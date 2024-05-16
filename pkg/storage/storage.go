package storage

import (
	"fmt"
	"log"
	"sync"

	"github.com/i5heu/ouroboros-db/internal/keyValStore"
	"github.com/i5heu/ouroboros-db/pkg/buzhashChunker"
	"github.com/i5heu/ouroboros-db/pkg/types"
)

type storage struct {
	kv *keyValStore.KeyValStore
}

type StoreFileOptions struct {
	EventToAppendTo types.Event
	Metadata        []byte
	File            []byte
	Temporary       bool
	FullTextSearch  bool
}

func NewStorage(kv *keyValStore.KeyValStore) StorageService {
	return &storage{
		kv: kv,
	}
}

func (ss *storage) Close() {
	ss.kv.Close()
}

// will store the file in the chunkStore and create new Event as child of given event
func (ss *storage) StoreFile(options StoreFileOptions) (types.Event, error) {
	// Validate options before proceeding
	err := options.ValidateOptions()
	if err != nil {
		log.Fatalf("Error validating options: %v", err)
		return types.Event{}, err
	}

	// Create channels to handle asynchronous results and errors
	fileChunkKeysChan := make(chan [][64]byte, 1)
	metadataChunkKeysChan := make(chan [][64]byte, 1)
	errorChan := make(chan error, 2) // buffer for two possible errors

	var wg sync.WaitGroup
	wg.Add(2)

	// Store file data in chunk store asynchronously
	go func() {
		defer wg.Done()
		keys, err := ss.storeDataInChunkStore(options.File)
		if err != nil {
			errorChan <- err
			return
		}
		fileChunkKeysChan <- keys
	}()

	// Store metadata in chunk store asynchronously
	go func() {
		defer wg.Done()
		keys, err := ss.storeDataInChunkStore(options.Metadata)
		if err != nil {
			errorChan <- err
			return
		}
		metadataChunkKeysChan <- keys
	}()

	// Wait for both goroutines to complete
	wg.Wait()
	close(errorChan)
	close(fileChunkKeysChan)
	close(metadataChunkKeysChan)

	// Check for errors
	for err := range errorChan {
		log.Printf("Error in storing data: %v", err)
		return types.Event{}, err
	}

	// Retrieve results from channels
	fileChunkKeys := <-fileChunkKeysChan
	metadataChunkKeys := <-metadataChunkKeysChan

	// Create a new event
	newEvent, err := ss.CreateNewEvent(EventOptions{
		ContentHashes:     fileChunkKeys,
		MetadataHashes:    metadataChunkKeys,
		HashOfParentEvent: options.EventToAppendTo.EventHash,
		Temporary:         options.Temporary,
		FullTextSearch:    options.FullTextSearch,
	})
	if err != nil {
		log.Fatalf("Error creating new event: %v", err)
		return types.Event{}, err
	}

	return newEvent, nil
}

func (options *StoreFileOptions) ValidateOptions() error {
	if options.EventToAppendTo.EventHash == [64]byte{} {
		return fmt.Errorf("Error storing file: Parent event was not defined")
	}

	if len(options.File) == 0 && len(options.Metadata) == 0 {
		return fmt.Errorf("Error storing file: Both file and metadata are empty")
	}

	if !options.Temporary && options.EventToAppendTo.Temporary {
		return fmt.Errorf("Error storing file: Parent event is Temporary and can not have non-Temporary children")
	}

	return nil
}

func (ss *storage) GetFile(eventOfFile types.Event) ([]byte, error) {
	file := []byte{}

	for _, hash := range eventOfFile.ContentHashes {
		chunk, err := ss.kv.Read(hash[:])
		if err != nil {
			return nil, fmt.Errorf("Error reading chunk from GetFile: %v", err)
		}

		file = append(file, chunk...)
	}

	return file, nil
}

func (ss *storage) GetMetadata(eventOfFile types.Event) ([]byte, error) {
	metadata := []byte{}

	for _, hash := range eventOfFile.MetadataHashes {
		chunk, err := ss.kv.Read(hash[:])
		if err != nil {
			return nil, fmt.Errorf("Error reading chunk from GetMetadata: %v", err)
		}

		metadata = append(metadata, chunk...)
	}

	return metadata, nil
}

func (ss *storage) storeDataInChunkStore(data []byte) ([][64]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("Error storing data: Data is empty")
	}

	chunks, err := buzhashChunker.ChunkBytes(data)
	if err != nil {
		log.Fatalf("Error chunking data: %v", err)
		return nil, err
	}

	var keys [][64]byte

	for _, chunk := range chunks {
		keys = append(keys, chunk.Hash)
	}

	err = ss.kv.BatchWriteNonExistingChunks(chunks)
	if err != nil {
		log.Fatalf("Error writing chunks: %v", err)
		return nil, err
	}

	return keys, nil
}

func (ss *storage) GarbageCollection() error {
	return ss.kv.Clean()
}

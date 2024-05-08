package storage

import (
	"fmt"
	"log"
	"time"

	"github.com/i5heu/ouroboros-db/internal/binaryCoder"
	"github.com/i5heu/ouroboros-db/pkg/types"
)

const (
	RootEventPrefix = "RootEvent:"
)

var RootEventParentEventHash [64]byte

func init() {
	RootEventParentEventHash = [64]byte{'R', 'o', 'o', 't', 'E', 'v', 'e', 'n', 't'}
}

// Same as Event struct in storageService.go but without some unnecessary fields

func (ss *Storage) CreateRootEvent(title string) (types.Event, error) {
	// Create a new IndexEvent
	item := types.Event{
		Key:            []byte{},
		Level:          time.Now().UnixNano(),
		EventHash:      [64]byte{},
		ContentHashes:  [][64]byte{},
		MetadataHashes: [][64]byte{},
	}

	// Check if RootEvent with the same title already exists
	otherRootEvent, err := ss.kv.GetItemsWithPrefix([]byte("RootEvent:" + title + ":"))
	if err != nil {
		log.Fatalf("Error getting keys: %v", err)
		return types.Event{}, err
	}

	if len(otherRootEvent) > 0 {
		log.Fatalf("Error creating new root event: RootEvent with the same title already exists")
		return types.Event{}, fmt.Errorf("Error creating new root event: RootEvent with the same title already exists")
	}

	item.MetadataHashes, err = ss.storeDataInChunkStore([]byte(title))
	if err != nil {
		log.Fatalf("Error storing metadata: %v", err)
		return types.Event{}, err
	}

	item.HashOfParentEvent = [64]byte{'R', 'o', 'o', 't', 'E', 'v', 'e', 'n', 't'}
	item.HashOfRootEvent = [64]byte{'R', 'o', 'o', 't', 'E', 'v', 'e', 'n', 't'}
	item.EventHash = item.CreateDetailsMetaHash()
	item.Key = GenerateKeyFromPrefixAndHash("Event:", item.EventHash)

	// Serialize the EventChainItem using gob
	data, err := binaryCoder.EventToByte(item)
	if err != nil {
		log.Fatalf("Error serializing EventChainItem: %v", err)
		return types.Event{}, err
	}

	// Write the EventChainItem to the keyValStore
	ss.kv.Write(item.Key, data)

	// define the event as entry point for the EventChain
	ss.kv.Write([]byte("RootEvent:"+title+":"+fmt.Sprint(item.Level)), item.Key)

	return item, err
}

func (ss *Storage) GetAllRootEvents() ([]types.Event, error) {
	// Get all keys from the keyValStore
	rootIndex, err := ss.GetRootIndex()
	if err != nil {
		log.Fatalf("Error getting root index: %v", err)
		return nil, err
	}

	if len(rootIndex) == 0 {
		return nil, nil
	}

	rootEvents := []types.Event{}

	for _, indexItem := range rootIndex {
		rootEvent, error := ss.GetEvent(indexItem.Hash)
		if error != nil {
			log.Fatalf("Error getting root event: %v", error)
			return nil, error
		}

		rootEvents = append(rootEvents, rootEvent)
	}

	return rootEvents, nil
}

func (ss *Storage) GetRootIndex() ([]types.RootEventsIndex, error) {
	// Get all keys from the keyValStore
	rootIndex, err := ss.kv.GetItemsWithPrefix([]byte("RootEvent:"))
	if err != nil {
		log.Fatalf("Error getting keys: %v", err)
		return nil, err
	}

	if len(rootIndex) == 0 {
		return nil, nil
	}

	revi := []types.RootEventsIndex{}
	for _, item := range rootIndex {
		revi = append(revi, types.RootEventsIndex{Title: item[0], Hash: GetEventHashFromKey(item[1])})
	}

	return revi, nil
}

func (ss *Storage) GetRootEventsWithTitle(title string) ([]types.Event, error) {
	rootIndex, err := ss.kv.GetItemsWithPrefix([]byte("RootEvent:" + title + ":"))
	if err != nil {
		log.Fatalf("Error getting keys: %v", err)
		return nil, err
	}

	if len(rootIndex) == 0 {
		return nil, nil
	}

	rootEvents := []types.Event{}
	for _, item := range rootIndex {
		rootEvent, error := ss.GetEvent(GetEventHashFromKey(item[1]))
		if error != nil {
			log.Fatalf("Error getting root event: %v", error)
			return nil, error
		}

		rootEvents = append(rootEvents, rootEvent)
	}

	return rootEvents, nil
}
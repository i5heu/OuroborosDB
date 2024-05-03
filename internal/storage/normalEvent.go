package storage

import (
	"bytes"
	"crypto/sha512"
	"encoding/gob"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"time"
)

// Event represents an event in the EventChain, the absolute top of a EventChain is a RootEvent, look at rootEvents.go
type Event struct {
	Key               []byte     // type:title:[DetailsMetaHash] title is only used for RootEvents, DetailsMetaHash is the hash of all ContentHashes, MetadataHashes, Level, HashOfParentEvent and HashOfSourceEvent and Temporary ("true" or "false" string)
	EventHash         [64]byte   // SHA-512 hash of all ContentHashes, MetadataHashes, Level, HashOfParentEvent and HashOfSourceEvent and Temporary ("true" or "false" string)
	Level             int64      // unix timestamp of creation time
	ContentHashes     [][64]byte // all hashes in a certain order that allow to reconstruct the content from the values in badgerDB
	MetadataHashes    [][64]byte // If no metadata is present the metadata of the the next previous event with metadata is used
	HashOfParentEvent [64]byte   // the event this item is the child of
	HashOfRootEvent   [64]byte   // the first event that marks a now EventChain
	Temporary         bool       // no non-Temporary event can have a Temporary event as Parent, Temporary events will be removed after some conditions are met, if one deletes a Temporary event all its children will be deleted too
	FullTextSearch    bool       // if true the ContentHashes and MetadataHashes of the event will be used for full-text search, Important: this only applies to the event itself, not to its children
}

type EventOptions struct {
	HashOfParentEvent [64]byte
	ContentHashes     [][64]byte // optional
	MetadataHashes    [][64]byte // optional
	Temporary         bool       // optional
	FullTextSearch    bool       // optional
}

func (ss *Service) CreateNewEvent(options EventOptions) (Event, error) {
	// Create a new Event
	item := Event{
		Key:               []byte{},
		EventHash:         [64]byte{},
		Level:             time.Now().UnixNano(),
		ContentHashes:     options.ContentHashes,
		MetadataHashes:    options.MetadataHashes,
		HashOfParentEvent: options.HashOfParentEvent,
		HashOfRootEvent:   [64]byte{},
		Temporary:         false,
		FullTextSearch:    false,
	}

	// check if the parent event exists
	if item.HashOfParentEvent == [64]byte{} {
		log.Fatalf("Error creating new event: Parent event was not defined")
		return Event{}, errors.New("Error creating new event: Parent event was not defined")
	}

	// check if the parent event exists
	parentEvent, err := ss.GetEvent(item.GetParentEventKey())
	if err != nil {
		log.Fatalf("Error creating new event: Parent event does not exist")
		return Event{}, errors.New("Error creating new event: Parent event does not exist")
	}

	if !item.Temporary {
		// check if the parent event is not Temporary
		if parentEvent.Temporary {
			log.Fatalf("Error creating new event: Parent event is Temporary and can not have non-Temporary children")
			return Event{}, errors.New("Error creating new event: Parent event is Temporary and can not have non-Temporary children")
		}
	}

	item.EventHash = item.CreateDetailsMetaHash()
	item.Key = GenerateKeyFromPrefixAndHash("Event:", item.EventHash)

	// Serialize the EventChainItem using gob
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(item); err != nil {
		log.Fatalf("Error encoding item: %v", err)
		return Event{}, err
	}

	// Write the EventChainItem to the keyValStore

	err = ss.kv.Write(item.Key, buf.Bytes())
	if err != nil {
		log.Fatalf("Error writing item: %v", err)
		return Event{}, err
	}

	return item, err
}

func (item *Event) CreateDetailsMetaHash() [64]byte {
	// create a buffer, put all bytes into the buffer and hash it
	buffer := make([]byte, 0)

	buffer = append(buffer, byte(item.Level))
	for _, hash := range item.ContentHashes {
		buffer = append(buffer, hash[:]...)
	}
	for _, hash := range item.MetadataHashes {
		buffer = append(buffer, hash[:]...)
	}
	buffer = append(buffer, item.HashOfParentEvent[:]...)
	buffer = append(buffer, item.HashOfRootEvent[:]...)
	if item.Temporary {
		buffer = append(buffer, []byte("true")...)
	} else {
		buffer = append(buffer, []byte("false")...)
	}

	return sha512.Sum512(buffer)
}

func (ss *Service) GetEvent(key []byte) (Event, error) {
	// Read the EventChainItem from the keyValStore
	value, err := ss.kv.Read(key)
	if err != nil {
		return Event{}, fmt.Errorf("Error reading Event with Key: %x: %v", hex.EncodeToString(key), err)
	}

	// Deserialize the EventChainItem using gob
	item := Event{}
	dec := gob.NewDecoder(bytes.NewReader(value))
	if err := dec.Decode(&item); err != nil {
		log.Fatalf("Error decoding RootEvent with Key: %x, Value: %x: %v", hex.EncodeToString(key), hex.EncodeToString(value), err)
		return Event{}, err
	}

	return item, nil
}
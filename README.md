 [![Go Reference](https://pkg.go.dev/badge/github.com/i5heu/ouroboros-db.svg)](https://pkg.go.dev/github.com/i5heu/ouroboros-db) [![Tests](https://github.com/i5heu/OuroborosDB/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/i5heu/OuroborosDB/actions/workflows/go.yml) ![](https://github.com/i5heu/OuroborosDB/blob/gh-pages/.badges/main/coverage.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/i5heu/ouroboros-db)](https://goreportcard.com/report/github.com/i5heu/ouroboros-db) [![wakatime](https://wakatime.com/badge/github/i5heu/ouroboros-db.svg)](https://wakatime.com/badge/github/i5heu/ouroboros-db)
 
<p align="center">
  <img src=".media/logo.jpeg"  width="300">
</p>

⚠️ This Project is still in development and not ready for production use. ⚠️

# OuroborosDB
A embedded database built around the concept of event trees, emphasizing data deduplication and data integrity checks. By structuring data into event trees, OuroborosDB ensures efficient and intuitive data management. Key features include:

- Data Deduplication: Eliminates redundant data through efficient chunking and hashing mechanisms.
- Data Integrity Checks: Uses SHA-512 hashes to verify the integrity of stored data.
- Event-Based Architecture: Organizes data hierarchically for easy retrieval and management.
- Scalable Concurrent Processing: Optimized for concurrent processing to handle large-scale data.
- Log Management and Indexing: Provides efficient logging and indexing for performance monitoring.
- Non-Deletable Events: Once stored, events cannot be deleted or altered, ensuring the immutability and auditability of the data.
- (To be implemented) Temporary Events: Allows the creation of temporary events that can be marked as temporary and safely cleaned up later for short-term data storage needs.


## Table of Contents

- [OuroborosDB](#ouroborosdb)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Initialization](#initialization)
    - [Storing Files](#storing-files)
    - [Retrieving Files](#retrieving-files)
    - [Event Tree Management](#event-tree-management)
      - [Creating Root Event](#creating-root-event)
      - [Fetching Root Events by Title](#fetching-root-events-by-title)
      - [Creating Child Events](#creating-child-events)
      - [Fetching Child Events](#fetching-child-events)
  - [Testing](#testing)
  - [Benchmarking](#benchmarking)
    - [Benchmark current state of the codebase](#benchmark-current-state-of-the-codebase)
    - [Benchmark Versions](#benchmark-versions)
  - [OuroborosDB Performance Version Differences](#ouroborosdb-performance-version-differences)
  - [OuroborosDB Performance Changelog](#ouroborosdb-performance-changelog)
  - [Current Features](#current-features)
  - [Future Features](#future-features)
  - [Current Problems and things to research:](#current-problems-and-things-to-research)
  - [Name and Logo](#name-and-logo)
  - [License](#license)

## Installation

OuroborosDB requires Go 1.21.5+

```bash
go get -u github.com/i5heu/OuroborosDB
```

## Usage

### Initialization

**OuroborosDB** can be initialized with a configuration struct that includes paths for storage and other settings.

```go
import "OuroborosDB"

func initializeDB() *OuroborosDB.OuroborosDB {
    db, err := OuroborosDB.NewOuroborosDB(OuroborosDB.Config{
        Paths:                     []string{"./data/storage"},
        MinimumFreeGB:             1,
        GarbageCollectionInterval: 10, // Minutes
    })
    if err != nil {
        log.Fatalf("Failed to initialize OuroborosDB: %v", err)
    }
    return db
}
```

### Storing Files

Files can be stored within events using the `StoreFile` method.

```go
import (
    "OuroborosDB/internal/storage"
    "OuroborosDB"
)

func storeFile(db *OuroborosDB.OuroborosDB, parentEvent storage.Event) storage.Event {
    fileContent := []byte("This is a sample file content")
    metadata := []byte("sample.txt")

    event, err := db.DB.StoreFile(storage.StoreFileOptions{
        EventToAppendTo: parentEvent,
        Metadata:        metadata,
        File:            fileContent,
    })
    if err != nil {
        log.Fatalf("Failed to store file: %v", err)
    }
    return event
}
```

### Retrieving Files

Files can be retrieved by providing the event from which they were stored.

```go
func retrieveFile(db *OuroborosDB.OuroborosDB, event storage.Event) []byte {
    content, err := db.DB.GetFile(event)
    if err != nil {
        log.Fatalf("Failed to retrieve file: %v", err)
    }
    return content
}
```

### Event Tree Management

#### Creating Root Event

Create a root event to represent the top level of an event tree.

```go
func createRootEvent(db *OuroborosDB.OuroborosDB, title string) storage.Event {
    rootEvent, err := db.DB.CreateRootEvent(title)
    if err != nil {
        log.Fatalf("Failed to create root event: %v", err)
    }
    return rootEvent
}
```

#### Fetching Root Events by Title

```go
func getRootEventsByTitle(db *OuroborosDB.OuroborosDB, title string) []storage.Event {
    events, err := db.DB.GetRootEventsWithTitle(title)
    if err != nil {
        log.Fatalf("Failed to fetch root events by title: %v", err)
    }
    return events
}
```

#### Creating Child Events

```go
func createChildEvent(db *OuroborosDB.OuroborosDB, parentEvent storage.Event) storage.Event {
    childEvent, err := db.DB.CreateNewEvent(storage.EventOptions{
        HashOfParentEvent: parentEvent.EventHash,
    })
    if err != nil {
        log.Fatalf("Failed to create child event: %v", err)
    }
    return childEvent
}
```

#### Fetching Child Events

```go
func getChildEvents(db *OuroborosDB.OuroborosDB, parentEvent storage.Event) []storage.Event {
    children, err := db.Index.GetDirectChildrenOfEvent(parentEvent.EventHash)
    if err != nil {
        log.Fatalf("Failed to fetch child events: %v", err)
    }
    return children
}
```

## Testing
  
```bash
  go test ./...
```

## Benchmarking
### Benchmark current state of the codebase
```bash
  go test -run='^$' -bench=.
```
### Benchmark Versions 
Works with committed changes and version/commits that are reachable by `git checkout`.  
You also need to have installed `benchstat` to compare the benchmarks, install it with `go get golang.org/x/perf/cmd/benchstat@latest`

```bash
  # add versions to bench.sh
  bash bench.sh
  # Now look in benchmarks/combined_benchmarks_comparison to see the results
```
## OuroborosDB Performance Version Differences
```bash
goos: linux
goarch: arm64
pkg: github.com/i5heu/ouroboros-db
                                                           │ benchmarks/v0.0.5.txt │       benchmarks/v0.0.8.txt        │       benchmarks/v0.0.14.txt        │
                                                           │        sec/op         │    sec/op     vs base              │    sec/op     vs base               │
_setupDBWithData/RebuildIndex-8                                       645.8m ±  9%   657.2m ±  7%       ~ (p=0.589 n=6)   669.7m ±  3%        ~ (p=0.093 n=6)
_Index_RebuildingIndex/RebuildIndex-8                                 20.78m ± 11%   19.98m ± 20%       ~ (p=0.699 n=6)   23.04m ± 11%  +10.86% (p=0.041 n=6)
_Index_GetDirectChildrenOfEvent/GetChildrenOfEvent-8                  4.500µ ±  6%   4.841µ ±  2%  +7.58% (p=0.004 n=6)   5.267µ ±  5%  +17.04% (p=0.002 n=6)
_Index_GetChildrenHashesOfEvent/GetChildrenHashesOfEvent-8            77.12n ±  3%   76.19n ±  4%       ~ (p=0.699 n=6)   77.81n ±  3%        ~ (p=0.310 n=6)
_DB_StoreFile/StoreFile-8                                             179.9µ ±  7%   185.9µ ±  5%       ~ (p=0.485 n=6)   187.9µ ±  6%        ~ (p=0.589 n=6)
_DB_GetFile/GetFile-8                                                 3.613µ ±  3%   3.667µ ±  2%       ~ (p=0.132 n=6)   3.575µ ±  3%        ~ (p=0.589 n=6)
_DB_GetEvent/GetEvent-8                                               6.099µ ±  3%   6.286µ ±  1%  +3.07% (p=0.002 n=6)   6.314µ ±  4%   +3.53% (p=0.026 n=6)
_DB_GetMetadata/GetMetadata-8                                         3.982µ ±  3%   3.958µ ±  4%       ~ (p=0.589 n=6)   3.929µ ±  4%        ~ (p=0.394 n=6)
_DB_GetAllRootEvents/GetAllRootEvents-8                               19.42m ±  5%   19.44m ±  3%       ~ (p=0.818 n=6)   19.62m ±  6%        ~ (p=0.180 n=6)
_DB_GetRootIndex/GetRootIndex-8                                       2.455m ±  2%   2.487m ±  5%       ~ (p=0.310 n=6)   2.453m ±  6%        ~ (p=0.818 n=6)
_DB_GetRootEventsWithTitle/GetRootEventsWithTitle-8                   12.33µ ±  1%   12.54µ ±  3%       ~ (p=0.180 n=6)   12.24µ ±  2%        ~ (p=0.589 n=6)
_DB_CreateRootEvent/CreateRootEvent-8                                 133.1µ ±  4%   131.8µ ±  2%       ~ (p=0.589 n=6)   133.3µ ±  9%        ~ (p=1.000 n=6)
_DB_CreateNewEvent/CreateNewEvent-8                                   49.86µ ± 19%   50.50µ ± 15%       ~ (p=0.818 n=6)   51.24µ ± 14%        ~ (p=0.818 n=6)
geomean                                                               108.4µ         109.7µ        +1.13%                 111.7µ         +3.05%
```

## OuroborosDB Performance Changelog

- **v0.0.14** - Major refactor of the Event type and introduction of FastMeta which should speed up search
- **v0.0.3** - Switch from `gob` to `protobuf` for serialization
- **v0.0.2** - Create tests and benchmarks


## Current Features
- [x] Data Deduplication
- [x] Basic Data Store and Retrieval
- [ ] Data Integrity Checks
- [x] Child to Parent Index

## Future Features
- [ ] Full Text Search
- [ ] Is the deletion of not Temporary Events a good idea?
  - Maybe if only some superUser can delete them with a key or something. 
- [ ] It would be nice to have pipelines that can run custom JS or webassembly to do arbitrary things.   
  - With http routing we could build a webserver that can run inside a "pipeline" in the database. sksksk
  - They should be usable as scraper or time or event based notificators.
  - Like if this event gets a child recursively, upload this tree to a server.
    - this would need a virtual folder structure that is represented in an event.    
    - with this we could also build a webdav server that can be used to access parts of the database. 

## Current Problems and things to research:
- [ ] Garbage Collection would delete Chunks that in the process of being used in a new event
- [ ] Deletion of Temporary Events is not yet discovered
- [ ] We have EventChilds that are used as either
  - A Item in the "category" of the Event
  - New Information that replaces it's Parent
  - Changes to the Parent (think patches)     
  We need to reflect this in the Event Structure to lower chunk lookups  
  If we implement a potential DeltaEvent, we need to provide tooling for it.
    - is it like git where we have a diff of the file?
    - is it a new file that replaces the old one?
    - we already have the chunk system in place. But this seams to not be suitable for text files - so we would need a text based delta event?


## Name and Logo

The name "OuroborosDB" is derived from the ancient symbol "Ouroboros," a representation of cyclical events, continuity, and endless return. Historically, it's been a potent symbol across various cultures, signifying the eternal cycle of life, death, and rebirth. In the context of this database, the Ouroboros symbolizes the perpetual preservation and renewal of data. While the traditional Ouroboros depicts a serpent consuming its tail, our version deviates, hinting at both reverence for historical cycles and the importance of continuous adaptation in the face of change.

## License
OuroborosDB (c) 2024 Mia Heidenstedt and contributors  
   
SPDX-License-Identifier: AGPL-3.0

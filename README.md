# UUID Version 7

## Introduction
UUID (Universally Unique Identifier) version 7 is a new draft standard that introduces a time-based structure to UUID generation. This version provides properties that are particularly advantageous for ordering and indexing in modern systems.

## Key Features of UUID Version 7

### 1. Time-Ordered Generation
UUIDv7 incorporates a timestamp as the primary component in its structure. This means that UUIDs generated sequentially are naturally ordered by time. Such time-ordering ensures that UUIDs can be used effectively for indexing and sorting in databases.

### 2. Structural Composition
- **Timestamp**: The most significant 48 bits of UUIDv7 represent a Unix timestamp in milliseconds.
- **Randomness**: The remaining bits are reserved for randomness, ensuring uniqueness even in distributed systems generating UUIDs at high throughput.
- **128-Bit Length**: Like other UUIDs, version 7 maintains a total size of 128 bits.

### 3. Lexicographic Properties
UUID version 7's time-based structure enables lexicographical (alphabetical) ordering. This property is especially useful for:
- **Efficient Indexing**: UUIDs maintain a natural sort order, making database indexing and time-based data retrieval efficient.
- **Range Queries**: Querying by a range of timestamps becomes straightforward due to the time component's positioning.
- **Predictable Sequence**: Sequential inserts benefit from the partially predictable nature of UUIDv7, as it sorts new UUIDs after older ones when arranged in ascending order.

## Advantages Over Older UUID Versions
- **Better Indexing**: UUIDv7 offers improved efficiency for indexing and searching compared to UUIDv4, which is fully random.
- **Reduced Collisions**: While UUIDv1 includes time but has potential privacy concerns and uniqueness limitations in distributed environments, UUIDv7 mitigates these issues by including randomness in the non-timestamp bits.

# Golang UUID v7 Generator and Parser

UUID 
This project is a Go package that provides functionality for generating and parsing UUID v7 values. The package is designed to support both millisecond and nanosecond granularity. The UUIDs generated follow the UUID v7 format, which is time-based and ideal for sorting by generation time.

## Overview
UUID v7 combines time-based data for ordering and randomness for uniqueness. This implementation ensures that:
- UUIDs have millisecond-level public granularity.
- Nanosecond-level precision is achieved privately by encoding additional random data.

## Installation
To use this package, import it into your Go project:

```go
import "github.com/emersonary/uuidv7"
```

## Public Methods
Below are the public methods exposed by this package:

### `NewUUIDv7`
Generates a new UUID v7 based on the current time.

**Signature:**
```go
func NewUUIDv7() uuid.UUID
```

### `UUIDv7FromTimeStamp`
Generates a UUID v7 from a provided `time.Time` object.

**Signature:**
```go
func UUIDv7FromTimeStamp(timeStamp time.Time) uuid.UUID
```

**Parameters:**
- `timeStamp` (type: `time.Time`): The time from which the UUID should be generated.

**Returns:**
- A `uuid.UUID` object.

### `TimeStampFromUUIDv7`
Extracts the `time.Time` timestamp from a UUID v7.

**Signature:**
```go
func TimeStampFromUUIDv7(uuid uuid.UUID) time.Time
```

**Parameters:**
- `uuid` (type: `uuid.UUID`): The UUID from which to extract the timestamp.

**Returns:**
- A `time.Time` object representing the timestamp.

## Example
```go
package main

import (
	"fmt"
	"time"
	"github.com/google/uuid"
	"github.com/yourusername/yourproject/utils"
)

func main() {
	// Generate a new UUID v7
	newUUID := utils.NewUUIDv7()
	fmt.Printf("Generated UUID v7: %s\n", newUUID)

	// Generate a UUID v7 from a specific timestamp
	timeStamp := time.Date(2024, 11, 15, 12, 0, 0, 0, time.UTC)
	specificUUID := utils.UUIDv7FromTimeStamp(timeStamp)
	fmt.Printf("UUID v7 from timestamp: %s\n", specificUUID)

	// Extract timestamp from UUID v7
	extractedTime := utils.TimeStampFromUUIDv7(newUUID)
	fmt.Printf("Extracted timestamp from UUID: %s\n", extractedTime)
}
```
---
Feel free to reach out with questions or contributions!


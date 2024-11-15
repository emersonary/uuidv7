# UUID v7 Generator and Parser

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


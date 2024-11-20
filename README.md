
# bincollection

A Go module to calculate the type of bin collection (e.g., Black bins or Red and Yellow bins) based on a start date and the current date.

## Features
- Determine bin collection type for a given week.
- Simple and reusable logic.

## Installation

To use bincollection in your project, add it as a dependency:

```
go get -u github.com/irrisdev/bin-collection
```

## Usage
```go
package main

import (
	"fmt"
	"time"

	"github.com/irrisdev/bin-collection/bincollection"
)

func main() {
	blackBinStartDate := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.Local)
	currentDate := time.Now()

	collectionType, err := bincollection.GetBinCollectionType(blackBinStartDate, currentDate)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("This week is for: %s collection.\n", collectionType)
}
```
## License
[MIT License](LICENSE)

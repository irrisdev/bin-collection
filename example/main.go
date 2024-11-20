package main

import (
	"fmt"
	"time"

	bincollection "github.com/irrisdev/bin-collection"
)

func main() {

	blackBinStartDate := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.Local)

	collectionType := bincollection.GetBinCollectionType(blackBinStartDate)

	fmt.Printf("This week is for: %s collection.\n", collectionType)
}
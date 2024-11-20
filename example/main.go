package main

import (
	"fmt"
	"time"

	bincollection "github.com/irrisdev/bin-collection"
)

func main() {

	blackBinStartDate := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.Local)

	currentDate := time.Now()

	collectionType, err := bincollection.GetBinCollectionType(blackBinStartDate, currentDate)
	if err != nil {
		fmt.Printf("Error determining collection type: %s\n", err)
		return
	}

	fmt.Printf("This week is for: %s collection.\n", collectionType)
}
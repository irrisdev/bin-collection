package bincollection

import (
	"errors"
	"log"
	"time"
)

type BinType string

const (
	BlackBin         BinType = "Black bins"
	RedAndYellowBins BinType = "Red and Yellow bins"
)

func GetBinCollectionType(blackBinStartDate time.Time) (BinType) {
	
	currentDate := time.Now()
	
	if blackBinStartDate.IsZero() || currentDate.IsZero() {
		log.Fatal(errors.New("invalid date provided"))
	}

	_, collectionWeek := blackBinStartDate.ISOWeek()
	_, currentWeek := currentDate.ISOWeek()

	if (currentWeek-collectionWeek)%2 == 0 {
		return BlackBin
	}
	return RedAndYellowBins
}
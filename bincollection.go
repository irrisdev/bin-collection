package bincollection

import (
	"errors"
	"time"
)

type BinType string

const (
	BlackBin         BinType = "Black bins"
	RedAndYellowBins BinType = "Red and Yellow bins"
)

func GetBinCollectionType(blackBinStartDate time.Time, currentDate time.Time) (BinType, error) {
	if blackBinStartDate.IsZero() || currentDate.IsZero() {
		return "", errors.New("invalid date provided")
	}

	_, collectionWeek := blackBinStartDate.ISOWeek()
	_, currentWeek := currentDate.ISOWeek()

	if (currentWeek-collectionWeek)%2 == 0 {
		return BlackBin, nil
	}
	return RedAndYellowBins, nil
}
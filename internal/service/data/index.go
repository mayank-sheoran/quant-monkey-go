package data

import (
	dataModels "quant-monkey/internal/common/models/data"
	"time"
)

type GenericHistoricalData interface {
	fetchOHLCdataForToken(
		token string,
		timeFrame string,
		from time.Time,
		to time.Time,
	) []dataModels.OHLC_TV
}

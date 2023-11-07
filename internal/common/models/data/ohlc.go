package data

import "time"

type OHLC_TV struct {
	Open  float64
	High  float64
	Low   float64
	Close float64

	TimeStamp time.Time
	Volume    float64
}

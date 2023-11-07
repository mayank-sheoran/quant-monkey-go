package main

import (
	"fmt"
	enums "quant_monkey/internal/common/constants/enums"
	"quant_monkey/internal/service/broker_clients"
	data2 "quant_monkey/internal/service/data"
	"time"
)

func main() {
	// Load ENV
	loadEnvFile()

	// Init Broker Clients
	broker_clients.InitializeAllBrokerClients()

	layout := "2006-01-02 15:04"
	fromStr := "2023-02-10 09:15"
	toTime := time.Now()
	fromTime, _ := time.Parse(layout, fromStr)
	data := data2.Get_OHLC_DataFromBroker(
		enums.ANGLE_ONE,
		"3045",
		"ONE_DAY",
		fromTime,
		toTime,
	)
	fmt.Println(data)

}

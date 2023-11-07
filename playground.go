package main

import (
	"fmt"
	enums "quant_monkey/internal/common/constants/enums"
	"quant_monkey/internal/service/broker_clients"
	"time"
)

func main() {
	// Load ENV
	loadEnvFile()

	// Init Broker Clients
	broker_clients.InitializeAllBrokerClients()

	angleOneMasterClient := broker_clients.GetBrokerMasterClient(enums.ANGLE_ONE)

	if client, ok := angleOneMasterClient.(*broker_clients.AngleOneMasterClient); ok {
		layout := "2006-01-02 15:04"
		fromStr := "2023-02-10 09:15"
		toTime := time.Now()
		fromTime, _ := time.Parse(layout, fromStr)
		data := client.FetchOHLCdataForToken("3045", "ONE_DAY", fromTime, toTime)
		fmt.Println("The value is an integer", data)
	} else {
		fmt.Println("The value is not an integer")
	}

}

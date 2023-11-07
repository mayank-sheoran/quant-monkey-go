package main

import (
	"github.com/joho/godotenv"
	"quant_monkey/cmd"
	"quant_monkey/internal/common/logger"
	"quant_monkey/internal/service/broker_clients"
)

func main1() {
	// Load ENV
	loadEnvFile()

	// Init Broker Clients
	broker_clients.InitializeAllBrokerClients()

	// Start server
	cmd.StartListeningToApis()
}

func loadEnvFile() {
	if err := godotenv.Load(); err != nil {
		logger.LoggerClient.Error(err.Error())
	}
}

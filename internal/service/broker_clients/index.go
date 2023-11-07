package broker_clients

import (
	"quant_monkey/internal/common/constants/enums"
)

type InitializeBrokerClient interface {
	initMasterClient()
}

func initializeBroker(ibc InitializeBrokerClient) {
	ibc.initMasterClient()
}

func InitializeAllBrokerClients() {
	initializeBroker(angleOneMasterCl)
}

func GetBrokerMasterClient(broker string) interface{} {
	switch broker {
	case enums.ANGLE_ONE:
		return angleOneMasterCl
	}
	return nil
}

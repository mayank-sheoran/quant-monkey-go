package data

import (
	dataModels "quant_monkey/internal/common/models/data"
	"quant_monkey/internal/service/broker_clients"
	"time"
)

func Get_OHLC_DataFromBroker(
	brokerName string,
	token string,
	timeFrame string,
	from time.Time,
	to time.Time,
) []dataModels.OHLC_TV {
	var brokerMasterClient = broker_clients.GetBrokerMasterClient(brokerName)
	switch client := brokerMasterClient.(type) {
	case *broker_clients.AngleOneMasterClient:
		return client.FetchOHLCdataForToken(token, timeFrame, from, to)
	}
	return nil
}

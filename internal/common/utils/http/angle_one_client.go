package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	http2 "quant_monkey/internal/common/constants/http"
	"quant_monkey/internal/common/logger"
	"quant_monkey/internal/common/models/data"
	"time"
)

func GetAngleOneGenericHttpHeaders(apiKey string, authToken string, headers http.Header) http.Header {
	if headers == nil {
		headers = map[string][]string{}
	}
	headers.Add("X-PrivateKey", apiKey)
	headers.Add("Authorization", "Bearer "+authToken)
	headers.Add("Content-Type", "application/json")
	headers.Add("X-ClientLocalIP", "localIp")
	headers.Add("X-ClientPublicIP", "publicIp")
	headers.Add("X-MACAddress", "mac")
	headers.Add("Accept", "application/json")
	headers.Add("X-UserType", "USER")
	headers.Add("X-SourceID", "WEB")
	return headers
}

func GetAngleOneGenericReqBody(symbolToken string, timeFrame string, from time.Time, to time.Time) []byte {
	return []byte(fmt.Sprintf(`{
      "exchange": "NSE",
      "symboltoken": "%s",
      "interval": "%s",
      "fromdate": "%s",
      "todate": "%s"
 }`, symbolToken, timeFrame, from.Format("2006-01-02 15:04"), to.Format("2006-01-02 15:04")))
}

type AngleOneSuccessEnvelope struct {
	Status    bool            `json:"status"`
	Message   string          `json:"message"`
	ErrorCode string          `json:"errorcode"`
	Data      [][]interface{} `json:"data"`
}

func EnvelopSuccessResponse(resp HTTPResponse) (AngleOneSuccessEnvelope, error) {
	if resp.Response.StatusCode >= http.StatusBadRequest {
		logger.LoggerClient.Error(http2.HTTP_BAD_RESPONSE_PARSE_FAILURE_ANGLE_ONE)
		return AngleOneSuccessEnvelope{}, fmt.Errorf(http2.HTTP_BAD_RESPONSE_PARSE_FAILURE_ANGLE_ONE)
	}
	var envelope AngleOneSuccessEnvelope
	if err := json.Unmarshal(resp.Body, &envelope); err != nil {
		logger.LoggerClient.Error(err.Error())
		return AngleOneSuccessEnvelope{}, err
	}
	return envelope, nil
}

func Get_OHLC_TV_fromSuccessEnvelope(envelope AngleOneSuccessEnvelope) []data.OHLC_TV {
	var ohlcData []data.OHLC_TV
	for _, row := range envelope.Data {
		ohlc := data.OHLC_TV{}
		timestampStr := row[0].(string)
		if t, err := time.Parse(time.RFC3339, timestampStr); err == nil {
			ohlc.TimeStamp = t
		}
		ohlc.Open = row[1].(float64)
		ohlc.High = row[2].(float64)
		ohlc.Low = row[3].(float64)
		ohlc.Close = row[4].(float64)
		ohlc.Volume = row[5].(float64)
		ohlcData = append(ohlcData, ohlc)
	}
	return ohlcData
}

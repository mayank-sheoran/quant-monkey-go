package broker_clients

import (
	"github.com/angel-one/smartapigo"
	"net/http"
	"os"
	"quant_monkey/internal/common/constants"
	httpConstants "quant_monkey/internal/common/constants/http"
	"quant_monkey/internal/common/logger"
	dataModels "quant_monkey/internal/common/models/data"
	"quant_monkey/internal/common/utils"
	httpUtils "quant_monkey/internal/common/utils/http"
	"time"
)

type AngleOneMasterClient struct {
	smartApiGoClient *smartapigo.Client
	userSession      smartapigo.UserSession
	httpClient       httpUtils.HTTPClient
}

const (
	ANGLE_ONE_BASE_URI             = "https://apiconnect.angelbroking.com"
	ANGLE_ONE_HISTORICAL_END_POINT = "/rest/secure/angelbroking/historical/v1/getCandleData"
)

var (
	angleOneMasterCl = &AngleOneMasterClient{
		smartApiGoClient: &smartapigo.Client{},
		userSession:      smartapigo.UserSession{},
	}
)

func (aomc *AngleOneMasterClient) initMasterClient() {
	var totp, err = utils.GenerateTOTP(os.Getenv(constants.ANGLE_ONE_TOTP_CODE))
	aomc.smartApiGoClient = smartapigo.New(
		os.Getenv(constants.ANGLE_ONE_CLIENT_ID),
		os.Getenv(constants.ANGLE_ONE_PASSWORD),
		os.Getenv(constants.ANGLE_ONE_API_KEY))
	aomc.httpClient = httpUtils.GenerateHttpClient(&http.Client{Timeout: httpConstants.API_REQUEST_TIMEOUT}, false)
	if err != nil {
		logger.LoggerClient.Error(httpConstants.TOTP_GENERATION_ERROR)
	} else {
		var userSession, sessionErr = aomc.smartApiGoClient.GenerateSession(totp)
		if sessionErr != nil {
			logger.LoggerClient.Error(httpConstants.TOTP_GENERATION_ERROR)
		} else {
			aomc.userSession = userSession
		}
	}
}

func (aomc *AngleOneMasterClient) FetchOHLCdataForToken(
	token string,
	timeFrame string,
	from time.Time,
	to time.Time,
) []dataModels.OHLC_TV {
	var apiKey = os.Getenv(constants.ANGLE_ONE_API_KEY)
	var headers = httpUtils.GetAngleOneGenericHttpHeaders(apiKey, aomc.userSession.AccessToken, nil)
	var reqBody = httpUtils.GetAngleOneGenericReqBody(token, timeFrame, from, to)
	response, _ := aomc.httpClient.DoRaw(http.MethodPost, ANGLE_ONE_BASE_URI+ANGLE_ONE_HISTORICAL_END_POINT, reqBody, headers)
	angleOneSuccessEnvelope, _ := httpUtils.EnvelopSuccessResponse(response)
	return httpUtils.Get_OHLC_TV_fromSuccessEnvelope(angleOneSuccessEnvelope)
}

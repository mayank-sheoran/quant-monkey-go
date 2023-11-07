package health_check

import (
	"github.com/gorilla/mux"
	"net/http"
	"quant_monkey/internal/common/constants"
	"quant_monkey/internal/common/models/response"
	http2 "quant_monkey/internal/common/utils/http"
)

func HandleHealthCheckRoutes(rootApiRouter *mux.Router) {
	rootApiRouter.HandleFunc(constants.HEALTH_CHECK, healthCheckHandler).Methods(constants.GET)
}

func healthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	genericResponseMapper := response.GenericResponse{
		Status:  200,
		Message: constants.HEALTH_CHECK_MESSAGE,
	}
	http2.SendHttpResponse(w, genericResponseMapper, http.StatusOK)
}

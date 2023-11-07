package cmd

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"quant_monkey/internal/common/constants"
	"quant_monkey/internal/common/logger"
	"quant_monkey/pkg/api"
)

var (
	rootApiRouter = mux.NewRouter()
)

func StartListeningToApis() {
	api.HandleAllRoutes(rootApiRouter)
	http.Handle("/", rootApiRouter)
	logger.LoggerClient.Info(constants.SERVER_STARTED + os.Getenv(constants.SERVER_PORT))
	err := http.ListenAndServe(":"+os.Getenv(constants.SERVER_PORT), nil)
	if err != nil {
		logger.LoggerClient.Error(err.Error())
	}
}

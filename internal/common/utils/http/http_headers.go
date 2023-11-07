package http

import (
	"encoding/json"
	"net/http"
	"quant_monkey/internal/common/logger"
)

func SendHttpResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.LoggerClient.Error(err.Error())
	}
}

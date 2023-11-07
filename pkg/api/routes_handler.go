package api

import (
	"github.com/gorilla/mux"
	"quant_monkey/pkg/api/routes/health_check"
)

func HandleAllRoutes(rootApiRouter *mux.Router) {
	// Routes
	health_check.HandleHealthCheckRoutes(rootApiRouter)
}

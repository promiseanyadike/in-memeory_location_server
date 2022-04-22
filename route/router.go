package route

import (
	"in-memory_location_server/service"
	"net/http"

	"github.com/gorilla/mux"
)

func NewHandler(locationService *service.Location) http.Handler {
	reportHandler := NewReportHandler(locationService)
	router := mux.NewRouter()
	router.HandleFunc("/location/{order_id}/now", reportHandler.AddLocation).Methods(http.MethodPost).
		HeadersRegexp("Content-Type", "application/json")

	router.HandleFunc("/location/{order_id}", reportHandler.GetLocation).Methods(http.MethodGet)

	router.HandleFunc("/location/{order_id}", reportHandler.DeleteLocation).Methods(http.MethodDelete)

	return router
}

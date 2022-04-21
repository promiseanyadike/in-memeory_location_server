package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"in-memory_location_server/model"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type ReportHandler struct {
	service *Location
}

func NewReportHandler(locationService *Location) *ReportHandler {
	return &ReportHandler{
		service: locationService,
	}
}

func (h *ReportHandler) AddLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId, ok := vars["order_id"]
	if !ok {
		err := fmt.Errorf("error during rrid parsing")
		log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	payload := new(model.Location)
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = payload.Validate()
	if err != nil {
		log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.service.AppendLocation(orderId, *payload)
	respondWithJSON(w, http.StatusOK, payload)
}

func (h *ReportHandler) GetLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId, ok := vars["order_id"]
	if !ok {
		err := fmt.Errorf("error during rrid parsing")
		log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var max int
	maxStr := r.URL.Query().Get("max")
	if len(maxStr) > 0 {
		var err error
		max, err = strconv.Atoi(maxStr)
		if err != nil {
			log.Err(err).Send()
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	response, err := h.service.GetLocation(orderId, max)
	if err != nil {
		if err != nil {
			log.Err(err).Send()
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}
	respondWithJSON(w, http.StatusOK, LocationPayload{
		OrderId: orderId,
		History: response,
	})
}

func (h *ReportHandler) DeleteLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId, ok := vars["order_id"]
	if !ok {
		err := fmt.Errorf("error during rrid parsing")
		log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := h.service.DeleteLocation(orderId)
	if err != nil {
		log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// respondWithJSON write json response format
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		log.Err(err).Send()
	}
}

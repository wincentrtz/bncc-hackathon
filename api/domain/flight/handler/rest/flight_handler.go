package handler

import (
	"encoding/json"
	"net/http"

	"github.com/wincentrtz/bncc/api/domain/flight"

	"github.com/gorilla/mux"
)

type FlightHandler struct {
	FlightUsecase flight.Usecase
}

func NewFlightHandler(r *mux.Router, fu flight.Usecase) {
	handler := &FlightHandler{
		FlightUsecase: fu,
	}
	r.HandleFunc("/api/flight", handler.FetchFlight).Methods("GET")
}

func (fh *FlightHandler) FetchFlight(w http.ResponseWriter, r *http.Request) {
	flights, err := fh.FlightUsecase.FetchFlight()
	if err != nil {
		panic("ERROR")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(flights)
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

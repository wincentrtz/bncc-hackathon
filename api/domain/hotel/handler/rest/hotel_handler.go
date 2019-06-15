package handler

import (
	"encoding/json"
	"net/http"

	"github.com/wincentrtz/bncc/api/domain/hotel"

	"github.com/gorilla/mux"
)

type HotelHandler struct {
	HotelUsecase hotel.Usecase
}

func NewHotelHandler(r *mux.Router, hu hotel.Usecase) {
	handler := &HotelHandler{
		HotelUsecase: hu,
	}
	r.HandleFunc("/api/hotel", handler.Hotel).Methods("GET")
}

func (hh *HotelHandler) Hotel(w http.ResponseWriter, r *http.Request) {
	hotels, err := hh.HotelUsecase.FetchHotel()
	if err != nil {
		panic("ERROR")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(hotels)
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

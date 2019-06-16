package handler

import (
	"encoding/json"
	"net/http"

	"github.com/wincentrtz/bncc/api/domain/album"

	"github.com/gorilla/mux"
)

// AlbumHandler struct
type AlbumHandler struct {
	AlbumUsecase album.Usecase
}

// NewAlbumHandler factory
func NewAlbumHandler(r *mux.Router, al album.Usecase) {
	handler := &AlbumHandler{
		AlbumUsecase: al,
	}
	r.HandleFunc("/api/album", handler.FetchAll).Methods("GET")
}

// FetchAll command
func (ah *AlbumHandler) FetchAll(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	// vars := mux.Vars(r)

	albums, err := ah.AlbumUsecase.FetchAll()
	if err != nil {
		panic("ERROR")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(albums)
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

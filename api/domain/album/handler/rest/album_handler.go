package handler

import (
	"encoding/json"
	"net/http"

	"github.com/wincentrtz/bncc/api/domain/album"

	"github.com/gorilla/mux"
)

type AlbumHandler struct {
	AlbumUsecase album.Usecase
}

func NewAlbumHandler(r *mux.Router, hu album.Usecase) {
	handler := &AlbumHandler{
		AlbumUsecase: hu,
	}
	r.HandleFunc("/api/album", handler.FetchAlbum).Methods("GET")
}

func (hh *AlbumHandler) FetchAlbum(w http.ResponseWriter, r *http.Request) {
	albums, err := hh.AlbumUsecase.FetchAlbum()
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

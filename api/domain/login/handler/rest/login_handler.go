package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/wincentrtz/bncc/api/domain/login"
	"github.com/wincentrtz/bncc/api/models/request"

	"github.com/gorilla/mux"
)

type LoginHandler struct {
	LoginUsecase login.Usecase
}

func NewLoginHandler(r *mux.Router, lu login.Usecase) {
	handler := &LoginHandler{
		LoginUsecase: lu,
	}
	r.HandleFunc("/api/login", handler.Login).Methods("POST")
}

func (lh *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	var loginRequest request.LoginRequest
	_ = json.NewDecoder(r.Body).Decode(&loginRequest)

	login, err := lh.LoginUsecase.Login(&loginRequest)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, `{"status": "Invalid login"}`)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(login)
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

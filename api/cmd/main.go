package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/wincentrtz/bncc/config"

	userHandler "github.com/wincentrtz/bncc/domain/user/handler/rest"
	_userRepository "github.com/wincentrtz/bncc/domain/user/repository"
	_userUsecase "github.com/wincentrtz/bncc/domain/user/usecase"
)

func main() {
	db := config.InitDb()
	defer db.Close()
	r := mux.NewRouter()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	registerUserHandler(r, timeoutContext, db)

	fmt.Println("Starting..")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func registerUserHandler(r *mux.Router, timeoutContext time.Duration, db *sql.DB) {
	ur := _userRepository.NewUserRepository(db)
	us := _userUsecase.NewUserUsecase(ur, timeoutContext)
	userHandler.NewUserHandler(r, us)
}

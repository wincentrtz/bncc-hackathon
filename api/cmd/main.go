package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/wincentrtz/bncc/api/config"

	userHandler "github.com/wincentrtz/bncc/api/domain/user/handler/rest"
	_userRepository "github.com/wincentrtz/bncc/api/domain/user/repository"
	_userUsecase "github.com/wincentrtz/bncc/api/domain/user/usecase"

	loginHandler "github.com/wincentrtz/bncc/api/domain/login/handler/rest"
	_loginRepository "github.com/wincentrtz/bncc/api/domain/login/repository"
	_loginUsecase "github.com/wincentrtz/bncc/api/domain/login/usecase"

	hotelHandler "github.com/wincentrtz/bncc/api/domain/hotel/handler/rest"
	_hotelRepository "github.com/wincentrtz/bncc/api/domain/hotel/repository"
	_hotelUsecase "github.com/wincentrtz/bncc/api/domain/hotel/usecase"

	albumHandler "github.com/wincentrtz/bncc/api/domain/album/handler/rest"
	_albumRepository "github.com/wincentrtz/bncc/api/domain/album/repository"
	_albumUsecase "github.com/wincentrtz/bncc/api/domain/album/usecase"
)

func main() {
	db := config.InitDb()
	defer db.Close()
	r := mux.NewRouter()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	registerUserHandler(r, timeoutContext, db)
	registerLoginHandler(r, timeoutContext, db)
	registerHotelHandler(r, timeoutContext, db)
	registerAlbumHandler(r, timeoutContext, db)

	fmt.Println("Starting..")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func registerUserHandler(r *mux.Router, timeoutContext time.Duration, db *sql.DB) {
	ur := _userRepository.NewUserRepository(db)
	us := _userUsecase.NewUserUsecase(ur, timeoutContext)
	userHandler.NewUserHandler(r, us)
}

func registerLoginHandler(r *mux.Router, timeoutContext time.Duration, db *sql.DB) {
	lr := _loginRepository.NewLoginRepository(db)
	ls := _loginUsecase.NewLoginUsecase(lr, timeoutContext)
	loginHandler.NewLoginHandler(r, ls)
}

func registerHotelHandler(r *mux.Router, timeoutContext time.Duration, db *sql.DB) {
	hr := _hotelRepository.NewHotelRepository(db)
	hs := _hotelUsecase.NewHotelUsecase(hr, timeoutContext)
	hotelHandler.NewHotelHandler(r, hs)
}

func registerAlbumHandler(r *mux.Router, timeoutContext time.Duration, db *sql.DB) {
	ar := _albumRepository.NewAlbumRepository(db)
	as := _albumUsecase.NewAlbumUsecase(ar, timeoutContext)
	albumHandler.NewAlbumHandler(r, as)
}

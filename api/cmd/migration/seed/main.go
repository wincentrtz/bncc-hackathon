package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/bxcodec/faker"
	"github.com/wincentrtz/bncc/api/config"
	"github.com/wincentrtz/bncc/api/models"
)

func insertUserDataToDB(user models.User, encryptedPassword string) {
	db := config.InitDb()
	defer db.Close()

	query := `
		INSERT INTO users 
		VALUES(DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := db.Exec(
		query,
		&user.Name,
		&user.Email,
		encryptedPassword,
		&user.Phone,
		&user.Address,
		&user.Gender,
		&user.Ktp,
		time.Now())

	if err != nil {
		panic(err)
	}
}

func insertHotelDataToDB(hotel models.Hotel) {
	db := config.InitDb()
	defer db.Close()

	query := `
		INSERT INTO hotel 
		VALUES(DEFAULT, $1, $2, $3, $4, $5, $6)
	`
	_, err := db.Exec(
		query,
		&hotel.Name,
		&hotel.Phone,
		&hotel.Address,
		&hotel.City,
		&hotel.Price,
		time.Now())

	if err != nil {
		panic(err)
	}
}

func populateUserData(number int) {
	for i := 1; i <= number; i++ {
		user := models.User{}
		err := faker.FakeData(&user)
		if err != nil {
			fmt.Println(err)
		}

		hasher := md5.New()
		hasher.Write([]byte("user" + strconv.Itoa(i)))
		encryptedPassword := hex.EncodeToString(hasher.Sum(nil))

		user.Name = "User " + strconv.Itoa(i)
		user.Email = "user" + strconv.Itoa(i) + "@gmail.com"

		insertUserDataToDB(user, encryptedPassword)
	}
}

func populateHotelData(number int) {
	city := []string{
		"Bali",
		"Bali",
		"Bali",
		"Bali",
		"Jakarta",
		"Bandung",
	}

	for i := 1; i <= number; i++ {
		hotel := models.Hotel{}
		err := faker.FakeData(&hotel)
		if err != nil {
			fmt.Println(err)
		}

		hotel.Name = "Hotel " + strconv.Itoa(i)
		hotel.City = city[i]
		hotel.Price = hotel.Price * 100000

		insertHotelDataToDB(hotel)
	}
}

func main() {
	populateUserData(20)
	populateHotelData(5)
}

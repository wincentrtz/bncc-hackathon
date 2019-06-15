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

func insertUserDataToDB(user models.User) {
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
		&user.Password,
		&user.Phone,
		&user.Address,
		&user.Gender,
		&user.Ktp,
		time.Now())

	if err != nil {
		panic(err)
	}
}

func populateUserData(number int) {
	for i := 0; i < number; i++ {
		user := models.User{}
		err := faker.FakeData(&user)
		if err != nil {
			fmt.Println(err)
		}

		hasher := md5.New()
		hasher.Write([]byte("user" + strconv.Itoa(i)))
		encryptedPassword := hex.EncodeToString(hasher.Sum(nil))

		user.Password = encryptedPassword
		user.Name = "User " + strconv.Itoa(i)
		user.Email = "user" + strconv.Itoa(i) + "@gmail.com"

		insertUserDataToDB(user)
	}
}

func main() {
	populateUserData(20)
}

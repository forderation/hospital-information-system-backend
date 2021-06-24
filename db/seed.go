package db

import (
	"crypto/md5"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"io"
	"log"
)

func SeedUser(db *DB, userLen int) {
	for i := 1; i <= userLen; i++ {
		randomProfile := randomdata.GenerateProfile(randomdata.RandomGender)
		h := md5.New()
		_, _ = io.WriteString(h, "password123")
		user := User{
			FirstName: randomProfile.Name.First,
			LastName:  randomProfile.Name.Last,
			Age:       uint(randomdata.Decimal(18, 30)),
			Email:     randomProfile.Email,
			Username:  randomProfile.Login.Username,
			Password:  fmt.Sprintf("%x", h.Sum(nil)),
			IsAdmin:   true,
		}
		err := db.CreateUser(&user)
		if err != nil {
			log.Fatalf("Could not create foreign key: %v", err.Error())
		}
	}
}

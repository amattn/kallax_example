package main

import (
	"log"
	"time"

	"./model"

	kallax "gopkg.in/src-d/go-kallax.v1"
)

func main() {

	db, err := openDB()
	if err != nil {
		log.Println(3103303200, err)
		return
	}
	log.Println(db)

	store := model.NewUserStore(db) // it just needs an instance of *sql.DB

	err = store.Insert(&model.User{
		Name:     "john",
		Email:    "john@doe.me",
		Passhash: "1234bunnies", // please properly salt and hash your passwords.
	})
	if err != nil {
		log.Println(3103303232, err)
		return
	}

	q := model.NewUserQuery().
		FindByName("john").
		FindByCreatedAt(kallax.Gt, time.Now().Add(-30*24*time.Hour))

	user, err := store.FindOne(q)
	if err != nil {
		log.Println(3103303283, err)
		return
	}
	log.Println(user)
}

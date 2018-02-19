package main

import (
	"log"
	"runtime"
	"time"

	"./model"

	kallax "gopkg.in/src-d/go-kallax.v1"
)

func main() {
	log.Printf("Starting kallax_example (%v, v%v, build %v)", runtime.Version(), Version(), BuildNumber())
	log.Println("Connecting to DB")

	db, err := openDB()
	if err != nil {
		log.Println(3103303200, "unable to open connection to DB\n", err)
		return
	}

	store := model.NewUserStore(db) // it just needs an instance of *sql.DB

	insert_user(store)

	query_user(store)
}

func insert_user(store *model.UserStore) {
	log.Println("Insert User")
	defer trace("insert_user", time.Now())

	err := store.Insert(&model.User{
		Name:     "john",
		Email:    "john@doe.me",
		Passhash: "1234bunnies", // please properly salt and hash your passwords.
	})
	if err != nil {
		log.Println(3103303232, err)
		return
	}
}

func query_user(store *model.UserStore) {
	log.Println("Query User")
	defer trace("query_user", time.Now())

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

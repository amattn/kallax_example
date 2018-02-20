package main

import (
	"log"
	"runtime"
	"time"

	"./model"

	"github.com/amattn/deeperror"

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

	costore := model.NewCompanyStore(db)
	userstore := model.NewUserStore(db)

	company, err := insert_company(costore)

	insert_user(userstore, company)

	query_user(userstore)
}

func insert_company(store *model.CompanyStore) (*model.Company, error) {
	log.Println("Insert Company")
	defer trace("insert_company", time.Now())

	comp := model.NewCompany()
	comp.Name = "SomeCo"
	comp.Address = "12345 Main St. Anytown, AA, USA"

	err := store.Insert(comp)
	if err != nil {
		derr := deeperror.New(2283340241, "Insert error", err)
		return nil, derr
	}

	return comp, nil
}

func insert_user(store *model.UserStore, comp *model.Company) {
	log.Println("Insert User")
	defer trace("insert_user", time.Now())

	err := store.Insert(&model.User{
		Name:     "john",
		Email:    "john@doe.me",
		Passhash: "1234bunnies", // please properly salt and hash your passwords.
		Company:  comp,
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

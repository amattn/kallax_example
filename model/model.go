package model

import kallax "gopkg.in/src-d/go-kallax.v1"

//go:generate kallax gen

type User struct {
	kallax.Model `table:"users" pk:"id,autoincr"`
	kallax.Timestamps
	ID       int64
	Email    string
	Salt     string
	Passhash string
	Name     string
	Phone    string

	Company *Company `fk:",inverse"`
}

type Company struct {
	kallax.Model `table:"companies" pk:"id,autoincr"`
	kallax.Timestamps
	ID      int64
	Name    string
	Address string

	Users []*User `fk:""`
}

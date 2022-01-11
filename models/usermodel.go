package models

import (
	"database/sql"
)

type User struct {
	FirstName   sql.NullString `schema:"firstName"`
	LastName    sql.NullString `schema:"lastName"`
	PhoneNumber sql.NullString `schema:"phoneNumber"`
	StreetAddrA sql.NullString `schema:"streetAddrA"`
	StreetAddrB sql.NullString `schema:"streetAddrB"`
	City        sql.NullString `schema:"city"`
	Email       sql.NullString `schema:"email"`
	State       sql.NullString `schema:"state"`
	Country     sql.NullString `schema:"country"`
	Zipcode     sql.NullInt32  `schema:"zipcode"`
	BirthDay    sql.NullInt32  `schema:"birthDay"`
	BirthMonth  sql.NullInt32  `schema:"birthMonth"`
	BirthYear   sql.NullInt32  `schema:"birthYear"`
	Title       sql.NullString `schema:"title"`
}

type UserStore interface {
	CreateUser(u User) bool
}

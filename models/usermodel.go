package models

import (
	"database/sql"
)

type User struct {
	FirstName       sql.NullString `schema:"firstName"`
	LastName        sql.NullString `schema:"lastName"`
	PhoneNumber     sql.NullString `schema:"phoneNumber"`
	Email           sql.NullString `schema:"email"`
	StreetAddrAShip sql.NullString `schema:"streetAddrA_ship"`
	StreetAddrBShip sql.NullString `schema:"streetAddrB_ship"`
	CityShip        sql.NullString `schema:"city_ship"`
	StateShip       sql.NullString `schema:"state_ship"`
	CountryShip     sql.NullString `schema:"country_ship"`
	ZipcodeShip     sql.NullInt32  `schema:"zipcode_ship"`
	StreetAddrABill sql.NullString `schema:"streetAddrA_bill"`
	StreetAddrBBill sql.NullString `schema:"streetAddrB_bill"`
	CityBill        sql.NullString `schema:"city_bill"`
	StateBill       sql.NullString `schema:"state_bill"`
	CountryBill     sql.NullString `schema:"country_bill"`
	ZipcodeBill     sql.NullInt32  `schema:"zipcode_bill"`
	BirthDay        sql.NullInt32  `schema:"birthDay"`
	BirthMonth      sql.NullInt32  `schema:"birthMonth"`
	BirthYear       sql.NullInt32  `schema:"birthYear"`
	Title           sql.NullString `schema:"title"`
}

type UserRequests interface {
	CreateUser(u User) error
}

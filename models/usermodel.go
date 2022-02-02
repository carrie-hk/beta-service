package models

type User struct {
	FirstName       *string `db:"firstName"`
	LastName        *string `db:"lastName"`
	PhoneNumber     *string `db:"phoneNumber"`
	Email           *string `db:"email"`
	StreetAddrAShip *string `db:"streetAddrA_ship"`
	StreetAddrBShip *string `db:"streetAddrB_ship"`
	CityShip        *string `db:"city_ship"`
	StateShip       *string `db:"state_ship"`
	CountryShip     *string `db:"country_ship"`
	ZipcodeShip     *int64  `db:"zipcode_ship"`
	StreetAddrABill *string `db:"streetAddrA_bill"`
	StreetAddrBBill *string `db:"streetAddrB_bill"`
	CityBill        *string `db:"city_bill"`
	StateBill       *string `db:"state_bill"`
	CountryBill     *string `db:"country_bill"`
	ZipcodeBill     *int32  `db:"zipcode_bill"`
	BirthDay        *int32  `db:"birthDay"`
	BirthMonth      *int32  `db:"birthMonth"`
	BirthYear       *int32  `db:"birthYear"`
	Title           *string `db:"title"`
}

type UserRequests interface {
	CreateUser(u User) error
}

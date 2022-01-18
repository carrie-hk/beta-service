package models

type User struct {
	FirstName       string `schema:"firstName" binding:"required" json:"firstName" db:"firstName"`
	LastName        string `schema:"lastName" binding:"required"`
	PhoneNumber     string `schema:"phoneNumber" binding:"required"`
	Email           string `schema:"email" binding:"required"`
	StreetAddrAShip string `schema:"streetAddrA_ship"`
	StreetAddrBShip string `schema:"streetAddrB_ship"`
	CityShip        string `schema:"city_ship"`
	StateShip       string `schema:"state_ship"`
	CountryShip     string `schema:"country_ship"`
	ZipcodeShip     int32  `schema:"zipcode_ship"`
	StreetAddrABill string `schema:"streetAddrA_bill"`
	StreetAddrBBill string `schema:"streetAddrB_bill"`
	CityBill        string `schema:"city_bill"`
	StateBill       string `schema:"state_bill"`
	CountryBill     string `schema:"country_bill"`
	ZipcodeBill     int32  `schema:"zipcode_bill"`
	BirthDay        int32  `schema:"birthDay"`
	BirthMonth      int32  `schema:"birthMonth"`
	BirthYear       int32  `schema:"birthYear"`
	Title           string `schema:"title"`
}

type UserRequests interface {
	CreateUser(u User) error
}

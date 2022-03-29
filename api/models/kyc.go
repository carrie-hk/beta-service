package models

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

type KYC struct {
	Wallet_PK    string `db:"wallet_pk" json:"wallet_pk" validate:"required"`
	First_Name   string `db:"first_name" json:"first_name" validate:"required"`
	Last_Name    string `db:"last_name" json:"last_name" validate:"required"`
	Phone_Number string `db:"phone_num" json:"phone_num" validate:"required,e164"`
	Email        string `db:"email" json:"email" validate:"required,email"`
	Ship_Addr_A  string `db:"ship_addr_a" json:"ship_addr_a" validate:"required"`
	Ship_Addr_B  string `db:"ship_addr_b" json:"ship_addr_b"`
	Ship_City    string `db:"ship_city" json:"ship_city" validate:"required"`
	Ship_State   string `db:"ship_state" json:"ship_state"`
	Ship_ZIP     string `db:"ship_zip" json:"ship_zip" validate:"required"`
	Ship_Country string `db:"ship_country" json:"ship_country" validate:"required"`
	Dob_Day      int32  `db:"dob_day" json:"dob_day" validate:"required,gte=1,lte=31"`
	Dob_Month    int32  `db:"dob_month" json:"dob_month" validate:"required,gte=1,lte=12"`
	Dob_Year     int32  `db:"dob_year" json:"dob_year" validate:"required,gte=1900"`
}

func (kyc *KYC) Validate() error {
	err := kyc.validateAge()
	if err != nil {
		return err
	}
	validate := validator.New()
	return validate.Struct(kyc)
}

func (kyc *KYC) validateAge() error {
	TWENTY_ONE_YEARS_NS := 662709600000000000
	birthDate := time.Date(int(kyc.Dob_Year), time.Month(kyc.Dob_Month), int(kyc.Dob_Day), 0, 0, 0, 0, time.UTC)
	if int(time.Since(birthDate)) < TWENTY_ONE_YEARS_NS {
		return errors.New("Error: does not minimum age requirement")
	}
	return nil
}

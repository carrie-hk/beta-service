package models

import "testing"

func Test_ValidateKycInfo(t *testing.T) {
	k := &KYC{
		Wallet_PK:    "1q2w3e4r5t6y7u8i9o0p9o8i7u6y5t4r",
		First_Name:   "Janet",
		Last_Name:    "Snakehole",
		Phone_Number: "+14204204204",
		Email:        "janet.snakehole@fireball.ca",
		Ship_Addr_A:  "420 Down-the-Drain Way",
		Ship_City:    "Pawnee",
		Ship_State:   "Indiana",
		Ship_Country: "US&A",
		Ship_ZIP:     "46001",
		Dob_Day:      20,
		Dob_Month:    04,
		Dob_Year:     1999,
	}

	err := k.Validate()

	if err != nil {
		t.Fatal(err)
	}
}

func Test_ValidateInternational(t *testing.T) {
	k := &KYC{
		Wallet_PK:    "1q2w3e4r5t6y7u8i9o0p9o8i7u6y5t4r",
		First_Name:   "Janet",
		Last_Name:    "Snakehole",
		Phone_Number: "+14204204204",
		Email:        "janet.snakehole@fireball.ca",
		Ship_Addr_A:  "420 Down-the-Drain Way",
		Ship_City:    "Lisbon",
		Ship_State:   "",
		Ship_Country: "Portugal",
		Ship_ZIP:     "0",
		Dob_Day:      20,
		Dob_Month:    04,
		Dob_Year:     1999,
	}

	err := k.Validate()

	if err != nil {
		t.Fatal(err)
	}
}

func Test_MissingFirstNameFails(t *testing.T) {
	k := &KYC{
		Wallet_PK: "1q2w3e4r5t6y7u8i9o0p9o8i7u6y5t4r",
		// First_Name:   "Janet",
		Last_Name:    "Snakehole",
		Phone_Number: "+14204204204",
		Email:        "janet.snakehole@fireball.ca",
		Ship_Addr_A:  "420 Down-the-Drain Way",
		Ship_City:    "Pawnee",
		Ship_State:   "Indiana",
		Ship_ZIP:     "46001",
		Ship_Country: "US&A",
		Dob_Day:      20,
		Dob_Month:    04,
		Dob_Year:     1999,
	}

	err := k.Validate()

	if err == nil {
		t.Fatal(err)
	}
}

func Test_IncorrectPhoneNumberFormatFails(t *testing.T) {
	k := &KYC{
		Wallet_PK:  "1q2w3e4r5t6y7u8i9o0p9o8i7u6y5t4r",
		First_Name: "Janet",
		Last_Name:  "Snakehole",
		// Phone number is missing country code
		Phone_Number: "4204204204",
		Email:        "janet.snakehole@fireball.ca",
		Ship_Addr_A:  "420 Down-the-Drain Way",
		Ship_City:    "Pawnee",
		Ship_State:   "Indiana",
		Ship_Country: "US&A",
		Ship_ZIP:     "46001",
		Dob_Day:      20,
		Dob_Month:    04,
		Dob_Year:     1999,
	}

	err := k.Validate()

	if err == nil {
		t.Fatal(err)
	}
}

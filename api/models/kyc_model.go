package models

type KYC struct {
	Wallet_PK    *string `db:"wallet_pk" json:"wallet_pk"`
	First_Name   *string `db:"first_name" json:"first_name"`
	Last_Name    *string `db:"last_name" json:"last_name"`
	Phone_Number *string `db:"phone_num" json:"phone_num"`
	Email        *string `db:"email" json:"email"`
	Ship_Addr_A  *string `db:"ship_addr_a" json:"ship_addr_a"`
	Ship_Addr_B  *string `db:"ship_addr_b" json:"ship_addr_b"`
	Ship_City    *string `db:"ship_city" json:"ship_city"`
	Ship_State   *string `db:"ship_state" json:"ship_state"`
	Ship_ZIP     *int64  `db:"ship_zip" json:"ship_zip"`
	Dob_Day      *int32  `db:"dob_day" json:"dob_day"`
	Dob_Month    *int32  `db:"dob_month" json:"dob_month"`
	Dob_Year     *int32  `db:"dob_year" json:"dob_year"`
	Title        *string `db:"title" json:"title"`
}

package models

import "github.com/go-playground/validator/v10"

type StatusUpdate struct {
	New_Status string `db:"new_status" json:"new_status" validate:"required"`
	AXU_ID     int64  `db:"axu_id" json:"axu_id" validate:"required"`
	Mint_Addr  string `db:"mint_addr" json:"mint_addr" validate:"required"`
}

func (su *StatusUpdate) Validate() error {
	validate := validator.New()
	return validate.Struct(su)
}

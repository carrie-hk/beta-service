package models

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type StatusUpdate struct {
	Asset_Status string `db:"asset_status" json:"asset_status" validate:"required"`
	Mint_Addr    string `db:"mint_addr" json:"mint_addr" validate:"required"`
}

func (su *StatusUpdate) Validate() error {
	err := su.validateStatus(STATUS_LIST)
	if err != nil {
		return err
	}
	validate := validator.New()
	return validate.Struct(su)
}

func (su *StatusUpdate) validateStatus(statuses []string) error {
	for _, status := range statuses {
		if status == su.Asset_Status {
			return nil
		}
	}
	return errors.New("Error: invalid status provided")
}

var STATUS_LIST = []string{"Minted", "Listed", "Sold", "Escrow", "Redeemed"}

package models

import "github.com/go-playground/validator/v10"

type RedemptionInfo struct {
	Wallet_PK                    string `db:"wallet_pk" json:"wallet_pk" validate:"required"`
	Redemption_Info_Account_Addr string `db:"redemption_info_accnt_addr" json:"redemption_info_accnt_addr" validate:"required"`
	Baxus_Escrow_Addr            string `db:"baxus_escrow_addr" json:"baxus_escrow_addr" validate:"required"`
	Mint_Addr                    string `db:"mint_addr" json:"mint_addr" validate:"required"`
}

func (ri *RedemptionInfo) Validate() error {
	validate := validator.New()
	return validate.Struct(ri)
}

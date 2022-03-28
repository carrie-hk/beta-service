package models

type RedemptionRequest struct {
	Mint_Addr string `db:"mint_addr" json:"mint_addr"`
}

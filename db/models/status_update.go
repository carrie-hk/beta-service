package models

type StatusUpdate struct {
	New_Status string `db:"asset_status" json:"asset_status"`
	AXU_ID     int64  `db:"axu_id" json:"axu_id"`
	Mint_Addr  string `db:"mint_addr" json:"mint_addr"`
}

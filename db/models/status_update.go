package models

type StatusUpdate struct {
	New_Status *string `db:"asset_status"`
	AXU_ID     *int64  `db:"axu_id"`
	Mint_Addr  *string `db:"mint_addr"`
}

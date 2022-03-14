package models

type StatusUpdate struct {
	AXU_ID     *int64  `db:"axu_id"`
	Mint_Addr  *string `db:"mint_addr"`
	New_Status *string `db:"status"`
}

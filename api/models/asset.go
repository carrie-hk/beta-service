package models

type AssetView struct {
	AXU_ID       *int64   `db:"axu_id" json:"axu_id"`
	ASC_Num      *int64   `db:"asc_num" json:"asc_num"`
	Time_Created *string  `db:"time_created" json:"time_created"`
	Asset_Status *string  `db:"asset_status" json:"asset_status"`
	Token_Addr   *string  `db:"token_addr" json:"token_addr"`
	Mint_Addr    *string  `db:"mint_addr" json:"mint_addr"`
	Update_Addr  *string  `db:"update_addr" json:"update_addr"`
	Price        *float64 `db:"price" json:"price"`
	Featured     *bool    `db:"featured" json:"featured"`
	HTML5        *string  `db:"html5" json:"html5"`
	Cover        *string  `db:"cover" json:"cover"`
	Name         *string  `db:"name" json:"name"`
	Age          *int64   `db:"age" json:"age"`
	Desc_Short   *string  `db:"desc_short" json:"desc_short"`
	Desc_Long    *string  `db:"desc_long" json:"desc_long"`
	ABV          *float64 `db:"abv" json:"abv"`
}

package models

import "database/sql"

type Asset struct {
	Name          sql.NullString  `db:"Name"`
	Files         sql.NullString  `db:"Files"`
	Link          sql.NullString  `db:"Link"`
	Price         sql.NullString  `db:"Price"`
	Status        sql.NullString  `db:"Status"`
	BottleID      sql.NullInt32   `db:"Bottle ID"`
	People        sql.NullString  `db:"People"`
	Description   sql.NullString  `db:"Description"`
	Age           sql.NullInt32   `db:"Age"`
	YearDistilled sql.NullString  `db:"Year Distilled"`
	Distillery    sql.NullString  `db:"Distillery"`
	Bottler       sql.NullString  `db:"Bottler"`
	Size          sql.NullInt32   `db:"Size"`
	YearBottled   sql.NullInt32   `db:"Year Bottled"`
	CaskType      sql.NullString  `db:"Cask Type"`
	BottleNum     sql.NullInt32   `db:"Bottle Num"`
	CaskNum       sql.NullInt32   `db:"Cask Num"`
	SerialNum     sql.NullInt32   `db:"Serial Num"`
	Grade         sql.NullString  `db:"Grade"`
	Packaging     sql.NullString  `db:"Packaging"`
	Series        sql.NullString  `db:"Series"`
	Spirit        sql.NullString  `db:"Spirit"`
	Country       sql.NullString  `db:"Country"`
	OgCaskYield   sql.NullString  `db:"Original Cask Yield"`
	Region        sql.NullString  `db:"Region"`
	ABV           sql.NullFloat64 `db:"ABV"`
	ASCNum        sql.NullInt32   `db:"ASC Token #"`
	Dropbox       sql.NullString  `db:"Dropbox"`
	PriceRealized sql.NullInt32   `db:"Price Realized"`
	BaxusRev      sql.NullInt32   `db:"BAXUS Rev"`
	COGS          sql.NullInt32   `db:"COGS"`
}

type AssetStore interface {
	Assets() []Asset
	FeaturedAssets() []Asset
}

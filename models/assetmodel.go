package models

type Asset struct {
	Name          *string  `db:"Name"`
	Files         *string  `db:"HTML5 Files"`
	Link          *string  `db:"MP4 Link"`
	Price         *string  `db:"Price"`
	Status        *string  `db:"Status"`
	BottleID      *int64   `db:"Bottle ID"`
	People        *string  `db:"People"`
	Description   *string  `db:"Description"`
	Age           *int64   `db:"Age"`
	YearDistilled *string  `db:"Year Distilled"`
	Distillery    *string  `db:"Distillery"`
	Bottler       *string  `db:"Bottler"`
	Size          *int64   `db:"Size"`
	YearBottled   *int64   `db:"Year Bottled"`
	CaskType      *string  `db:"Cask Type"`
	BottleNum     *string  `db:"Bottle Num"`
	CaskNum       *string  `db:"Cask Num"`
	SerialNum     *string  `db:"Serial Num"`
	Grade         *string  `db:"Grade"`
	Packaging     *string  `db:"Packaging"`
	Series        *string  `db:"Series"`
	Spirit        *string  `db:"Spirit"`
	Country       *string  `db:"Country"`
	OgCaskYield   *string  `db:"Original Cask Yield"`
	Region        *string  `db:"Region"`
	ABV           *float64 `db:"ABV"`
	ASCNum        *int64   `db:"ASC Token #"`
	Dropbox       *string  `db:"Dropbox"`
	PriceRealized *int64   `db:"Price Realized"`
	BaxusRev      *int64   `db:"BAXUS Rev"`
	COGS          *int64   `db:"COGS"`
	Box           *int64   `db:"Box"`
	Issues        *string  `db:"Issues"`
	ShelfLoc      *string  `db:"Shelf Loc"`
	S3            *string  `db:"S3"`
	Featured      *bool    `db:"Featured"`
	MSRP          *int64   `db:"MSRP"`
	AnnualizedRet *int64   `db:"Annualized_return"`
}

type AssetRequests interface {
	Assets() ([]Asset, error)
	FeaturedAssets() ([]Asset, error)
	TestAssets() ([]Asset, error)
}

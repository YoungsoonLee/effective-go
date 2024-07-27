package main

import (
	"fmt"
	"time"
)

// StockInfo is information about a stock
type StockInfo struct {
	Date   time.Time
	Symbol string
	Open   float64
	High   float64
	Low    float64
	Close  float64
}

// Key is a key for a stock
type Key struct {
	year   int
	month  time.Month
	day    int
	symbol string
}

// InfoDB is in memory stock database
type InfoDB struct {
	m map[Key]StockInfo
}

// Get return information for stock at date. If information is not found the
// second return value is false.
func (db *InfoDB) Get(date time.Time, symbol string) (StockInfo, bool) {
	k := Key{date.Year(), date.Month(), date.Day(), symbol}
	v, ok := db.m[k]
	return v, ok
}

func main() {
	// Init the database
	db := InfoDB{m: make(map[Key]StockInfo)}

	// Add some data
	db.m[Key{2021, time.January, 1, "AAPL"}] = StockInfo{
		Date:   time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		Symbol: "AAPL",
		Open:   132.69,
		High:   133.61,
		Low:    132.69,
		Close:  133.52,
	}
	//fmt.Println(db.Get(time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC), "AAPL"))
	//fmt.Printf("%+v\n", db)

	// Add dup key data
	db.m[Key{2021, time.January, 1, "AAPL"}] = StockInfo{
		Date:   time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		Symbol: "TSLA",
		Open:   132.69,
		High:   133.61,
		Low:    132.69,
		Close:  133.52,
	}
	fmt.Printf("%+v\n", db)

	// check if the key is in the map
	if _, ok := db.m[Key{2021, time.January, 1, "AAPL"}]; ok {
		fmt.Println("key found")
	} else {
		fmt.Println("key not found")
	}

}

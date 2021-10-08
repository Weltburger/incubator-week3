package models

type Trade struct {
	EventType     string `json:"e"`
	EventTime     int    `json:"E"`
	Symbol        string `json:"s"`
	TradeID       int    `json:"t"`
	Price         string `json:"p"`
	Quantity      string `json:"q"`
	BuyerOrderID  int    `json:"b"`
	SellerOrderID int    `json:"a"`
	TradeTime     int    `json:"T"`
	MarketMaker   bool   `json:"m"`
	Ignore        bool   `json:"M"`
}

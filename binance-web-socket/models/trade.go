package models

type Trade struct {
	EventType    string `json:"e"`
	EventTime    int    `json:"E"`
	Symbol       string `json:"s"`
	TradeID      int    `json:"a"`
	Price        string `json:"p"`
	Quantity     string `json:"q"`
	FirstTradeID int    `json:"f"`
	LastTradeID  int    `json:"l"`
	TradeTime    int    `json:"T"`
	MarketMaker  bool   `json:"m"`
}

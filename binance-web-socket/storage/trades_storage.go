package storage

import (
	"binance-web-socket/models"
	"database/sql"
	"github.com/google/uuid"
	"log"
	"strconv"
	"time"
)

type TradesStorage struct {
	database *Database
}

func (tradesStorage *TradesStorage) Prepare() *sql.Stmt {
	st, err := tradesStorage.database.Tx.Prepare(`
		INSERT INTO trade.trades (
			uuid, 
			event_type, 
			event_time, 
			symbol, 
			trade_id, 
			price, 
			quantity, 
			first_trade_id, 
			last_trade_id, 
			trade_time, 
			market_maker,
		    ignore
		) VALUES (
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
		)`)

	if err != nil {
		log.Fatal(err)
	}

	return st
}

func (tradesStorage *TradesStorage) Exc(data *models.Trade) {
	etm := strconv.Itoa(data.EventTime)
	s, _ := strconv.ParseInt(etm[:10], 10, 64)
	ns, _ := strconv.ParseInt(etm[10:], 10, 64)
	eventTime := time.Unix(s, ns)
	ttm := strconv.Itoa(data.TradeTime)
	s, _ = strconv.ParseInt(ttm[:10], 10, 64)
	ns, _ = strconv.ParseInt(ttm[10:], 10, 64)
	tradeTime := time.Unix(s, ns)

	if _, err := tradesStorage.database.Stmt.Exec(
		uuid.New(),
		data.EventType,
		eventTime,
		data.Symbol,
		data.TradeID,
		data.Price,
		data.Quantity,
		data.BuyerOrderID,
		data.SellerOrderID,
		tradeTime,
		data.MarketMaker,
		data.Ignore,
	); err != nil {
		log.Fatal(err)
	}
}

func (tradesStorage *TradesStorage) Cmt() {
	if err := tradesStorage.database.Tx.Commit(); err != nil {
		log.Fatal(err)
	}
	tradesStorage.database.Tx, _ = tradesStorage.database.DB.Begin()
	tradesStorage.database.Stmt = tradesStorage.Prepare()
}

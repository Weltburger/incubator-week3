package storage

import (
	"binance-web-socket/models"
	"context"
	"log"
	"strconv"
	"time"
)

type TradesStorage struct {
	database *Database
}

func (tradesStorage *TradesStorage) AddTrade(ctx context.Context, trade *models.Trade) {
	etm := strconv.Itoa(trade.EventTime)
	s, _ := strconv.ParseInt(etm[:10], 10, 64)
	ns, _ := strconv.ParseInt(etm[10:], 10, 64)
	eventTime := time.Unix(s, ns)
	ttm := strconv.Itoa(trade.TradeTime)
	s, _ = strconv.ParseInt(ttm[:10], 10, 64)
	ns, _ = strconv.ParseInt(ttm[10:], 10, 64)
	tradeTime := time.Unix(s, ns)

	_, err := tradesStorage.database.db.ExecContext(ctx,
		`INSERT INTO "public"."trades" (type, event_time, symbol, trade_id, price, quantity, buyer_order_id, seller_order_id, trade_time, market_maker, ignore) 
			   VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		trade.EventType,
		eventTime,
		trade.Symbol,
		trade.TradeID,
		trade.Price,
		trade.Quantity,
		trade.BuyerOrderID,
		trade.SellerOrderID,
		tradeTime,
		trade.MarketMaker,
		trade.Ignore,
	)
	if err != nil {
		log.Fatal(err)
	}
}

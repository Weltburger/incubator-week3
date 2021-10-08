CREATE TABLE IF NOT EXISTS trade.trades (
    uuid            UUID,
    event_type      String,
    event_time      DateTime,
    symbol          String,
    trade_id        UInt64,
    price           String,
    quantity        String,
    buyer_order_id  UInt64,
    seller_order_id UInt64,
    trade_time      DateTime,
    market_maker    UInt8,
    ignore          UInt8
) engine = MergeTree()
      PARTITION BY toYYYYMMDD(event_time)
      order by  (event_type, symbol)

-------------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS trade.trades_rep (
    uuid            UUID,
    event_type      String,
    event_time      DateTime,
    symbol          String,
    trade_id        UInt64,
    price           String,
    quantity        String,
    buyer_order_id  UInt64,
    seller_order_id UInt64,
    trade_time      DateTime,
    market_maker    UInt8,
    ignore          UInt8
) engine = ReplacingMergeTree()
PARTITION BY toYYYYMMDD(event_time)
order by  (trade_id)
TTL event_time + INTERVAL 1 WEEK
SETTINGS index_granularity=262144,
         index_granularity_bytes=0

-------------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS trade.trades_sum (
    uuid            UUID,
    event_type      String,
    event_time      DateTime,
    symbol          String,
    trade_id        UInt64,
    price           String,
    quantity        String,
    buyer_order_id  UInt64,
    seller_order_id UInt64,
    trade_time      DateTime,
    market_maker    UInt8,
    ignore          UInt8
) engine = SummingMergeTree()
      PARTITION BY toYYYYMMDD(event_time)
      order by  (trade_id)
      TTL event_time + INTERVAL 1 WEEK
      SETTINGS index_granularity=262144,
          index_granularity_bytes=0

-------------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS trade.trades_agg (
    uuid            UUID,
    event_type      String,
    event_time      DateTime,
    symbol          String,
    trade_id        UInt64,
    price           String,
    quantity        String,
    buyer_order_id  UInt64,
    seller_order_id UInt64,
    trade_time      DateTime,
    market_maker    UInt8,
    ignore          UInt8
) engine = AggregatingMergeTree()
      PARTITION BY toYYYYMMDD(event_time)
      order by  (trade_id)
      TTL event_time + INTERVAL 1 WEEK
      SETTINGS index_granularity=262144,
          index_granularity_bytes=0
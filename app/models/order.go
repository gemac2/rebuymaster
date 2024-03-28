package models

import (
	"time"

	"github.com/adshao/go-binance/v2/futures"
	"github.com/gofrs/uuid"
	"github.com/hirokisan/bybit/v2"
)

type Order struct {
	ID                   uuid.UUID `db:"id" json:"id"`
	OrderType            string    `db:"order_type" json:"order_type"`
	CurrencyName         string    `db:"currency_name" json:"currency_name"`
	BuybackPercentage    float64   `db:"buyback_percentage" json:"buyback_percentage"`
	CurrencyPercentage   float64   `db:"currency_percentage" json:"currency_percentage"`
	StopLoss             float64   `db:"stop_loss" json:"stop_loss"`
	OrderPrice           float64   `db:"order_price" json:"order_price"`
	CurrencyQuantity     float64   `db:"currency_quantity" json:"currency_quantity"`
	Leverage             int       `db:"leverage" json:"leverage"`
	CreatedAt            time.Time `db:"created_at" json:"created_at"`
	UpdatedAt            time.Time `db:"updated_at" json:"updated_at"`
	IsBuybacksEnabled    bool      `db:"is_buybacks_enabled" json:"is_buybacks_enabled"`
	IsOrderPosition      bool      `db:"is_order_position" json:"is_order_position"`
	StopLossTaken        bool      `db:"stop_loss_taken" json:"stop_loss_taken"`
	TakeProfitAchieved   bool      `db:"take_profit_achieved" json:"take_profit_achieved"`
	TradeWon             bool      `db:"trade_won" json:"trade_won"`
	TradeLoss            bool      `db:"trade_loss" json:"trade_loss"`
	TakeProfitPercentage float64   `db:"take_profit" json:"take_profit"`
	Profit               float64   `db:"profit" json:"profit"`
}

// Orders struct
type Orders []Order

func GetBybitSymbol(currencyName string) bybit.SymbolV5 {
	switch currencyName {
	case "BTCUSDT":
		return bybit.SymbolV5BTCUSDT
	case "ETHUSDT":
		return bybit.SymbolV5ETHUSDT
	case "XRPUSDT":
		return bybit.SymbolV5XRPUSDT
	case "ADAUSDT":
		return bybit.SymbolV5ADAUSDT
	case "LTCUSDT":
		return bybit.SymbolV5LTCUSDT
	case "BNBUSDT":
		return bybit.SymbolV5BNBUSDT
	case "SOLUSDT":
		return bybit.SymbolV5SOLUSDT
	case "SXPUSDT":
		return bybit.SymbolV5SXPUSDT
	case "SUIUSDT":
		return bybit.SymbolV5SUIUSDT
	case "BLZUSDT":
		return bybit.SymbolV5BLZUSDT
	case "LPTUSDT":
		return bybit.SymbolV5LPTUSDT
	case "1000PEPEUSDT":
		return bybit.SymbolV51000PEPEUSDT
	case "MATICUSDT":
		return bybit.SymbolV5MATICUSDT
	case "TOMOUSDT":
		return bybit.SymbolV5TOMOUSDT
	case "LINAUSDT":
		return bybit.SymbolV5LINAUSDT
	case "RUNEUSDT":
		return bybit.SymbolV5RUNEUSDT
	case "OPUSDT":
		return bybit.SymbolV5OPUSDT
	case "DOGEUSDT":
		return bybit.SymbolV5DOGEUSDT
	case "UNFIUSDT":
		return bybit.SymbolV5UNFIUSDT
	case "APEUSDT":
		return bybit.SymbolV5APEUSDT
	case "GALAUSDT":
		return bybit.SymbolV5GALAUSDT
	case "APTUSDT":
		return bybit.SymbolV5APTUSDT
	case "LINKUSDT":
		return bybit.SymbolV5LINKUSDT
	case "XLMUSDT":
		return bybit.SymbolV5XLMUSDT
	case "CYBERUSDT":
		return bybit.SymbolV5CYBERUSDT
	case "STMXUSDT":
		return bybit.SymbolV5STMXUSDT
	case "PERPUSDT":
		return bybit.SymbolV5PERPUSDT
	case "TRBUSDT":
		return bybit.SymbolV5TRBUSDT
	case "HIFIUSDT":
		return bybit.SymbolV5HIFIUSDT
	case "OGNUSDT":
		return bybit.SymbolV5OGNUSDT
	case "AMBUSDT":
		return bybit.SymbolV5AMBUSDT
	case "LQTYUSDT":
		return bybit.SymbolV5LQTYUSDT
	case "LOOMUSDT":
		return bybit.SymbolV5LOOMUSDT
	case "JOEUSDT":
		return bybit.SymbolV5JOEUSDT
	case "BNTUSDT":
		return bybit.SymbolV5BNTUSDT
	case "POLYXUSDT":
		return bybit.SymbolV5POLYXUSDT
	case "MEMEUSDT":
		return bybit.SymbolV5MEMEUSDT
	case "ALGOUSDT":
		return bybit.SymbolV5ALGOUSDT
	case "10000LADYSUSDT":
		return bybit.SymbolV510000LADYSUSDT
	case "USTCUSDT":
		return bybit.SymbolV5USTCUSDT
	case "SUPERUSDT":
		return bybit.SymbolV5SUPERUSDT
	case "SEIUSDT":
		return bybit.SymbolV5SEIUSDT
	case "FRONTUSDT":
		return bybit.SymbolV5FRONTUSDT
	case "1000RATSUSDT":
		return bybit.SymbolV51000RATSUSDT
	case "GMTUSDT":
		return bybit.SymbolV5GMTUSDT
	case "BIGTIMEUSDT":
		return bybit.SymbolV5BIGTIMEUSDT
	case "SANDUSDT":
		return bybit.SymbolV5SANDUSDT
	case "LDOUSDT":
		return bybit.SymbolV5LDOUSDT
	case "LUNA2USDT":
		return bybit.SymbolV5LUNA2USDT
	case "XAIUSDT":
		return bybit.SymbolV5XAIUSDT
	case "BLURUSDT":
		return bybit.SymbolV5BLURUSDT
	case "CHZUSDT":
		return bybit.SymbolV5CHZUSDT
	case "NFPUSDT":
		return bybit.SymbolV5NFPUSDT
	case "ETCUSDT":
		return bybit.SymbolV5ETCUSDT
	case "UMAUSDT":
		return bybit.SymbolV5UMAUSDT
	case "MAVUSDT":
		return bybit.SymbolV5MAVUSDT
	case "MANTAUSDT":
		return bybit.SymbolV5MANTAUSDT
	case "WIFUSDT":
		return bybit.SymbolV5WIFUSDT
	case "PROMUSDT":
		return bybit.SymbolV5PROMUSDT
	case "PENDLEUSDT":
		return bybit.SymbolV5PENDLEUSDT
	case "JUPUSDT":
		return bybit.SymbolV5JUPUSDT
	case "ZETAUSDT":
		return bybit.SymbolV5ZETAUSDT
	case "PYTHUSDT":
		return bybit.SymbolV5PYTHUSDT
	case "ALTUSDT":
		return bybit.SymbolV5ALTUSDT
	case "ARKUSDT":
		return bybit.SymbolV5ARKUSDT
	case "TIAUSDT":
		return bybit.SymbolV5TIAUSDT
	case "DYMUSDT":
		return bybit.SymbolV5DYMUSDT
	case "CTSIUSDT":
		return bybit.SymbolV5CTSIUSDT
	case "COTIUSDT":
		return bybit.SymbolV5COTIUSDT
	case "QIUSDT":
		return bybit.SymbolV5QIUSDT
	case "INJUSDT":
		return bybit.SymbolV5INJUSDT
	case "LSKUSDT":
		return bybit.SymbolV5LSKUSDT
	case "MYROUSDT":
		return bybit.SymbolV5MYROUSDT
	case "SPELLUSDT":
		return bybit.SymbolV5SPELLUSDT
	case "CHRUSDT":
		return bybit.SymbolV5CHRUSDT
	case "DOTUSDT":
		return bybit.SymbolV5DOTUSDT
	}
	return ""
}

func GetBybitSide(orderType string) bybit.Side {
	switch orderType {
	case "Long":
		return bybit.SideBuy
	case "Short":
		return bybit.SideSell
	default:
		return bybit.SideBuy
	}
}

func GetBinanceSide(orderType string) futures.SideType {
	switch orderType {
	case "Long":
		return futures.SideTypeBuy
	case "Short":
		return futures.SideTypeSell
	default:
		return futures.SideTypeBuy
	}
}

func GetBinancePositionType(orderType string) futures.PositionSideType {
	switch orderType {
	case "Long":
		return futures.PositionSideTypeLong
	case "Short":
		return futures.PositionSideTypeShort
	default:
		return futures.PositionSideTypeBoth
	}
}

func GetBinanceStopLossSide(orderType string) futures.SideType {
	switch orderType {
	case "Long":
		return futures.SideTypeSell
	case "Short":
		return futures.SideTypeBuy
	default:
		return futures.SideTypeBuy
	}
}

func GetBinanceStopLossPositionType(orderType string) futures.PositionSideType {
	switch orderType {
	case "Long":
		return futures.PositionSideTypeShort
	case "Short":
		return futures.PositionSideTypeLong
	default:
		return futures.PositionSideTypeBoth
	}
}

func (o Order) SetStopLossPrice() float64 {
	posicion := (o.OrderPrice * o.CurrencyQuantity)
	slPrice := 0.0

	if o.OrderType == "Short" {
		slPrice = (posicion + o.StopLoss) / o.CurrencyQuantity
	}

	if o.OrderType == "Long" {
		slPrice = (posicion - o.StopLoss) / o.CurrencyQuantity
	}

	return slPrice
}

func GetBybitPositionMode(orderType string) *bybit.PositionIdx {
	position := 0
	switch orderType {
	case "Long":
		position = int(bybit.PositionIdxHedgeBuy)
		pointer := &position
		return (*bybit.PositionIdx)(pointer)
	case "Short":
		position = int(bybit.PositionIdxHedgeSell)
		pointer := &position
		return (*bybit.PositionIdx)(pointer)
	default:
		position = int(bybit.PositionIdxOneWay)
		pointer := &position
		return (*bybit.PositionIdx)(pointer)
	}
}

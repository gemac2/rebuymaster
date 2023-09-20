package models

import (
	"time"

	"github.com/adshao/go-binance/v2/futures"
	"github.com/gofrs/uuid"
	"github.com/hirokisan/bybit/v2"
)

type Order struct {
	ID                 uuid.UUID `db:"id" json:"id"`
	OrderType          string    `db:"order_type" json:"order_type"`
	CurrencyName       string    `db:"currency_name" json:"currency_name"`
	BuybackPercentage  float64   `db:"buyback_percentage" json:"buyback_percentage"`
	CurrencyPercentage float64   `db:"currency_percentage" json:"currency_percentage"`
	StopLoss           float64   `db:"stop_loss" json:"stop_loss"`
	OrderPrice         float64   `db:"order_price" json:"order_price"`
	CurrencyQuantity   float64   `db:"currency_quantity" json:"currency_quantity"`
	Leverage           int       `db:"leverage" json:"leverage"`
	CreatedAt          time.Time `db:"created_at" json:"created_at"`
	UpdatedAt          time.Time `db:"updated_at" json:"updated_at"`
	IsBuybacksEnabled  bool      `db:"is_buybacks_enabled" json:"is_buybacks_enabled"`
	IsOrderPosition    bool      `db:"is_order_position" json:"is_order_position"`
	StopLossTaken      bool      `db:"stop_loss_taken" json:"stop_loss_taken"`
	TakeProfitAchieved bool      `db:"take_profit_achieved" json:"take_profit_achieved"`
	TradeWon           bool      `db:"trade_won" json:"trade_won"`
	TradeLoss          bool      `db:"trade_loss" json:"trade_loss"`
	TakeProfit         float64   `db:"take_profit" json:"take_profit"`
	Profit             float64   `db:"profit" json:"profit"`
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

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
	// Return the symbol as a string since the constants may have changed
	return bybit.SymbolV5(currencyName)
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

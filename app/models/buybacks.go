package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
)

type Buyback struct {
	ID               uuid.UUID `db:"id" json:"id"`
	OrderID          uuid.UUID `db:"order_id" json:"order_id"`
	Price            float64   `db:"price" json:"price"`
	CurrencyAmount   float64   `db:"currency_amount" json:"currency_amount"`
	Margin           float64   `db:"margin" json:"margin"`
	MarginSum        float64   `db:"margin_sum" json:"margin_sum"`
	Average          float64   `db:"average" json:"average"`
	StopLossPrice    float64   `db:"stop_loss_price" json:"stop_loss_price"`
	Value            float64   `db:"value" json:"value"`
	TotalCurrency    float64   `db:"total_currency" json:"total_currency"`
	LiquidationPrice float64   `db:"liquidation_price" json:"liquidation_price"`
	ValueSum         float64   `db:"value_sum" json:"value_sum"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time `db:"updated_at" json:"updated_at"`
}

// Buybacks struct
type Buybacks []Buyback

func GenerateBuybacks(o Order) []Buyback {
	buybacks := make([]Buyback, 11)
	buybacksPercentage := o.BuybackPercentage
	if o.OrderType == "Long" {
		buybacksPercentage = -o.BuybackPercentage
	}

	if o.CurrencyPercentage == 100 {
		precioPrimeraRecompra := o.OrderPrice * ((buybacksPercentage / 100) + 1)
		buybacks[0] = Buyback{Price: o.OrderPrice, CurrencyAmount: o.CurrencyQuantity}
		buybacks[1] = Buyback{Price: precioPrimeraRecompra, CurrencyAmount: o.CurrencyQuantity}

		for i := 2; i < 11; i++ {
			buybacks[i].Price = buybacks[i-1].Price * (1 + buybacksPercentage/100)
			buybacks[i].CurrencyAmount = buybacks[i-1].CurrencyAmount + buybacks[i-1].CurrencyAmount*(o.CurrencyPercentage/100)
		}

		setMargin(buybacks, o.Leverage)
		getBuybackFinalPrice(buybacks)
		calculateStopLossPrice(o, buybacks)

		return buybacks
	}

	buybacks[0] = Buyback{Price: o.OrderPrice, CurrencyAmount: o.CurrencyQuantity}

	for i := 1; i < 11; i++ {
		buybacks[i].Price = buybacks[i-1].Price * (1 + buybacksPercentage/100)
		buybacks[i].CurrencyAmount = buybacks[i-1].CurrencyAmount + buybacks[i-1].CurrencyAmount*(o.CurrencyPercentage/100)
	}

	setMargin(buybacks, o.Leverage)
	getBuybackFinalPrice(buybacks)
	calculateStopLossPrice(o, buybacks)

	return buybacks
}

func FilterBuybacksByStopLoss(buybacks []Buyback, orderType string) []Buyback {
	var filteredBuybacks []Buyback
	for _, r := range buybacks {
		if orderType == "Short" {
			if r.Price < r.StopLossPrice {
				filteredBuybacks = append(filteredBuybacks, r)
			}
		}

		if orderType == "Long" {
			if r.Price > r.StopLossPrice {
				filteredBuybacks = append(filteredBuybacks, r)
			}
		}
	}
	return filteredBuybacks
}

func setMargin(buybacks []Buyback, apalancamiento int) {
	for i := range buybacks {
		buybacks[i].Margin = (buybacks[i].Price * buybacks[i].CurrencyAmount) / float64(apalancamiento)
	}

	buybacks[0].MarginSum = buybacks[0].Margin

	for i := 1; i < len(buybacks); i++ {
		buybacks[i].MarginSum = buybacks[i-1].MarginSum + buybacks[i].Margin
	}
}

func getBuybackFinalPrice(buybacks []Buyback) {
	buybacks[0].Value = buybacks[0].Price * buybacks[0].CurrencyAmount
	buybacks[0].TotalCurrency = buybacks[0].CurrencyAmount
	for i := 1; i < len(buybacks); i++ {
		buybacks[i].Value = buybacks[i].Price * buybacks[i].CurrencyAmount
		buybacks[i].TotalCurrency = buybacks[i].CurrencyAmount + buybacks[i-1].TotalCurrency
	}

	buybacks[0].ValueSum = buybacks[0].Value
	for i := 1; i < len(buybacks); i++ {
		buybacks[i].ValueSum = buybacks[i].Value + buybacks[i-1].ValueSum
	}

	buybacks[0].Average = buybacks[0].ValueSum / buybacks[0].TotalCurrency

	for i := 1; i < len(buybacks); i++ {
		buybacks[i].Average = (buybacks[i].Value + buybacks[i-1].ValueSum) / buybacks[i].TotalCurrency
	}
}

func calculateStopLossPrice(o Order, buybacks []Buyback) {
	liquidationPercentage := setLiquidationPercentage(o.Leverage)
	for i := range buybacks {
		posicion := (buybacks[i].Average * buybacks[i].TotalCurrency)

		if o.OrderType == "Short" {
			buybacks[i].LiquidationPrice = buybacks[i].Average * (1 + (float64(liquidationPercentage) / 100))
			buybacks[i].StopLossPrice = (posicion + o.StopLoss) / buybacks[i].TotalCurrency
		}

		if o.OrderType == "Long" {
			buybacks[i].LiquidationPrice = (buybacks[i].Average * ((float64(liquidationPercentage) / 100) - 1)) * (-1)
			buybacks[i].StopLossPrice = (posicion - o.StopLoss) / buybacks[i].TotalCurrency
		}
	}
}

func setLiquidationPercentage(leverage int) (liquidationPercentage float64) {
	switch leverage {
	case 5:
		liquidationPercentage = 20
	case 10:
		liquidationPercentage = 10
	case 20:
		liquidationPercentage = 5
	case 25:
		liquidationPercentage = 4
	case 50:
		liquidationPercentage = 2
	case 75:
		liquidationPercentage = 1.33
	case 100:
		liquidationPercentage = 1
	default:
		fmt.Println("\n Invalid option")
	}

	return liquidationPercentage
}

func SetPriceForExchanges(currencyName string, buybackPrice float64) string {
	price := ""
	switch currencyName {
	case "BTCUSDT":
		price = strconv.FormatFloat(buybackPrice, 'f', 1, 64)
	case "ETHUSDT":
		price = strconv.FormatFloat(buybackPrice, 'f', 2, 64)
	case "LPTUSDT":
		price = strconv.FormatFloat(buybackPrice, 'f', 3, 64)
	case "NMRUSDT":
		price = strconv.FormatFloat(buybackPrice, 'f', 1, 64)
	case "UNFIUSDT":
		price = strconv.FormatFloat(buybackPrice, 'f', 3, 64)
	case "BLZUSDT":
		price = strconv.FormatFloat(buybackPrice, 'f', 5, 64)
	case "PERPUSDT":
		price = strconv.FormatFloat(buybackPrice, 'f', 5, 64)
	case "TRBUSDT":
		price = strconv.FormatFloat(buybackPrice, 'f', 3, 64)
	case "HIFIUSDT":
		price = strconv.FormatFloat(buybackPrice, 'f', 5, 64)
	case "AMBUSDT":
		price = strconv.FormatFloat(buybackPrice, 'f', 6, 64)
	case "GALAUSDT":
		price = strconv.FormatFloat(buybackPrice, 'f', 5, 64)
	case "CHRUSDT":
		price = strconv.FormatFloat(buybackPrice, 'f', 4, 64)
	case "DOTUSDT":
		price = strconv.FormatFloat(buybackPrice, 'f', 3, 64)
	case "XRPUSDT":
		price = strconv.FormatFloat(buybackPrice, 'f', 4, 64)
	case "WUSDT":
		price = strconv.FormatFloat(buybackPrice, 'f', 3, 64)
	case "REEFUSDT":
		price = strconv.FormatFloat(buybackPrice, 'f', 6, 64)
	case "XVGUSDT":
		price = strconv.FormatFloat(buybackPrice, 'f', 6, 64)
	case "NEARUSDT":
		price = strconv.FormatFloat(buybackPrice, 'f', 3, 64)
	case "SAGAUSDT":
		price = strconv.FormatFloat(buybackPrice, 'f', 3, 64)
	default:
		price = strconv.FormatFloat(buybackPrice, 'f', 7, 64)
	}
	return price
}

package models

import (
	"time"

	"github.com/gofrs/uuid"
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
}

// Orders struct
type Orders []Order

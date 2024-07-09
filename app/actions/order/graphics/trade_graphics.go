package graphics

import (
	"bytes"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"rebuymaster/app/models"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/pkg/errors"
)

func GenerateGraphics(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	orders := []models.Order{}
	if err := tx.Where("is_order_position = true").All(&orders); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	chartWLRateBuffer, err := winLossRateGraphic(orders)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "GenerateGraphics - error when generate rate graphic"))
	}

	chartStrategyRateBuffer, err := strategyRateGraphics(tx, orders)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "GenerateGraphics - error when generate rate graphic"))
	}

	c.Set("chartRateContent", template.HTML(chartWLRateBuffer.String()))
	c.Set("chartStrategyContent", template.HTML(chartStrategyRateBuffer.String()))

	return c.Render(http.StatusOK, r.HTML("orders/graphics.plush.html"))
}

func winLossRateGraphic(orders []models.Order) (bytes.Buffer, error) {
	winCount := 0
	lossCount := 0
	for _, order := range orders {
		if order.TradeWon {
			winCount++
		} else if order.TradeLoss {
			lossCount++
		}
	}

	winColor := "#198754"
	lossColor := "#DC3545"

	pie := charts.NewPie()

	labelOpts := opts.Label{
		Show:      true,
		Formatter: "{b}: {c}",
	}

	tooltipOpts := opts.Tooltip{
		Show:      true,
		Formatter: "{a} <br/>{b}: {d}%",
	}

	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Trades Rate",
		}),

		charts.WithTooltipOpts(tooltipOpts),

		charts.WithInitializationOpts(opts.Initialization{
			Width:  "500px",
			Height: "300px",
		}),
	)

	winPercentage := setPercentageWinAndLoss(len(orders), winCount)
	lossPercentage := setPercentageWinAndLoss(len(orders), lossCount)

	pie.SetSeriesOptions(
		charts.WithLabelOpts(labelOpts),
	)

	pie.AddSeries("Trades",
		generatePieItems(winPercentage, lossPercentage, winColor, lossColor),
	)

	var chartBuffer bytes.Buffer
	if err := pie.Render(&chartBuffer); err != nil {
		return chartBuffer, err
	}

	return chartBuffer, nil
}

func strategyRateGraphics(tx *pop.Connection, orders []models.Order) (bytes.Buffer, error) {
	winTwoOneBuybackCount := 0
	lossTwoOneBuybackCount := 0
	winCalculatorCount := 0
	lossCalculatorCount := 0
	winTwoOneCount := 0
	lossTwoOneCount := 0
	for _, order := range orders {
		buybacks := []models.Buyback{}
		if err := tx.Where("order_id = ?", order.ID).All(&buybacks); err != nil {
			return bytes.Buffer{}, errors.WithStack(errors.Wrap(err, "strategyRateGraphics - error getting all buybacks"))
		}

		filterBuybacks := models.FilterBuybacksByStopLoss(buybacks, order.OrderType)

		if order.IsBuybacksEnabled {
			if len(filterBuybacks) == 1 {
				setCounts(order, winTwoOneBuybackCount, lossTwoOneBuybackCount)
			} else {
				setCounts(order, winCalculatorCount, lossCalculatorCount)
			}
		} else {
			setCounts(order, winTwoOneCount, lossTwoOneCount)
		}
	}

	fmt.Printf("LenOrders ==> %v\n", len(orders))

	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Trading Strategies",
		Subtitle: "It's extremely easy to use, right?",
	}))

	// Put data into instance
	bar.SetXAxis([]string{"Calculator", "two&one", "two&one with buyback"}).
		AddSeries("Win", generateBarItems(len(orders))).
		AddSeries("Loss", generateBarItems(len(orders)))

	var chartStrategyBuffer bytes.Buffer
	if err := bar.Render(&chartStrategyBuffer); err != nil {
		return chartStrategyBuffer, err
	}

	return chartStrategyBuffer, nil
}

func generatePieItems(winPercentage, lossPercentage float64, winColor, lossColor string) []opts.PieData {
	return []opts.PieData{
		{Value: winPercentage, Name: "Win", ItemStyle: &opts.ItemStyle{Color: winColor}},
		{Value: lossPercentage, Name: "Loss", ItemStyle: &opts.ItemStyle{Color: lossColor}},
	}
}

func setPercentageWinAndLoss(totalOrders, tradeTypeCount int) float64 {
	percentage := (float64(tradeTypeCount) / float64(totalOrders)) * 100
	return percentage
}

func setCounts(order models.Order, winCount, lossCount int) {
	if order.TradeWon {
		winCount++
	} else if order.TradeLoss {
		lossCount++
	}
}

func generateBarItems(lenOrders int) []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < 3; i++ {
		items = append(items, opts.BarData{Value: rand.Intn(lenOrders)})
	}
	return items
}

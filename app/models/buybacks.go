package models

import (
	"fmt"
)

type Recompra struct {
	Precio            float64
	Cantidad          float64
	Margen            float64
	SumaMargen        float64
	Promedio          float64
	PrecioStopLoss    float64
	Valor             float64
	TotalMonedas      float64
	PrecioLiquidacion float64
	SumaValor         float64
}

func GenerarRecompras(o Order) []Recompra {
	recompras := make([]Recompra, 11)
	porcentajeRecompras := o.BuybackPercentage
	if o.OrderType == "long" {
		porcentajeRecompras = -o.BuybackPercentage
	}

	if o.CurrencyPercentage == 100 {
		precioPrimeraRecompra := o.OrderPrice * ((porcentajeRecompras / 100) + 1)
		recompras[0] = Recompra{Precio: o.OrderPrice, Cantidad: o.CurrencyQuantity}
		recompras[1] = Recompra{Precio: precioPrimeraRecompra, Cantidad: o.CurrencyQuantity}

		for i := 2; i < 11; i++ {
			recompras[i].Precio = recompras[i-1].Precio * (1 + porcentajeRecompras/100)
			recompras[i].Cantidad = recompras[i-1].Cantidad + recompras[i-1].Cantidad*(o.CurrencyPercentage/100)
		}

		poblarMargen(recompras, o.Leverage)
		obtenerPrecioFinaldeCompras(recompras)
		calcularPrecioStopLoss(o, recompras)

		return recompras
	}

	recompras[0] = Recompra{Precio: o.OrderPrice, Cantidad: o.CurrencyQuantity}

	for i := 1; i < 11; i++ {
		recompras[i].Precio = recompras[i-1].Precio * (1 + porcentajeRecompras/100)
		recompras[i].Cantidad = recompras[i-1].Cantidad + recompras[i-1].Cantidad*(o.CurrencyPercentage/100)
	}

	poblarMargen(recompras, o.Leverage)
	obtenerPrecioFinaldeCompras(recompras)
	calcularPrecioStopLoss(o, recompras)

	return recompras
}

func ImprimirRecompras(recompras []Recompra, tipoDeCompra string) {
	for _, r := range recompras {
		if tipoDeCompra == "short" {
			if r.Precio < r.PrecioStopLoss {
				fmt.Printf("Precio: %.4f, Cantidad: %.4f, Valor:%.4f, TotalMonedas:%.4f, Promedio: %.4f, PrecioStopLoss: %.4f\n", r.Precio, r.Cantidad, r.Valor, r.TotalMonedas, r.Promedio, r.PrecioStopLoss)
			}
		}

		if tipoDeCompra == "long" {
			if r.Precio > r.PrecioStopLoss {
				fmt.Printf("Precio: %.4f, Cantidad: %.4f, Valor:%.4f, TotalMonedas:%.4f, Promedio: %.4f, PrecioStopLoss: %.4f\n", r.Precio, r.Cantidad, r.Valor, r.TotalMonedas, r.Promedio, r.PrecioStopLoss)
			}
		}
	}
	fmt.Println()
}

func poblarMargen(recompras []Recompra, apalancamiento int) {
	for i := range recompras {
		recompras[i].Margen = (recompras[i].Precio * recompras[i].Cantidad) / float64(apalancamiento)
	}

	recompras[0].SumaMargen = recompras[0].Margen

	for i := 1; i < len(recompras); i++ {
		recompras[i].SumaMargen = recompras[i-1].SumaMargen + recompras[i].Margen
	}
}

func obtenerPrecioFinaldeCompras(recompras []Recompra) {
	recompras[0].Valor = recompras[0].Precio * recompras[0].Cantidad
	recompras[0].TotalMonedas = recompras[0].Cantidad
	for i := 1; i < len(recompras); i++ {
		recompras[i].Valor = recompras[i].Precio * recompras[i].Cantidad
		recompras[i].TotalMonedas = recompras[i].Cantidad + recompras[i-1].TotalMonedas
	}

	recompras[0].SumaValor = recompras[0].Valor
	for i := 1; i < len(recompras); i++ {
		recompras[i].SumaValor = recompras[i].Valor + recompras[i-1].SumaValor
	}

	recompras[0].Promedio = recompras[0].SumaValor / recompras[0].TotalMonedas

	for i := 1; i < len(recompras); i++ {
		recompras[i].Promedio = (recompras[i].Valor + recompras[i-1].SumaValor) / recompras[i].TotalMonedas
	}
}

func calcularPrecioStopLoss(o Order, recompras []Recompra) {
	porcentajeLiquidacion := porcentajeLiquidacion(o.Leverage)
	for i := range recompras {
		posicion := (recompras[i].Promedio * recompras[i].TotalMonedas)

		if o.OrderType == "short" {
			recompras[i].PrecioLiquidacion = recompras[i].Promedio * (1 + (float64(porcentajeLiquidacion) / 100))
			recompras[i].PrecioStopLoss = (posicion + o.StopLoss) / recompras[i].TotalMonedas
		}

		if o.OrderType == "long" {
			recompras[i].PrecioLiquidacion = (recompras[i].Promedio * ((float64(porcentajeLiquidacion) / 100) - 1)) * (-1)
			recompras[i].PrecioStopLoss = (posicion - o.StopLoss) / recompras[i].TotalMonedas
		}
	}
}

func porcentajeLiquidacion(apalancamiento int) (porcentajeLiquidacion float64) {
	switch apalancamiento {
	case 5:
		porcentajeLiquidacion = 20
	case 10:
		porcentajeLiquidacion = 10
	case 20:
		porcentajeLiquidacion = 5
	case 50:
		porcentajeLiquidacion = 2
	case 75:
		porcentajeLiquidacion = 1.33
	case 100:
		porcentajeLiquidacion = 1
	default:
		fmt.Println("\nOpción inválida")
	}

	return porcentajeLiquidacion
}

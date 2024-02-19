package provider

import (
	"github.com/shopspring/decimal"
)

type RateMap map[string]map[string]decimal.Decimal

type Provider interface {
	Currencies() ([]string, error)
	Coins() ([]string, error)
	Rates(coins, currencies []string) (RateMap, error)
}

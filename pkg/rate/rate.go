package rate

import (
	"log"
	"sync"

	"github.com/ochenane/currency/pkg/provider"
	"github.com/shopspring/decimal"
)

const chunkSize = 1000

type Rate struct {
	coins      []string
	currencies []string
	data       map[string]decimal.Decimal
}

func New(coins, currencies []string) *Rate {
	return &Rate{
		coins:      coins,
		currencies: currencies,
		data:       make(map[string]decimal.Decimal, len(coins)*len(currencies)),
	}
}

func (r *Rate) Get(coin, currency string) (decimal.Decimal, bool) {
	result, ok := r.data[formatKey(coin, currency)]
	return result, ok
}

func (r *Rate) Update(p provider.Provider) {
	ch := make(chan provider.RateMap)
	var wg sync.WaitGroup
	for i := 0; i < len(r.coins); i += chunkSize {
		wg.Add(1)
		go func(i int) {
			data, err := p.Rates(r.coins[i:min(i+chunkSize, len(r.coins))], r.currencies)
			if err != nil {
				log.Println("cannot get rates:", err)
			}
			ch <- data
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for data := range ch {
		for coin, v := range data {
			for currency, price := range v {
				r.data[formatKey(coin, currency)] = price
			}
		}
	}
}

func formatKey(coin, rate string) string {
	return coin + ":" + rate
}

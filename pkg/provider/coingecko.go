package provider

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const baseUrl string = "https://api.coingecko.com/api/v3"

type coingecko struct {
	// FIXME: Can contain api key
}

type coingeckoCoin struct {
	Id     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

func NewCoingecko() Provider {
	return &coingecko{}
}

func (*coingecko) Coins() ([]string, error) {
	resp, err := http.Get(baseUrl + "/coins/list?include_platform=false")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code for coins: %d", resp.StatusCode)
	}

	coins := []coingeckoCoin{}
	if err := json.NewDecoder(resp.Body).Decode(&coins); err != nil {
		return nil, err
	}
	result := make([]string, 0, len(coins))
	for _, coin := range coins {
		result = append(result, coin.Id)
	}
	return result, nil
}

func (*coingecko) Currencies() ([]string, error) {
	resp, err := http.Get(baseUrl + "/simple/supported_vs_currencies")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code for currencies: %d", resp.StatusCode)
	}

	result := make([]string, 0)
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (*coingecko) Rates(coins, currencies []string) (RateMap, error) {
	q := url.Values{}
	q.Add("ids", strings.Join(coins, ","))
	q.Add("vs_currencies", strings.Join(currencies, ","))

	resp, err := http.Get(baseUrl + "/simple/price?" + q.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code for prices: %d", resp.StatusCode)
	}

	result := make(RateMap)
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

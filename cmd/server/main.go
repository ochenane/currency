package main

import (
	"fmt"
	"log"

	"github.com/ochenane/currency/pkg/provider"
	"github.com/ochenane/currency/pkg/rate"
)

// This is just a sample, this should be changed to a server
func main() {
	coingecko := provider.NewCoingecko()

	coins, err := coingecko.Coins()
	if err != nil {
		log.Fatal(err)
	}
	currencies, err := coingecko.Currencies()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: filter coins and/or currencies as needed, this just filters it to given list
	// removing it causes to get all coins
	coins = []string{"bitcoin", "ethereum", "solana", "ada", "atom", "tether"}

	rate := rate.New(coins, currencies)

	// TODO: run this as a job
	rate.Update(coingecko)

	// These are just for testing
	if p, ok := rate.Get("bitcoin", "usd"); ok {
		fmt.Println("Bitcoin in USD", p)
	}
	if p, ok := rate.Get("bitcoin", "eur"); ok {
		fmt.Println("Bitcoin in Euro", p)
	}
	if p, ok := rate.Get("ethereum", "usd"); ok {
		fmt.Println("Ethereum in USD", p)
	}
	if p, ok := rate.Get("solana", "usd"); ok {
		fmt.Println("Solana in USD", p)
	}
}

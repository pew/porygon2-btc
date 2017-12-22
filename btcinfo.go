package bot

import (
	"fmt"
	"encoding/json"
	"net/http"
	"log"
	"github.com/0x263b/porygon2"
)

type coins []struct {
	Id string
	Symbol string
	PriceUsd float64 `json:"price_usd,string"`
	PriceEur float64 `json:"price_eur,string"`
}

func btcInfo(command *bot.Cmd, matches []string) (msg string, err error) {
	resp, err := http.Get("https://api.coinmarketcap.com/v1/ticker/?convert=EUR")
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	var currencies coins

	if err := dec.Decode(&currencies); err != nil {
		log.Println(err)
		return "", err
	}

	var output string
	formatString := "%s: $%.2f / %.2f€ - "

	for _, v := range currencies {
		// we could do better here. 2018 (TM)
		if v.Id == "bitcoin" || v.Id == "ethereum" || v.Id == "monero" || v.Id == "bitcoin-cash" {
			output += fmt.Sprintf(formatString, v.Symbol, v.PriceUsd, v.PriceEur)
		}
	}

	// help
	output = output[:len(output)-3]
	return fmt.Sprintf(output), nil
}

func init() {
	bot.RegisterCommand("^coin$", btcInfo)
	log.Println("cryptocoin plugin loaded")
}
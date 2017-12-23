package bot

import (
	"fmt"
	"encoding/json"
	"net/http"
	"log"
	"github.com/0x263b/porygon2"
	"sort"
	"strings"
)

type coins []struct {
	Id string
	Symbol string
	Name string
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

	var output []string
	formatString := "%s: $%.2f / %.2f€"

	for _, v := range currencies {
		// we could do better here. 2018 (TM)
		if v.Name == "Bitcoin" || v.Name == "Bitcoin Cash" || v.Name == "Ethereum" || v.Name == "Monero" || v.Name == "IOTA" {
			output = append(output, fmt.Sprintf(formatString, v.Symbol, v.PriceUsd, v.PriceEur))
		}
	}

	sort.Strings(output)
	return fmt.Sprintf(strings.Join(output, " - ")), nil
}

func init() {
	bot.RegisterCommand("^coin$", btcInfo)
	log.Println("cryptocoin plugin loaded")
}
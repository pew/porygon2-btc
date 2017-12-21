package bot

import (
	"fmt"
	"encoding/json"
	"net/http"
	"log"
	"github.com/0x263b/porygon2"
)

type Currency struct {
	USD struct {
		Last float64
		Symbol string
	}
	EUR struct {
		Last float64
		Symbol string
	}
}

func btcInfo(command *bot.Cmd, matches []string) (msg string, err error) {
	resp, err := http.Get("https://blockchain.info/ticker")
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	var currency Currency

	if err := dec.Decode(&currency); err != nil {
		log.Println(err)
		return "", err
	}

	formatString := "1 BTC: %s %.2f - %s %.2f"

	return fmt.Sprintf(formatString, currency.USD.Symbol, currency.USD.Last, currency.EUR.Symbol, currency.EUR.Last), nil
}

func init() {
	bot.RegisterCommand("^btc$", btcInfo)
	log.Println("btc plugin loaded")
}
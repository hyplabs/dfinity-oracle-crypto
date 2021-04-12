package main

import (
	"os"
	"strings"
	"time"

	framework "github.com/hyplabs/dfinity-oracle-framework"
	"github.com/hyplabs/dfinity-oracle-framework/models"
)

func generateEndpoints(coinName string) []models.Endpoint {
	coinMarketCapAPIKey := os.Getenv("COINMARKETCAP_API_KEY")
	lowercaseCoinName := strings.ToLower(coinName)
	return []models.Endpoint{
		{
			Endpoint: "https://api.coingecko.com/api/v3/simple/price?" + lowercaseCoinName + "&vs_currencies=usd",
			JSONPaths: map[string]string{
				"usd_per_token": "$." + lowercaseCoinName + ".usd",
			},
		},
		{
			Endpoint: "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest?slug=" + lowercaseCoinName + "&CMC_PRO_API_KEY=" + coinMarketCapAPIKey,
			JSONPaths: map[string]string{
				"usd_per_token": "$.data..quote.USD.price",
			},
		},
	}
}

func main() {
	config := models.Config{
		CanisterName:   "crypto_oracle",
		UpdateInterval: 5 * time.Second,
	}

	engine := models.Engine{
		Metadata: []models.MappingMetadata{
			{Key: "Bitcoin", Endpoints: generateEndpoints("Bitcoin")},
			{Key: "Ethereum", Endpoints: generateEndpoints("Ethereum")},
			{Key: "Cardano", Endpoints: generateEndpoints("Cardano")},
			{Key: "Polkadot", Endpoints: generateEndpoints("Polkadot")},
			{Key: "Uniswap", Endpoints: generateEndpoints("Uniswap")},
			{Key: "Litecoin", Endpoints: generateEndpoints("Litecoin")},
			{Key: "Chainlink", Endpoints: generateEndpoints("Chainlink")},
			{Key: "Stellar", Endpoints: generateEndpoints("Stellar")},
			{Key: "Filecoin", Endpoints: generateEndpoints("Filecoin")},
			{Key: "TRON", Endpoints: generateEndpoints("TRON")},
			{Key: "Dogecoin", Endpoints: generateEndpoints("Dogecoin")},
			{Key: "Solana", Endpoints: generateEndpoints("Solana")},
			{Key: "EOS", Endpoints: generateEndpoints("EOS")},
			{Key: "Monero", Endpoints: generateEndpoints("Monero")},
			{Key: "Terra", Endpoints: generateEndpoints("Terra")},
			{Key: "IOTA", Endpoints: generateEndpoints("IOTA")},
			{Key: "Cosmos", Endpoints: generateEndpoints("Cosmos")},
			{Key: "Algorand", Endpoints: generateEndpoints("Algorand")},
			{Key: "Tezos", Endpoints: generateEndpoints("Tezos")},
			{Key: "Avalanche", Endpoints: generateEndpoints("Avalanche")},
		},
	}

	oracle := framework.NewOracle(&config, &engine)
	oracle.Bootstrap()
	oracle.Run()
}

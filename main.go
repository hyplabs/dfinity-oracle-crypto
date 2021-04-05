package main

import (
	"time"

	framework "github.com/hyplabs/dfinity-oracle-framework"
	"github.com/hyplabs/dfinity-oracle-framework/models"
)

func main() {
	config := models.Config{
		CanisterName:   "crypto_oracle",
		UpdateInterval: 5 * time.Second,
	}

	engine := models.Engine{
		Metadata: []models.MappingMetadata{
			{
				Key: "Bitcoin",
				Endpoints: []models.Endpoint{
					{
						Endpoint:  "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd",
						JSONPaths: map[string]string{"usd_per_token": "$.bitcoin.usd"},
					},
				},
			},
			{
				Key: "Ethereum",
				Endpoints: []models.Endpoint{
					{
						Endpoint:  "https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd",
						JSONPaths: map[string]string{"usd_per_token": "$.ethereum.usd"},
					},
				},
			},
		},
	}

	oracle := framework.NewOracle(&config, &engine)
	oracle.Bootstrap()
	oracle.Run()
}

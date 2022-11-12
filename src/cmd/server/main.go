package main

import (
	"github.com/misikch/stocktickers/src/internal/app"
	"github.com/misikch/stocktickers/src/internal/config"
)

func main() {
	cfg := config.InitConfig()

	err := app.Run(cfg)
	if err != nil {
		panic(err)
	}
}

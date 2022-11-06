package main

import (
	"github.com/IMDb-searcher/internal/cli"
	"github.com/IMDb-searcher/internal/logger"
)

func main() {
	log := logger.GetLogger()

	ui, err := cli.New(log)
	if err != nil {
		log.Panic(err)
	}
	ui.Run()
}

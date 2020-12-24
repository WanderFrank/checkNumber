package main

import (
	"gotest/app"
	"gotest/game"
	// "gotest/io/telegram"
	"gotest/lang"
	"gotest/io/console"
	"log"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	t := console.New()

	g := game.New()

	langData, err := lang.LoadLangData(lang.Ru)
	if err != nil {
		return err
	}

	app := app.New(g, t, langData)
	err = app.Run()
	if err != nil {
		return err
	}

	return nil
}

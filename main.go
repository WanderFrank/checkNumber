package main

import (
	"gotest/app"
	"gotest/game"
	"gotest/game/telegram"
)

func main() {
	const (
		token = "1458054360:AAEmSNc1AtgHvVPhhLY0mkpIPSygiaRIHPM"
		chatID = 302066235
	)
	g := game.New()
	t := telegram.New(token, chatID)

	app := app.New(g, t, t)
	app.Run()
}

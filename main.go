package main

import (
	"gotest/app"
	"gotest/game"
	"gotest/game/telegram"
	// "gotest/game/console"
	"gotest/app/lang"
)

func main() {
	
	const (
		token = "1458054360:AAEmSNc1AtgHvVPhhLY0mkpIPSygiaRIHPM"
		chatID = 302066235
	)
	t := telegram.New(token, chatID)
	//t := console.New()
	
	g := game.New()
	
	langData := lang.GetLangData(lang.Ru)

	app := app.New(g, t, langData)
	app.Run()
}

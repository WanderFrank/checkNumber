package main

import (
	"gotest/app"
	"gotest/game"
	"gotest/game/input/console"
	"gotest/game/viewer/factory"
)

func main() {
	g := game.New()
	v := factory.NewByLanguage(factory.Ru)
	c := console.New()

	app := app.New(g, v, c)
	app.Run()
}

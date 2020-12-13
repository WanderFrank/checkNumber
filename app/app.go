package app

import (
	"gotest/game"
	"gotest/game/io"
)

type App interface {
	Run()
}

type app struct {
	game   game.Game
	viewer io.Viewer
	input  io.InputProvider
}

func New(game game.Game, viewer io.Viewer, input io.InputProvider) *app {
	return &app{
		game:   game,
		viewer: viewer,
		input:  input,
	}
}

func (a *app) Run() {
	var inputChannel chan int = make(chan int)

	a.viewer.PrintTitle()

	go func() {
		for {
			inputChannel <- a.input.GetNumber()
		}
	}()

	for {
		a.viewer.PrintInputPrompt(a.game.CheckCount())
		number := <-inputChannel
		checkResult := a.game.Check(number)

		a.viewer.PrintCheckResult(toViewCheckResult(checkResult))
		if checkResult == game.Equal {
			break
		}
	}
}

func toViewCheckResult(value game.CheckResult) io.CheckResult {
	toResultMap := map[game.CheckResult]io.CheckResult{
		game.Equal: io.Equal,
		game.More:  io.More,
		game.Less:  io.Less,
	}

	return toResultMap[value]
}

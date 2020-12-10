package app

import (
	"gotest/game"
	"gotest/game/input"
	"gotest/game/viewer"
)

type App interface {
	Run()
}

type app struct {
	game   game.Game
	viewer viewer.Viewer
	input  input.InputProvider
}

func New(game game.Game, viewer viewer.Viewer, input input.InputProvider) *app {
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

func toViewCheckResult(value game.CheckResult) viewer.CheckResult {
	toResultMap := map[game.CheckResult]viewer.CheckResult{
		game.Equal: viewer.Equal,
		game.More:  viewer.More,
		game.Less:  viewer.Less,
	}

	return toResultMap[value]
}

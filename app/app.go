package app

import (
	"gotest/game"
	"gotest/game/io"
	"fmt"
)

type App interface {
	Run()
}

type app struct {
	game   game.Game
	transport io.Transport
	langData LanguageData
}

type LanguageData struct {
	Title       string
	InputNumber string
	Less        string
	More        string
	Equal       string
}

func New(game game.Game, transport io.Transport, langData LanguageData) *app {
	return &app{
		game:   game,
		transport: transport,
		langData:  langData,
	}
}

func (a *app) Run() {
	var inputChannel chan int = make(chan int)

	a.transport.Print(a.langData.Title)

	go func() {
		for {
			inputChannel <- a.transport.GetNumber()
		}
	}()

	for {
		a.transport.Print(fmt.Sprintf("%d) "+a.langData.InputNumber, a.game.CheckCount()+1))
		number := <-inputChannel
		checkResult := a.game.Check(number)

		a.printCheckResult(checkResult)

		if checkResult == game.Equal {
			break
		}
	}
}

func (a *app) printCheckResult(result game.CheckResult) {
	var message string
	if result == game.Less {
		message = a.langData.Less
	} else if result == game.More {
		message = a.langData.More
	} else if result == game.Equal {
		message = a.langData.Equal
	}

	a.transport.Print(message)
}

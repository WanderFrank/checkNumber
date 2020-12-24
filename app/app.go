package app

import (
	"gotest/game"
	"gotest/io"
	"fmt"
	"gotest/lang"
)

type app struct {
	game  game.Game
	transport io.Transport
	langData lang.LanguageData
}

func New(game game.Game, transport io.Transport, langData lang.LanguageData) *app {
	return &app{
		game:   game,
		transport: transport,
		langData:  langData,
	}
}

func (a *app) Run() (err error) {
	defer func () {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	inputChannel := make(chan int)
	inputErrorChannel := make(chan error)

	defer close(inputChannel)
	defer close(inputErrorChannel)
	
	go a.inputLoop(inputChannel, inputErrorChannel)

	a.print(a.langData.Title)

	for {
		a.print(fmt.Sprintf("%d) "+a.langData.InputNumber, a.game.CheckCount()+1))

		select {
		case number := <- inputChannel:
			checkResult := a.game.Check(number)

			a.printCheckResult(checkResult)

			if checkResult == game.Equal {
				break
			}

		case err := <- inputErrorChannel:
			a.print(err.Error())
		}

		
	}

	return
}

func (a *app) print(text string) {
	err := a.transport.Print(text)
	if err != nil {
		panic(err)
	}
}

func (a *app) inputLoop(resultChannel chan int, errorChannel chan error) {
	for {
		result, err := a.transport.GetNumber()
		if err != nil {
			
			errorChannel <- err
			continue
		}

		resultChannel <- result
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

	a.print(message)
}

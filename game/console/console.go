package console

import (
	"fmt"
	"gotest/game/viewer"
)

type LanguageData struct {
	Title       string
	InputNumber string
	Less        string
	More        string
	Equal       string
}

type consoleViewer struct {
	langData LanguageData
}

func New(lang LanguageData) *consoleViewer {
	return &consoleViewer{langData: lang}
}

func (v *consoleViewer) PrintInputPrompt(currentCheckCount int) {
	fmt.Printf("%d) "+v.langData.InputNumber, currentCheckCount+1)
}

func (v *consoleViewer) PrintTitle() {
	fmt.Println(v.langData.Title)
}

func (v *consoleViewer) PrintCheckResult(result viewer.CheckResult) {
	if result == viewer.Less {
		fmt.Println(v.langData.Less)
	} else if result == viewer.More {
		fmt.Println(v.langData.More)
	} else if result == viewer.Equal {
		fmt.Println(v.langData.Equal)
	}
}

func (c *consoleViewer) GetNumber() int {
	var number int
	fmt.Scanf("%d\n", &number)

	return number
}

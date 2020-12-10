package console

import (
	"fmt"
	"gotest/game/viewer"
)

type consoleViewer struct {
	langData viewer.LanguageData
}

func New(lang viewer.LanguageData) *consoleViewer {
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

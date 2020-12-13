package en

import "gotest/game/viewer/console"

// Get data
func Get() console.LanguageData {
	return console.LanguageData{
		Title:       "Tell number!",
		InputNumber: "Input number:",
		Less:        "Less",
		More:        "More",
		Equal:       "Equal",
	}
}

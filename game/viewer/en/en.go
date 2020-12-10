package en

import "gotest/game/viewer"

// Get data
func Get() viewer.LanguageData {
	return viewer.LanguageData{
		Title:       "Tell number!",
		InputNumber: "Input number:",
		Less:        "Less",
		More:        "More",
		Equal:       "Equal",
	}
}

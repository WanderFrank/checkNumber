package en

import "gotest/app"

// Get data
func Get() app.LanguageData {
	return app.LanguageData{
		Title:       "Tell number!",
		InputNumber: "Input number:",
		Less:        "Less",
		More:        "More",
		Equal:       "Equal",
	}
}

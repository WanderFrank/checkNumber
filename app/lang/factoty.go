package lang

import (
	"gotest/app"
	"gotest/app/lang/ru"
	"gotest/app/lang/en"
)

type language string
const (
	Ru language = "ru"
	En language = "en"
)

func GetLangData(lang language) app.LanguageData {
	var langData app.LanguageData
	switch lang {
	case Ru:
		langData = ru.Get()
	case En:
		langData = en.Get()
	default:
		langData = ru.Get()
	}

	return langData
}
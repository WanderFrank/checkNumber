package factory

import (
	"gotest/game/viewer"
	"gotest/game/viewer/console"
	"gotest/game/viewer/en"
	"gotest/game/viewer/ru"
)

type language string

const (
	Ru language = "ru"
	En language = "en"
)

func NewByLanguage(lang language) viewer.Viewer {
	var langData viewer.LanguageData
	switch lang {
	case Ru:
		langData = ru.Get()
	case En:
		langData = en.Get()
	default:
		langData = ru.Get()
	}

	return console.New(langData)
}

package lang

import (
	"encoding/json"
	"io/ioutil"
)

type language string

const (
	Ru language = "ru"
	En language = "en"
)

func LoadLangData(lang language) (LanguageData, error) {
	var langData LanguageData

	file, err := ioutil.ReadFile("lang/data/" + string(lang) + ".json")
	if err != nil {
		return langData, err
	}

	err = json.Unmarshal(file, &langData)

	return langData, err
}

package ru

import "gotest/game/viewer/console"

func Get() console.LanguageData {
	return console.LanguageData{
		Title:       "Загадано число от 1 до 100. Угдай!",
		InputNumber: "Введите число:",
		Less:        "Загаданное число меньше",
		More:        "Загаданное число больше",
		Equal:       "Вы угадали",
	}
}

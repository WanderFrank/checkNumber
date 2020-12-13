package ru

import "gotest/app"

func Get() app.LanguageData {
	return app.LanguageData{
		Title:       "Загадано число от 1 до 100. Угдай!",
		InputNumber: "Введите число:",
		Less:        "Загаданное число меньше",
		More:        "Загаданное число больше",
		Equal:       "Вы угадали",
	}
}

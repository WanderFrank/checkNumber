package telegram

import (
	"gotest/game/io"
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"strconv"
)

type telegramViewer struct {
	bot *tgbotapi.BotAPI
	chatId int64
	updatesChannel tgbotapi.UpdatesChannel
}

func New(token string, chatId int64) *telegramViewer {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	// инициализируем канал, куда будут прилетать обновления от API
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	updatesChannel, err := bot.GetUpdatesChan(ucfg)
	if err != nil {
		log.Panic(err)
	}

	return &telegramViewer{
		bot: bot,
		chatId: chatId,
		updatesChannel: updatesChannel,
	}
}

func (v *telegramViewer) PrintInputPrompt(currentCheckCount int) {
	v.sendMessage("Введите число")
}

func (v *telegramViewer) PrintTitle() {
	v.sendMessage("Загадано число от 1 до 100. Угадай!")
}
func (v *telegramViewer) PrintCheckResult(result io.CheckResult) {
	if result == io.Less {
		v.sendMessage("Меньше")
	} else if result == io.More {
		v.sendMessage("Больше")
	} else if result == io.Equal {
		v.sendMessage("Угадали!")
	}
}

func (v *telegramViewer) sendMessage(message string) {
	msg := tgbotapi.NewMessage(v.chatId, message)
	// и отправляем его
	v.bot.Send(msg)
}

func (i *telegramViewer) GetNumber() int {
	update := <- i.updatesChannel
	text := update.Message.Text

	number, _ := strconv.Atoi(text)

	return number
}
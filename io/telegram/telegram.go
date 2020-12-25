package telegram

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"strconv"
)

type telegramViewer struct {
	bot            *tgbotapi.BotAPI
	chatId         int64
	updatesChannel tgbotapi.UpdatesChannel
}

func New(token string, chatId int64) (*telegramViewer, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	// инициализируем канал, куда будут прилетать обновления от API
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	updatesChannel, err := bot.GetUpdatesChan(ucfg)
	if err != nil {
		return nil, err
	}

	return &telegramViewer{
			bot:            bot,
			chatId:         chatId,
			updatesChannel: updatesChannel,
		},
		nil
}

func (v *telegramViewer) Print(message string) error {
	msg := tgbotapi.NewMessage(v.chatId, message)
	// и отправляем его
	_, err := v.bot.Send(msg)

	return err

}

func (i *telegramViewer) GetNumber() (int, error) {
	update := <-i.updatesChannel
	text := update.Message.Text

	number, err := strconv.Atoi(text)

	return number, err
}

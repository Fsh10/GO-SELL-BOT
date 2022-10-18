package bot

import (
	"fmt"
	"go.uber.org/zap"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type replyKeyboardValue string

const (
	ReplyCategories = replyKeyboardValue("Categories")
	ReplyReviews    = replyKeyboardValue("Reviews")
	ReplyHelp       = replyKeyboardValue("Help")
)

const (
	inviteUrl = "https://t.me/+1-lMGCQ7zOphODUy"
)

func (b *bot) StartCmd(upd tgbotapi.Update) {
	name := upd.Message.From.UserName
	if name == "" {
		name = upd.Message.From.FirstName
	}
	message := `
Welcome to <b>Xbox Store | Bot-shop for subscriptions and games</b>, %s!

Don't forget to subscribe to <a href='%s'>our main channel</a> and follow the news about game releases and promotions.
`
	reply := tgbotapi.NewMessage(upd.Message.Chat.ID, fmt.Sprintf(message, name, inviteUrl))
	reply.ParseMode = "html"

	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(string(ReplyCategories)),
			tgbotapi.NewKeyboardButton(string(ReplyReviews)),
			tgbotapi.NewKeyboardButton(string(ReplyHelp)),
		),
	)
	reply.ReplyMarkup = keyboard
	reply.DisableWebPagePreview = true

	if err := b.apiRequest(reply); err != nil {
		b.logger.Error("failed to send start message", zap.Error(err))
	}
}

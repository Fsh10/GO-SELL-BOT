package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

func (b *bot) HelpCmd(upd tgbotapi.Update) {
	message := `
💬 <b>Support:</b> @noobmaster111, write if you have any problems.
⌚ <b>Online:</b> Approximately from 10:00 - 00:00 MSK.
❗I don't buy anything and don't take on consignment, there is no advertising in the bot.
	`
	reply := tgbotapi.NewMessage(upd.Message.Chat.ID, message)
	reply.ParseMode = "html"

	if err := b.apiRequest(reply); err != nil {
		b.logger.Error("failed to send help message", zap.Error(err))
	}
}

package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type (
	commandEntity struct {
		key    commandKey
		desc   string
		action func(upd tgbotapi.Update)
	}
)

type commandKey string

const (
	StartCmdKey   = commandKey("start")
	ShopCmdKey    = commandKey("shop")
	ReviewsCmdKey = commandKey("reviews")
	HelpCmdKey    = commandKey("help")
)

func (b *bot) initCommands() error {
	commands := []commandEntity{
		{
			key:    StartCmdKey,
			desc:   "Start the bot",
			action: b.StartCmd,
		},
		{
			key:    ShopCmdKey,
			desc:   "Show products",
			action: b.ShopCmd,
		},
		{
			key:    ReviewsCmdKey,
			desc:   "Customer reviews",
			action: b.ReviewsCmd,
		},
		{
			key:    HelpCmdKey,
			desc:   "Support",
			action: b.HelpCmd,
		},
	}

	tgCommands := make([]tgbotapi.BotCommand, 0, len(commands))
	for _, cmd := range commands {
		b.commands[cmd.key] = cmd
		tgCommands = append(tgCommands, tgbotapi.BotCommand{
			Command:     "/" + string(cmd.key),
			Description: cmd.desc,
		})
	}

	config := tgbotapi.NewSetMyCommands(tgCommands...)
	return b.apiRequest(config)
}

func (b *bot) replyToCommand(text string) (commandEntity, bool) {
	switch replyKeyboardValue(text) {
	case ReplyCategories:
		cmd, ok := b.commands[ShopCmdKey]
		return cmd, ok
	case ReplyReviews:
		cmd, ok := b.commands[ReviewsCmdKey]
		return cmd, ok
	case ReplyHelp:
		cmd, ok := b.commands[HelpCmdKey]
		return cmd, ok
	}

	return commandEntity{}, false
}

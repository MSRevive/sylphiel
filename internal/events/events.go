package events

import (
	//"fmt"

	"github.com/msrevive/sylphiel/cmd/dbot"
	//"github.com/msrevive/sylphiel/internal/response"

	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/bot"
	//"github.com/disgoorg/disgo/discord"
)

func OnReady(b *dbot.Bot) bot.EventListener {
	/*return bot.NewListenerFunc(func(e *events.Ready) {
		e.Client().Rest().CreateMessage(b.Config.Discord.DevChannel, discord.NewMessageCreateBuilder().SetContent("Bot loaded.").Build())
	})*/
	return bot.NewListenerFunc(func(e *events.Ready) {
		b.Logger.Debug("OnReady event called")
	})
}
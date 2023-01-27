package events

import (
	"fmt"

	"github.com/msrevive/sylphiel/cmd/dbot"

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

func GuildMemberJoin(b *dbot.Bot) bot.EventListener {
	return bot.NewListenerFunc(func(e *events.GuildMemberJoin) {
		b.Logger.Debug("GuildMemberJoin event called")
		if e.GuildID == b.Config.Discord.GuildID {
			if b.Config.Discord.DefaultRole != 0 && !e.Member.User.Bot {
				e.Client().Rest().AddMemberRole(e.GuildID, e.Member.User.ID, b.Config.Discord.DefaultRole)
				fmt.Println("AddMemberGuild")
			}
		}
	})
}
package events

import (
	//"fmt"
	"context"

	"github.com/msrevive/sylphiel/cmd/dbot"
	//"github.com/msrevive/sylphiel/internal/response"

	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/gateway"
)

func OnReady(b *dbot.Bot) bot.EventListener {
	// return bot.NewListenerFunc(func(e *events.Ready) {
	// 	e.Client().Rest().CreateMessage(b.Config.Disc.DevChannel, discord.NewMessageCreateBuilder().SetContent("Bot loaded.").Build())
	// })
	return bot.NewListenerFunc(func(e *events.Ready) {
		if err := b.Client.SetPresence(context.TODO(),
			gateway.WithListeningActivity("gobby slobby"),
			gateway.WithOnlineStatus(discord.OnlineStatusOnline),
		); err != nil {
			b.Logger.Errorf("Failed to set presence: %s", err)
		}
	})
}

func ReactionAdd(b *dbot.Bot) bot.EventListener {
	return bot.NewListenerFunc(func(e *events.GuildMessageReactionAdd) {
		if e.UserID == b.Client.ID() {
			return
		}

		if e.ChannelID == b.Config.Roles.RoleChannel {
			if *e.Emoji.Name == "🎙️" {
				e.Client().Rest().AddMemberRole(e.GuildID, e.UserID, b.Config.Roles.VARole)
			}
		}
	})
}
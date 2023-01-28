package events

import (
	//"fmt"

	"github.com/msrevive/sylphiel/cmd/dbot"
	"github.com/msrevive/sylphiel/internal/response"

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
		if e.GuildID == b.Config.Disc.GuildID {
			if b.Config.Disc.DefaultRole != 0 && !e.Member.User.Bot {
				e.Client().Rest().AddMemberRole(e.GuildID, e.Member.User.ID, b.Config.Disc.DefaultRole)

				err := response.AuditMemberJoined(b.Webhook, e)
				if err != nil {
					b.Logger.Error(err)
				}
			}
		}
	})
}

func GuildMemberLeave(b *dbot.Bot) bot.EventListener {
	return bot.NewListenerFunc(func(e *events.GuildMemberLeave) {
		b.Logger.Debug("GuildMemberLeave event called")
		if e.GuildID == b.Config.Disc.GuildID {
			if !e.Member.User.Bot {
				err := response.AuditMemberLeft(b.Webhook, e)
				if err != nil {
					b.Logger.Error(err)
				}
			}
		}
	})
}

func GuildVoiceJoin(b *dbot.Bot) bot.EventListener {
	return bot.NewListenerFunc(func(e *events.GuildVoiceJoin) {
		b.Logger.Debug("GuildVoiceJoin event called")
		err := response.AuditVoiceJoined(b.Webhook, e)
		if err != nil {
			b.Logger.Error(err)
		}
	})
}

func GuildVoiceLeave(b *dbot.Bot) bot.EventListener {
	return bot.NewListenerFunc(func(e *events.GuildVoiceLeave) {
		b.Logger.Debug("GuildVoiceLeave event called")
		err := response.AuditVoiceLeft(b.Webhook, e)
		if err != nil {
			b.Logger.Error(err)
		}
	})
}

func MessageDelete(b *dbot.Bot) bot.EventListener {
	return bot.NewListenerFunc(func(e *events.MessageDelete) {
		b.Logger.Debug("MessageDelete event called")
		err := response.AuditMessageDelete(b.Webhook, e)
		if err != nil {
			b.Logger.Error(err)
		}
	})
}

func MessageUpdate(b *dbot.Bot) bot.EventListener {
	return bot.NewListenerFunc(func(e *events.MessageUpdate) {
		b.Logger.Debug("MessageUpdate event called")
		err := response.AuditMessageUpdate(b.Webhook, e)
		if err != nil {
			b.Logger.Error(err)
		}
	})
}
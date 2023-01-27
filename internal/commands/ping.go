package commands

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var ping = discord.SlashCommandCreate {
	Name: "ping",
	Description: "Ping command",
	Options: []discord.ApplicationCommandOption {
		discord.ApplicationCommandOptionBool{
			Name: "ephemeral",
			Description: "If the response should only be visible to you",
			Required: true,
		},
	},
}

func PingHandler(e *handler.CommandEvent) error {
	return e.CreateMessage(discord.NewMessageCreateBuilder().
		SetContent("pong!").
		Build(),
	)
}
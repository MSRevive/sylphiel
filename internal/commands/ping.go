package commands

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var ping = discord.SlashCommandCreate {
	Name: "ping",
	Description: "Ping command",
}

func HandlePing(e *handler.CommandEvent) error {
	var gatewayPing string
	if e.Client().HasGateway() {
		gatewayPing = e.Client().Gateway().Latency().String()
	}

	return e.CreateMessage(discord.NewMessageCreateBuilder().
		SetEphemeral(true).
		SetContentf("pong!, responded in %s", gatewayPing).
		Build(),
	)
}
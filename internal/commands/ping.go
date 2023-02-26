package commands

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var ping = discord.SlashCommandCreate {
	Name: "ping",
	Description: "Ping command",
}

func PingHandler(e *handler.CommandEvent) error {
	var gatewayPing string
	if e.Client().HasGateway() {
		gatewayPing = e.Client().Gateway().Latency().String()
	}

	

	embed := discord.NewEmbedBuilder().
		SetTitle("Pong!").
		SetColor(0x009999).
		SetTimestamp(time.Now()).
		AddField("Gateway", gatewayPing, false).
		AddField("FuzzNet", gatewayPing, false).
		SetFooter("Sylphiel", "https://winterfang.com/assets/gfx/bot-avatar.png").
	Build()

	return e.CreateMessage(discord.NewMessageCreateBuilder().
		SetEphemeral(true).
		SetContentf("pong!, responded in %s", gatewayPing).
		Build(),
	)
}
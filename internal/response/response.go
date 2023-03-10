package response

import (
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

func NotImplemented(e *handler.CommandEvent) error {
	embed := discord.NewEmbedBuilder().
		SetTitle("501 Not Implemented").
		SetColor(0xcc0000).
		SetDescription("Command not yet implemented.").
		SetTimestamp(time.Now()).
		SetFooter("Sylphiel", "https://winterfang.com/assets/gfx/bot-avatar.png").
	Build()

	return e.Respond(discord.InteractionResponseTypeCreateMessage, discord.NewMessageCreateBuilder().
		SetEmbeds(embed).
		Build(),
	)
}
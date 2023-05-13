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

func Error(e *handler.CommandEvent, err error) error {
	embed := discord.NewEmbedBuilder().
		SetTitle("500 Internal Server Error").
		SetColor(0xcc0000).
		SetDescription(err.Error()).
		SetTimestamp(time.Now()).
		SetFooter("Sylphiel", "https://winterfang.com/assets/gfx/bot-avatar.png").
	Build()

	return e.Respond(discord.InteractionResponseTypeCreateMessage, discord.NewMessageCreateBuilder().
		SetEphemeral(true).
		SetEmbeds(embed).
		Build(),
	)
}

func NoPermission(e *handler.CommandEvent) error {
	embed := discord.NewEmbedBuilder().
		SetTitle("Incorrect Permission").
		SetColor(0xcc0000).
		SetDescription("You don't have permission to this command!").
		SetTimestamp(time.Now()).
		SetFooter("Sylphiel", "https://winterfang.com/assets/gfx/bot-avatar.png").
	Build()

	return e.Respond(discord.InteractionResponseTypeCreateMessage, discord.NewMessageCreateBuilder().
		SetEphemeral(true).
		SetEmbeds(embed).
		Build(),
	)
}

func Success(e *handler.CommandEvent, msg string, emp bool) error {
	embed := discord.NewEmbedBuilder().
		SetTitle("Command Successful").
		SetColor(0x009999).
		SetDescription(msg).
		SetTimestamp(time.Now()).
		SetFooter("Sylphiel", "https://winterfang.com/assets/gfx/bot-avatar.png").
	Build()

	return e.Respond(discord.InteractionResponseTypeCreateMessage, discord.NewMessageCreateBuilder().
		SetEphemeral(emp).
		SetEmbeds(embed).
		Build(),
	)
}
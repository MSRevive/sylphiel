package response

import (
	"github.com/disgoorg/disgo/webhook"
	"github.com/disgoorg/disgo/discord"
)

func AuditMemberJoined(c webhook.Client) error {
	c.CreateEmbed(discord.NewEmbedBuilder().
		SetDescription("hello world!").
		Build(),
	)
}
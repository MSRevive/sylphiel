package commands

import (
	"github.com/msrevive/sylphiel/cmd/dbot"
	"github.com/msrevive/sylphiel/internal/response"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var setup = discord.SlashCommandCreate {
	Name: "setup",
	Description: "Setup stuffs",
	Options: []discord.ApplicationCommandOption{
		discord.ApplicationCommandOptionSubCommand{
			Name: "roles",
			Description: "Setup ping roles.",
		},
		discord.ApplicationCommandOptionSubCommand{
			Name: "serverlist",
			Description: "Setup server list.",
		},
	},
}

func HandleRolesSetup(b *dbot.Bot) handler.CommandHandler {
	return func(e *handler.CommandEvent) error {
		if e.Member().Permissions.Missing(discord.PermissionAdministrator) {
			return response.NoPermission(e)
		}

		embed := discord.NewEmbedBuilder().
			SetColor(0x7851a9).
			AddField(":studio_microphone: - Voice Acting", "", false).
			AddField(":hammer: - Mapping", "", false).
			AddField("<:blender:1106808241633120317> - Modeling/Animation", "", false).
		Build()

		msgBuilder := discord.NewMessageCreateBuilder().
			SetContent("React to this message if you're interested in helping us.").
			SetEmbeds(embed).
		Build()

		msg, err := e.Client().Rest().CreateMessage(e.ChannelID(), msgBuilder)
		if err != nil {
			b.Logger.Error(err)
			return response.Error(e, err)
		}

		if err := e.Client().Rest().AddReaction(e.ChannelID(), msg.ID, ":hammer:"); err != nil {
			b.Logger.Error(err)
			return response.Error(e, err)
		}

		return nil
	}
}

func HandleServerListSetup(e *handler.CommandEvent) error {
	return response.NotImplemented(e)
}
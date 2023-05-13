package commands

import (
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

func HandleRolesSetup(e *handler.CommandEvent) error {
	return response.NotImplemented(e)
}

func HandleServerListSetup(e *handler.CommandEvent) error {
	return response.NotImplemented(e)
}
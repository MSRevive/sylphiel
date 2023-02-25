package commands

import (
	"github.com/msrevive/sylphiel/internal/response"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var restore = discord.SlashCommandCreate {
	Name: "restore",
	Description: "Restore a recently deleted character.",
	Options: []discord.ApplicationCommandOption {
		discord.ApplicationCommandOptionString{
			Name: "steamid64",
			Description: "Player's SteamID 64 to restore.",
			Required: true,
		},
		discord.ApplicationCommandOptionInt{
			Name: "slot",
			Description: "Player's character slot.",
			Required: true,
			Choices: []discord.ApplicationCommandOptionChoiceInt{
				discord.ApplicationCommandOptionChoiceInt{
					Name: "First Slot",
					Value: 0,
				},
				discord.ApplicationCommandOptionChoiceInt{
					Name: "Second Slot",
					Value: 1,
				},
				discord.ApplicationCommandOptionChoiceInt{
					Name: "Third Slot",
					Value: 2,
				},
			},
		},
	},
}

func RestoreHandler(e *handler.CommandEvent) error {
	return response.NotImplemented(e)
}
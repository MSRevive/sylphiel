package dbot

import (
	"context"

	"github.com/disgoorg/log"
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/cache"
	"github.com/disgoorg/disgo/webhook"
)

var (
	Version = "canary"
)

type Bot struct {
	Client bot.Client
	Handler handler.Router
	Webhook webhook.Client
	Logger log.Logger
	Config *Config
	Ctx context.Context
}

func New(ctx context.Context, logger log.Logger, cfg *Config) *Bot {
	return &Bot {
		Ctx: ctx,
		Handler: handler.New(),
		Logger: logger,
		Config: cfg,
	}
}

func (b *Bot) Setup(listeners ...bot.EventListener) (err error) {
	b.Client, err = disgo.New(b.Config.Core.Token,
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentGuilds,
				gateway.IntentGuildMessages,
				gateway.IntentGuildVoiceStates,
				gateway.IntentGuildModeration,
				gateway.IntentGuildMembers,
				gateway.IntentGuildWebhooks,
				gateway.IntentGuildIntegrations,
				gateway.IntentMessageContent,
			),
		),
		bot.WithCacheConfigOpts(
			cache.WithCaches(cache.FlagsAll),
			cache.WithMemberCachePolicy(func(member discord.Member) bool {
				return member.User.ID == b.Client.ID()
			}),
		),
		bot.WithLogger(b.Logger),
		bot.WithEventListeners(listeners...),
	)

	return err
}

// func (b *Bot) SyncCommands() {
// 	// if b.Config.Disc.GuildID.String() == "" {
// 	// 	b.Logger.Info("Syncing commands globally...");
// 	// 	if _, err := b.Client.Rest().SetGlobalCommands(b.Client.ApplicationID(), commands.Commands); err != nil {
// 	// 		b.Logger.Errorf("Failed to sync commands: %s", err)
// 	// 	}
// 	// 	return
// 	// }

// 	b.Logger.Infof("Syncing commands with guild: %s...", b.Config.Disc.GuildID);
// 	if _, err := b.Client.Rest().SetGuildCommands(b.Client.ApplicationID(), b.Config.Disc.GuildID, commands.Commands); err != nil {
// 		b.Logger.Errorf("Failed to sync commands: %s", err)
// 	}

// 	// if err := b.Client.Rest().DeleteGlobalCommand(b.Client.ApplicationID(), 1067673174449340418); err != nil {
// 	// 	b.Logger.Errorf("Failed to sync commands: %s", err)
// 	// }

// 	// if err := b.Client.Rest().DeleteGlobalCommand(b.Client.ApplicationID(), 1079171151831498762); err != nil {
// 	// 	b.Logger.Errorf("Failed to sync commands: %s", err)
// 	// }
// }

func (b *Bot) Start() error {
	if (b.Config.Webhook.Enabled) {
		b.Logger.Info("Events logging enabled.")
		b.Webhook = webhook.New(b.Config.Webhook.ID, b.Config.Webhook.Token)
	}

	return b.Client.OpenGateway(b.Ctx);
}

func (b *Bot) Close() {
	if (b.Config.Webhook.Enabled) {
		b.Webhook.Close(context.TODO())
	}
	b.Client.Close(b.Ctx)
}
package dbot

import (
	"context"

	"github.com/disgoorg/log"
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/disgo/cache"
)

var (
	Version = "canary"
)

type Bot struct {
	Ctx context.Context

	Client bot.Client
	Handler handler.Router

	Logger log.Logger
	Config *Config
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
			),
		),
		bot.WithCacheConfigOpts(
			cache.WithCaches(
				cache.FlagMembers, 
				cache.FlagChannels, 
				cache.FlagGuilds, 
				cache.FlagRoles,
			),
			cache.WithMemberCachePolicy(func(member discord.Member) bool {
				return member.User.ID == b.Client.ID()
			}),
		),
		bot.WithEventListeners(listeners...),
	)

	return err
}

func (b *Bot) Start() error {
	return b.Client.OpenGateway(b.Ctx);
}
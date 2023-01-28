package response

import (
	"fmt"
	"time"

	"github.com/disgoorg/disgo/webhook"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func AuditTest(c webhook.Client, e *events.Ready) error {
	embeds := make([]discord.Embed, 1)
	embeds[0] = discord.NewEmbedBuilder().
	SetColor(12345).
	SetTimestamp(time.Now()).
	SetAuthorName("Test").
	SetTitle("Test Embed").
	SetDescription("Test Description").
	SetFooterText("Footer").
	Build()

	_, err := c.CreateEmbeds(embeds)

	return err
}

func AuditMemberJoined(c webhook.Client, a *events.GuildMemberJoin) error {
	var userAvatar string
	userName := fmt.Sprintf("%s#%s", a.Member.User.Username, a.Member.User.Discriminator)
	userID := fmt.Sprintf("%s", a.Member.User.ID)
	userDesc := fmt.Sprintf("%s created %s", a.Member.Mention(), a.Member.CreatedAt().Format(time.RFC1123))

	if a.Member.Avatar == nil {
		userAvatar = "https://winterfang.com/assets/gfx/bot-avatar.png"
	}else{
		userAvatar = fmt.Sprintf("%s", a.Member.AvatarURL())
	}

	embeds := make([]discord.Embed, 1)
	embeds[0] = discord.NewEmbedBuilder().
	SetColor(0x009999).
	SetTimestamp(time.Now()).
	SetAuthorIcon(userAvatar).
	SetAuthorName(userName).
	SetTitle("Member Joined").
	SetDescription(userDesc).
	SetFooterText(userID).
	Build()

	_, err := c.CreateEmbeds(embeds)

	return err
}

func AuditMemberLeft(c webhook.Client, a *events.GuildMemberLeave) error {
	var userAvatar string
	userName := fmt.Sprintf("%s#%s", a.User.Username, a.User.Discriminator)
	userID := fmt.Sprintf("%s", a.User.ID)
	userDesc := fmt.Sprintf("%s at %s", a.User.ID, time.Now().Format(time.RFC1123))

	if a.Member.Avatar == nil {
		userAvatar = "https://winterfang.com/assets/gfx/bot-avatar.png"
	}else{
		userAvatar = fmt.Sprintf("%s", a.Member.AvatarURL())
	}

	embeds := make([]discord.Embed, 1)
	embeds[0] = discord.NewEmbedBuilder().
	SetColor(0xcc0000).
	SetTimestamp(time.Now()).
	SetAuthorIcon(userAvatar).
	SetAuthorName(userName).
	SetTitle("Member Left").
	SetDescription(userDesc).
	SetFooterText(userID).
	Build()

	_, err := c.CreateEmbeds(embeds)

	return err
}

func AuditVoiceJoined(c webhook.Client, a *events.GuildVoiceJoin) error {
	var userAvatar string
	userName := fmt.Sprintf("%s#%s", a.Member.User.Username, a.Member.User.Discriminator)
	userID := fmt.Sprintf("%s", a.Member.User.ID)
	chanID := &a.VoiceState.ChannelID
	desc := fmt.Sprintf("%s joined %s", userName, discord.ChannelMention(chanID))

	if a.Member.Avatar == nil {
		userAvatar = "https://winterfang.com/assets/gfx/bot-avatar.png"
	}else{
		userAvatar = fmt.Sprintf("%s", a.Member.AvatarURL())
	}

	embeds := make([]discord.Embed, 1)
	embeds[0] = discord.NewEmbedBuilder().
	SetColor(0x009999).
	SetTimestamp(time.Now()).
	SetAuthorIcon(userAvatar).
	SetAuthorName(userName).
	SetTitle("Joined Voice Channel").
	SetDescription(desc).
	SetFooterText(userID).
	Build()

	_, err := c.CreateEmbeds(embeds)

	return err
}

func AuditVoiceLeft(c webhook.Client, a *events.GuildVoiceLeave) error {
	var userAvatar string
	userName := fmt.Sprintf("%s#%s", a.Member.User.Username, a.Member.User.Discriminator)
	userID := fmt.Sprintf("%s", a.Member.User.ID)
	chanID := &a.VoiceState.ChannelID
	desc := fmt.Sprintf("%s left %s", userName, discord.ChannelMention(chanID))
	
	if a.Member.Avatar == nil {
		userAvatar = "https://winterfang.com/assets/gfx/bot-avatar.png"
	}else{
		userAvatar = fmt.Sprintf("%s", a.Member.AvatarURL())
	}

	embeds := make([]discord.Embed, 1)
	embeds[0] = discord.NewEmbedBuilder().
	SetColor(0xcc0000).
	SetTimestamp(time.Now()).
	SetAuthorIcon(userAvatar).
	SetAuthorName(userName).
	SetTitle("Left Voice Channel").
	SetDescription(desc).
	SetFooterText(userID).
	Build()

	_, err := c.CreateEmbeds(embeds)

	return err
}
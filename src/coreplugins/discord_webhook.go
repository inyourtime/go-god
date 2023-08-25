package coreplugins

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var Discord *discordgo.Session

func NewDiscord() {
	var err error
	Discord, err = discordgo.New("Bot")
	if err != nil {
		log.Printf("DiscordGo error: %v", err)
	}
}

func WebhookSend(text string) {
	hookMessage := &discordgo.WebhookParams{
		Content: "```json\n" + text + "\n```",
	}

	if Discord == nil {
		Discord, _ = discordgo.New("Bot")
	}

	_, err := Discord.WebhookExecute(Config.Discord.WebHookID, Config.Discord.Token, false, hookMessage)
	if err != nil {
		log.Printf("Discord webhook error: %v", err)
	}
}

func WebhookSqlLogSend(text string) {
	hookMessage := &discordgo.WebhookParams{
		Content: "```sql\n" + text + "\n```",
	}

	if Discord == nil {
		Discord, _ = discordgo.New("Bot")
	}

	_, err := Discord.WebhookExecute(Config.Discord.SqlWebHookID, Config.Discord.SqlToken, false, hookMessage)
	if err != nil {
		log.Printf("Discord webhook error: %v", err)
	}
}

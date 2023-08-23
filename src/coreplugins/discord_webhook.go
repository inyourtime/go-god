package coreplugins

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func NewDiscord() *discordgo.Session {
	discord, err := discordgo.New("Bot")
	if err != nil {
		log.Printf("DiscordGo error: %v", err)
		return nil
	}
	return discord
}

func WebhookSend(dc *discordgo.Session, text string) {
	config, _ := Env()

	hookMessage := &discordgo.WebhookParams{
		Content: "```json\n" + text + "\n```" + "\n@everyone",
	}
	_, err := dc.WebhookExecute(config.Discord.WebHookID, config.Discord.Token, false, hookMessage)
	if err != nil {
		log.Printf("Discord webhook error: %v", err)
	}
}

func WebhookSqlLogSend(dc *discordgo.Session, text string) {
	config, _ := Env()
	
	hookMessage := &discordgo.WebhookParams{
		Content: "```sql\n" + text + "\n```",
	}
	_, err := dc.WebhookExecute(config.Discord.SqlWebHookID, config.Discord.SqlToken, false, hookMessage)
	if err != nil {
		log.Printf("Discord webhook error: %v", err)
	}
}

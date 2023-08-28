package coreplugins

import (
	"encoding/json"
	"gopher/src/model"
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

func WebhookSend(text string, ev model.ServerEnvironment) {
	data := model.DiscordErrorLog{}
	// Unmarshal the input JSON string into the map
	if err := json.Unmarshal([]byte(text), &data); err != nil {
		panic(err)
	}
	// Marshal the map back into a pretty-printed JSON string
	// prettyJSON, err := json.MarshalIndent(data, "", "  ")
	// if err != nil {
	// 	panic(err)
	// }

	emb := []*discordgo.MessageEmbed{
		{
			Type:        "rich",
			Title:       "Environment",
			Description: "**End Point**\n" + ev.Hostname,
			Color:       15548997,
			Timestamp:   data.Timestamp,
		},
	}

	// fmt.Println(Ctx)
	s := "\n**[" + ev.Method + "]** `" + ev.Url + "`"
	// s := ""

	hookMessage := &discordgo.WebhookParams{
		Embeds:  emb,
		Content: "**Server Error :boom:** - TypeError: " + data.Message + " at " + data.Caller + s,
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

package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/helltf/go-discord-bot/env"
)

var (
	BotToken = "DISCORD_BOT_TOKEN"
)

func CreateNewBotConnection() {
	discordSecret := GetDiscordSecret()
	discordBot := CreateNewBot(discordSecret)
	discordBot.AddHandler(messageHandler)
	OpenConnection(discordBot)
}


func OpenConnection(discordBot *discordgo.Session){
	errOpen := discordBot.Open()

	if errOpen != nil {
		panic(errOpen)
	}
}

func GetDiscordSecret() string {
	return env.GetEnvVariable(BotToken)
}

func CreateNewBot(discordSecret string) *discordgo.Session {

	discordBot, err := discordgo.New("Bot " + discordSecret)

	if err != nil {
		panic(err)
	}

	return discordBot
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if HasPrefix(m.Content) {

	}
}

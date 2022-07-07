package main

import (
	"Pavel-dis-tele-bot/util"
	"fmt"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	telegramToken := viper.GetString("telegram_token")
	if telegramToken == "" {
		panic(fmt.Errorf("cannot find telegramToken"))
	}
	discordToken := viper.GetString("discord_token")
	if discordToken == "" {
		panic(fmt.Errorf("cannot find discordoken"))
	}

	messages := make(chan string)

	util.TeleConnect(telegramToken, messages)
	/*	if err != nil {
		panic(fmt.Errorf("fatal error accessing telegram bot: %w", err))
	}*/
	util.DiscInit(discordToken, messages)

}

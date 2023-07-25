package notification

import (
	"cloudflare-pages-hook/pkg/parseConfig"
	"log"
)

func initNotifier() {
	if parseConfig.NotifierType == "telegram" || parseConfig.NotifierType == "tg" {
		log.Println("init telegram")
		N = InitTelegram(parseConfig.NotifierToken, parseConfig.NotifierChatId)
	}
}

func Init() {
	initNotifier()
}

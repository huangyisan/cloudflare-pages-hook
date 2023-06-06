package notification

import (
	flag "github.com/spf13/pflag"
	"log"
)

var notifierType string
var notifierToken string
var notifierChatId string

func init() {
	flag.StringVarP(&notifierType, "notifier", "n", "", "notifier type")
	flag.StringVarP(&notifierToken, "token", "t", "", "notifier token")
	flag.StringVarP(&notifierChatId, "chatId", "d", "", "notifier chatId")

	flag.Parse()
	initNotifier()

}
func initNotifier() {
	if notifierType == "telegram" {
		log.Println("init telegram")
		N = InitTelegram(notifierToken, notifierChatId)
	}
}

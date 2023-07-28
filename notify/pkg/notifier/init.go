package notifier

import "log"

func InitNotifier(notifierType, notifierToken string) INotifier {
	if notifierType == "telegram" || notifierType == "tg" {
		log.Println("init telegram")
		N = InitTelegram(notifierToken)
	}
	return N
}

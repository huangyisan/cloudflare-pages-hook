package notification

import (
	"cloudflare-pages-hook/common"
	"fmt"
	"log"
)

type Telegram struct {
	Token  string
	ChatId string
}

func (t *Telegram) send(msg string) bool {
	url := t.url()
	body := make(map[string]string)
	body = map[string]string{"text": msg}
	res, err := common.Post(url, body, "", nil)
	if err != nil {
		log.Printf("%+v", err)
		return false
	} else {
		log.Printf("%s", res)
		return true
	}
}

func (t *Telegram) url() string {
	parseMode := "markdown"
	url := fmt.Sprintf("https://api.Telegram.org/%s/sendMessage?chat_id=%s&parse_mode=%s", t.Token, t.ChatId, parseMode)
	return url
}

func InitTelegram(Token, chatId string) *Telegram {
	return &Telegram{
		Token:  Token,
		ChatId: chatId,
	}
}

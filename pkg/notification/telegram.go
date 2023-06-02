package notification

import (
	"cloudflare-pages-hook/common"
	"fmt"
	"log"
)

type telegram struct {
	token  string
	chatId string
}

func (t *telegram) send(msg string) {
	url := t.url()
	body := make(map[string]string)
	body = map[string]string{"text": msg}
	res, err := common.Post(url, body, "", nil)
	if err != nil {
		log.Printf("%+v", err)
	} else {

	}

}

func (t *telegram) url() string {
	parseMode := "markdown"
	url := fmt.Sprintf("https://api.telegram.org/%s/sendMessage?chat_id=%s&parse_mode=%s", t.token, t.chatId, parseMode)
	return url
}

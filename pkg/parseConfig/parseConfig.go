package parseConfig

import (
	flag "github.com/spf13/pflag"
	"log"
	"time"
)

var NotifierType string
var NotifierToken string
var NotifierChatId string
var CFScanWaitTime time.Duration

func init() {
	flag.StringVarP(&NotifierType, "notifier", "n", "", "notifier type")
	flag.StringVarP(&NotifierToken, "token", "t", "", "notifier token")
	flag.StringVarP(&NotifierChatId, "chatId", "d", "", "notifier chatId")

	flag.DurationVarP(&CFScanWaitTime, "scan", "s", 2*time.Minute, "scan wait time")
	flag.Parse()
	log.Printf("\nArgs:\nNotifierType: %s\nNotifierToken: %s\nNotifierChatId: %s\nCFScanWaitTime: %v\n", NotifierType,
		NotifierToken, NotifierChatId, CFScanWaitTime)
}

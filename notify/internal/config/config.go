package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Notifier struct {
		Type string
		// read from env
		Token string `json:",env=NOTIFIER_TOKEN"`
	}
}

package main

import (
	"cloudflare-pages-hook/notify/internal/config"
	"cloudflare-pages-hook/notify/internal/handler"
	"cloudflare-pages-hook/notify/internal/svc"
	"cloudflare-pages-hook/notify/pkg/notifier"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/notify.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	// init Notifier
	notifier.InitNotifier(c.Notifier.Type, c.Notifier.Token)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

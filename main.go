package main

import (
	"cloudflare-pages-hook/pkg/notification"
	"cloudflare-pages-hook/routers"
	"fmt"
)

func main() {
	// init notification
	notification.Init()

	r := routers.SetupRouter()
	if err := r.Run(); err != nil {
		fmt.Println("service failed, err:%v\n", err)
	}
}

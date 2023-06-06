package main

import (
	"cloudflare-pages-hook/routers"
	"fmt"
)

func main() {

	r := routers.SetupRouter()
	if err := r.Run(); err != nil {
		fmt.Println("service failed, err:%v\n", err)
	}
}

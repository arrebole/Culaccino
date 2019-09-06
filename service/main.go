package main

import (
	"github.com/arrebole/culaccino/service/config"
	"github.com/arrebole/culaccino/service/route"
)

func main() {
	server := route.New()
	server.Run(config.ListenPort)
}

package main

import (
	"flag"
	"fmt"

	"github.com/raismaulana/digilibP/application"
	"github.com/raismaulana/digilibP/application/registry"
)

func main() {
	appMap := map[string]func() application.RegistryContract{
		"app": registry.NewApp(),
	}

	flag.Parse()

	app, exist := appMap[flag.Arg(0)]
	if exist {
		application.Run(app())
	} else {
		fmt.Println("You may try this app name:")
		for appName := range appMap {
			fmt.Printf("%s\n", appName)
		}
	}
}

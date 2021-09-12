package main

import (
	"flag"
	"fmt"

	"github.com/raismaulana/my-gogen-rest-setup/application"
	"github.com/raismaulana/my-gogen-rest-setup/application/registry"
)

func main() {
	appMap := map[string]func() application.RegistryContract{
		"usingdb": registry.NewApp(),
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

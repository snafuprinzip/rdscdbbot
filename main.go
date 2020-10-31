package main

import (
	"fmt"

	"github.com/snafuprinzip/renscbot/bot"
	"github.com/snafuprinzip/renscbot/config"
	"github.com/snafuprinzip/renscbot/starcitizen"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	starcitizen.CreateNewShips()
	for _, ship := range starcitizen.Ships {
		fmt.Println(ship)
		fmt.Println(string(ship.Yaml()))
	}

	starcitizen.CreateNewComponents()

	for _, cmp := range starcitizen.Components {
		fmt.Println(cmp)
		fmt.Println(string(cmp.Yaml()))
	}

	starcitizen.CreateNewHardPoints()

	for _, hp := range starcitizen.HardPoints {
		fmt.Println(hp)
		fmt.Println(string(hp.Yaml()))
	}

	bot.Start()

	<-make(chan struct{})
	return
}

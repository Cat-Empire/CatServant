package main

import(
	"fmt"

	"github.com/Cat-Empire/CatServant/bot"
	"github.com/Cat-Empire/CatServant/config"
)

func main(){
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bot.Start()
	<-make(chan struct{})
	return

}
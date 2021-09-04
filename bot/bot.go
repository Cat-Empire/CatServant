package bot

import (
	"fmt"
	"github.com/Cat-Empire/CatServant/config"
	"github.com/bwmarrin/discordgo"

)
var BotID string
var CatServant *discordgo.Session 

func Start(){
	
	//make a Bot
	CatServant, err := discordgo.New("Bot " + config.Token)
	
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	
	u, err := CatServant.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID
	
	//Register messageHandler as a callback for the messageCreate events.	
	CatServant.AddHandler(messageHandler)

	

	err = CatServant.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Bot is running! id:%q\n",BotID)
	


}
func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == BotID{
		return
	}

	 if message.Content == "Koda" {

		_, _ = session.ChannelMessageSend(message.ChannelID, "Hi Koda May!")
		
	 }

}
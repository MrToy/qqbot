package main

import (
	//"fmt"
	"github.com/MrToy/qqbot"
)

func main() {
	store := qqbot.NewUserStore("./user")
	defer store.DB.Close()
	user := store.Get()
	if err := user.Login(); err != nil {
		qqbot.WebHandler(user.WaitVerify())
		user.Login()
	}
	store.Put(user)
	user.Login()

	user.GetOnline()

	for msg := range user.Poll() {
		go func() {
			if msg.Type == "message" {
				user.SendMessage(msg.From, qqbot.Tulin(msg.From, msg.Content, "c9577a0e99dde2b9be80b6d56e54a1bc"))
			}
			if msg.Type == "group_message" {
				//go user.SendGroupMessage(msg.From, qqbot.Tulin(msg.From, msg.Content, "c9577a0e99dde2b9be80b6d56e54a1bc"))
			}
		}()
	}
}

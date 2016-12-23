package main

import (
	"os"
	"path/filepath"

	"github.com/MrToy/qqbot"
)

func main() {
	store := qqbot.NewUserStore(filepath.Join(os.TempDir(), "qqbot"))
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
			//个人消息
			if msg.Type == "message" {
				user.SendMessage(msg.From, qqbot.Tulin(msg.From, msg.Content, "c9577a0e99dde2b9be80b6d56e54a1bc"))
			}
			//群消息
			if msg.Type == "group_message" {
				user.SendGroupMessage(msg.From, qqbot.Tulin(msg.From, msg.Content, "c9577a0e99dde2b9be80b6d56e54a1bc"))
			}
			//讨论组消息
			if msg.Type == "discu_message" {
				user.SendDiscuMessage(msg.From, qqbot.Tulin(msg.From, msg.Content, "c9577a0e99dde2b9be80b6d56e54a1bc"))
			}
		}()
	}
}

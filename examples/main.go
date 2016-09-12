package main

import (
	"fmt"
	"github.com/MrToy/qqbot"
)

func main() {
	fmt.Println("start")
	store := qqbot.NewUserStore()
	defer store.DB.Close()
	user := store.Get()
	if err := user.Login(); err != nil {
		qqbot.WebHandler(user.WaitVerify())
		user.Login()
	}
	store.Put(user)
	user.Login()
	fmt.Println(user.SendMessage(452719485, "2333"))
	for msg := range user.Poll() {
		fmt.Println(msg)
	}
}

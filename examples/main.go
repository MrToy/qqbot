package main

import (
	"fmt"
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
	fmt.Println(user.GetGroups())
}

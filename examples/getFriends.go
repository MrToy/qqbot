package main

import (
	"fmt"
	"github.com/MrToy/qqbot"
)

func main() {
	user := qqbot.NewUser()
	qqbot.WebHandler(user.WaitVerify())
	user.Login()
	fmt.Println(user.GetFriends())
}

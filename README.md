This is a client of the smart-web-qq in the Go programming language.

Installation
-----------

    go get github.com/MrToy/qqbot/examples
    cd $GOPATH/github.com/MrToy/qqbot/examples
    go build robot.go

or directly

    go run robot.go

Usage
-----------

* This is how to get friends list
```go
import "github.com/MrToy/qqbot"
...
user := qqbot.NewUser()
qqbot.WebHandler(user.WaitVerify())
user.Login()
fmt.Println(user.GetFriends())
...
```
start qqbot and open http://loacalhost:8081/captcha ,then use mobile qq scan this captcha


* This is how to login with cache
```go
import "github.com/MrToy/qqbot"
...
store := qqbot.NewUserStore("./user")
defer store.DB.Close()
user := store.Get()
if err := user.Login(); err != nil {
	qqbot.WebHandler(user.WaitVerify())
	user.Login()
}
store.Put(user)
fmt.Println(user.GetFriends())
...
```

* Message Get and Send 【you must be logined】
```go
...
user.SendMessage(uin,"some text")
...
for msg := range user.Poll() {
   fmt.Println(msg)
}
...
```

And You can use ["图灵机器人"](http://www.tuling123.com/) API for reduce the message
```go
qqbot.Tulin(uin int64,msg string, key string) string
//This uin is a tag who speaked
//This key you can get by 图灵机器人's website
``` 

More usage examples please view
https://github.com/MrToy/qqbot/examples

Documentation
-----------
You can read APIs in package documentation here
https://godoc.org/github.com/MrToy/qqbot
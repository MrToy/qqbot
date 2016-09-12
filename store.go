package qqbot

import (
	"encoding/json"
	"github.com/syndtr/goleveldb/leveldb"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type UserStore struct {
	DB *leveldb.DB
}

func NewUserStore() *UserStore {
	db, _ := leveldb.OpenFile(filepath.Join(os.TempDir(), "qqbot"), nil)
	return &UserStore{DB: db}
}

func (this *UserStore) Put(user *User) {
	u, _ := url.Parse("http://s.web2.qq.com/api/getvfwebqq")
	u2, _ := url.Parse("http://d1.web2.qq.com/channel/login2")
	cstr, _ := json.Marshal(user.Client.Jar.Cookies(u))
	cstr2, _ := json.Marshal(user.Client.Jar.Cookies(u2))
	ustr, _ := json.Marshal(user)
	this.DB.Put([]byte("cookies"), cstr, nil)
	this.DB.Put([]byte("cookies2"), cstr2, nil)
	this.DB.Put([]byte("user"), ustr, nil)
}

func (this *UserStore) Get() *User {
	u, _ := url.Parse("http://s.web2.qq.com/api/getvfwebqq")
	u2, _ := url.Parse("http://d1.web2.qq.com/channel/login2")
	cstr, _ := this.DB.Get([]byte("cookies"), nil)
	cstr2, _ := this.DB.Get([]byte("cookies2"), nil)
	ustr, _ := this.DB.Get([]byte("user"), nil)
	user := NewUser()
	var cookies, cookies2 []*http.Cookie
	json.Unmarshal(ustr, &user)
	json.Unmarshal(cstr, &cookies)
	json.Unmarshal(cstr2, &cookies2)
	user.Client.Jar.SetCookies(u, cookies)
	user.Client.Jar.SetCookies(u2, cookies2)
	return user
}

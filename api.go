package qqbot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
)

type Friend struct {
	Nick     string
	MarkName string
	Uin      int64
}

func (this *User) GetFriends() []Friend {
	req, _ := http.NewRequest("POST", "http://s.web2.qq.com/api/get_user_friends2", bytes.NewReader([]byte("r=%7B%22vfwebqq%22%3A%22"+this.Vfwebqq+"%22%2C%22hash%22%3A%22"+TxHash(this.Uin, this.Ptwebqq)+"%22%7D")))
	req.Header.Add("Referer", "http://s.web2.qq.com/proxy.html?v=20130916001&callback=1&id=1")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, _ := this.Client.Do(req)
	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	var result struct {
		RetCode int
		Result  struct {
			Info      []Friend
			Marknames []Friend
		}
	}
	if result.RetCode != 0 {
		fmt.Println(string(data))
	}
	json.Unmarshal(data, &result)
	friendsMap := make(map[int64]Friend)
	for _, it := range result.Result.Info {
		friendsMap[it.Uin] = it
	}
	for _, it := range result.Result.Marknames {
		tmp := friendsMap[it.Uin]
		tmp.MarkName = it.MarkName
		friendsMap[it.Uin] = tmp
	}
	friends := []Friend{}
	for _, it := range friendsMap {
		friends = append(friends, it)
	}
	return friends
}

type Group struct {
	Code int64
	Name string
	Gid  int64
}

func (this *User) GetGroups() []Group {
	req, _ := http.NewRequest("POST", "http://s.web2.qq.com/api/get_group_name_list_mask2", bytes.NewReader([]byte("r=%7B%22vfwebqq%22%3A%22"+this.Vfwebqq+"%22%2C%22hash%22%3A%22"+TxHash(this.Uin, this.Ptwebqq)+"%22%7D")))
	req.Header.Add("Referer", "http://s.web2.qq.com/proxy.html?v=20130916001&callback=1&id=1")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, _ := this.Client.Do(req)
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	var result struct {
		RetCode int
		Result  struct {
			Gnamelist []Group
		}
	}
	if result.RetCode != 0 {
		fmt.Println(string(data))
	}
	json.Unmarshal(data, &result)
	return result.Result.Gnamelist
}

type SelfInfo struct {
	Account  int64
	Nick     string
	Gender   string
	Country  string
	Province string
	City     string
}

func (this *User) GetSelfInfo() SelfInfo {
	req, _ := http.NewRequest("GET", "http://s.web2.qq.com/api/get_self_info2", nil)
	req.Header.Add("Referer", "http://s.web2.qq.com/proxy.html?v=20130916001&callback=1&id=1")
	res, _ := this.Client.Do(req)
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	var result struct {
		RetCode int
		Result  SelfInfo
	}
	if result.RetCode != 0 {
		fmt.Println(string(data))
	}
	json.Unmarshal(data, &result)
	return result.Result
}

func (this *User) SendMessage(uin int64, content string) error {
	uinStr := strconv.FormatInt(uin, 10)
	msgId := strconv.Itoa(rand.Intn(8))
	req, _ := http.NewRequest("POST", "https://d1.web2.qq.com/channel/send_buddy_msg2", bytes.NewReader([]byte("r=%7B%22to%22%3A"+uinStr+"%2C%22content%22%3A%22%5B%5C%22"+content+"%5C%22%2C%5B%5C%22font%5C%22%2C%7B%5C%22name%5C%22%3A%5C%22%E5%AE%8B%E4%BD%93%5C%22%2C%5C%22size%5C%22%3A10%2C%5C%22style%5C%22%3A%5B0%2C0%2C0%5D%2C%5C%22color%5C%22%3A%5C%22000000%5C%22%7D%5D%5D%22%2C%22face%22%3A525%2C%22clientid%22%3A53999199%2C%22msg_id%22%3A"+msgId+"%2C%22psessionid%22%3A%22"+this.Pssesionid+"%22%7D")))
	req.Header.Add("Referer", "https://d1.web2.qq.com/cfproxy.html?v=20151105001&callback=1")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, _ := this.Client.Do(req)
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	var result struct {
		Retcode int
		ErrCode int
		Msg     string
	}
	json.Unmarshal(data, &result)
	if result.ErrCode != 0 || result.Retcode != 0 {
		return errors.New(result.Msg)
	}
	return nil
}

type Message struct {
	Type    string //one of message|group_message
	From    int64
	To      int64
	Content string
}

func (this *User) Poll() chan Message {
	c := make(chan Message)
	go func() {
		req, _ := http.NewRequest("POST", "https://d1.web2.qq.com/channel/poll2", bytes.NewReader([]byte("r=%7B%22ptwebqq%22%3A%22"+this.Ptwebqq+"%22%2C%22clientid%22%3A53999199%2C%22psessionid%22%3A%22"+this.Pssesionid+"%22%2C%22key%22%3A%22%22%7D")))
		req.Header.Add("Referer", "https://d1.web2.qq.com/cfproxy.html?v=20151105001&callback=1")
		res, _ := this.Client.Do(req)
		defer res.Body.Close()
		data, _ := ioutil.ReadAll(res.Body)
		var result struct {
			Retcode int
			Result  []struct {
				Poll_type string
				Value     struct {
					From_uin int64
					To_uin   int64
					Content  [2]string
				}
			}
		}
		fmt.Println(string(data))
		json.Unmarshal(data, &result)
		for _, it := range result.Result {
			msg := Message{
				Type:    it.Poll_type,
				From:    it.Value.From_uin,
				To:      it.Value.To_uin,
				Content: it.Value.Content[1],
			}
			c <- msg
		}
	}()
	return c
}

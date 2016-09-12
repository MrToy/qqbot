package qqbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func TxHash(uin int, pt string) string {
	n := make([]byte, 4)
	for i := 0; i < len(pt); i++ {
		n[i%4] ^= byte(pt[i])
	}
	u := []byte{'E', 'C', 'O', 'K'}
	v := make([]byte, 4)
	for i := 0; i < 4; i++ {
		v[i] = byte(uin>>((3-uint(i))*8)&255 ^ int(u[i]))
	}
	r := make([]byte, 8)
	for i := 0; i < 8; i++ {
		if i%2 == 0 {
			r[i] = n[i>>1]
		} else {
			r[i] = v[i>>1]
		}
	}
	var res string
	for _, v := range r {
		hex := strconv.FormatInt(int64(v), 16)
		if len(hex) == 1 {
			res += "0" + hex
		} else {
			res += hex
		}
	}
	return strings.ToUpper(res)
}

func WebHandler(c chan image.Image) {
	var img image.Image
	r := http.NewServeMux()
	r.HandleFunc("/captcha", func(w http.ResponseWriter, r *http.Request) {
		if img != nil {
			png.Encode(w, img)
		}
	})
	l, _ := net.Listen("tcp", ":8081")
	go func() {
		for img = range c {
		}
		l.Close()
	}()
	fmt.Println("Need verify,please open http://localhost:8081/captcha")
	server := http.Server{Handler: r}
	server.Serve(l)
}

func Tulin(uid int64, info, key string) string {
	fmt.Println("from:", uid, info)
	uidStr := strconv.FormatInt(uid, 10)
	str := `{"key":"` + key + `","info":"` + info + `","userid":"` + uidStr + `"}`
	res, _ := http.Post("http://www.tuling123.com/openapi/api", "application/json;charset=utf-8", bytes.NewReader([]byte(str)))
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	var result struct {
		Code int64
		Text string
		Url  string
	}
	json.Unmarshal(data, &result)
	fmt.Println("to:", uid, result.Text)
	return result.Text
}

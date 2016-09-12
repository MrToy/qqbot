package qqbot

import (
	"fmt"
	"image"
	"image/png"
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

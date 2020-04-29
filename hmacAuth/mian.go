package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func HmacSha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	//return hex.EncodeToString(h.Sum(nil))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func Sha1(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func main() {
	s := "date: Tue, 28 Apr 2020 10:01:50 GMT\nGET /exp1 HTTP/1.1"
	fmt.Println(HmacSha256(s, "secret456"))
}

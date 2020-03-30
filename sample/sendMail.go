package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"strings"
)

func getPublicIp() (string, error) {
	url := "http://ifconfig.co"
	resp, err := http.Get(url)

	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	//fmt.Print(string(ip))
	return string(ip), err
}

func main() {
	host := "smtp.163.com"
	username := "xxx@163.com"
	password := "xxx"
	sender := "xxx@163.com"
	myIp, err := getPublicIp()
	if err != nil {
		log.Fatal(err)
	}

	// Set up authentication information.
	auth := smtp.PlainAuth("", username, password, host)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"x1@qq.com", "x2@163.com"}
	msg := []byte("To: " + strings.Join(to, ",") + "\r\n" +
		"Subject: My Public IP\r\n" +
		"\r\n" +
		myIp + "\r\n")
	err = smtp.SendMail(host+":25", auth, sender, to, msg)
	if err != nil {
		log.Fatal()
	}
	log.Println("Send success!")
}

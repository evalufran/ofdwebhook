package main

import (
	"log"
	"net/http"
	"mynetsmtp"
	"path"
)

func main() {

	http.HandleFunc("/", Receive)
	err := http.ListenAndServe(":8666", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func Receive(w http.ResponseWriter, req *http.Request) {
	send("Warning, you received this email because ELK stack triggered an alarm for: " + path.Base(req.URL.Path))
}

func send(body string) {
	from := "provawatcher@gmail.com"
	pass := "T4FquMPbTSzWAHT"
	to := "eva.francucci@sctech.info"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: ELK Alarm Received\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent!")
}

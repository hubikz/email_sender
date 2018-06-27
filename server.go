package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Email struct {
	Recipient   string `json:"recipient"`
	Sender      string `json:"sender"`
	Subject     string `json:"subject"`
	ContentType string `json:"content_type"`
	Body        string `json:"body"`
}

func sendEmail(w http.ResponseWriter, r *http.Request) {

	// Read body
	body, err := ioutil.ReadAll(r.Body)
	log.Printf("%v", body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	message := Email{}
	err = json.Unmarshal(body, &message)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// log.Println("Message with %s recived", message.Subject)
	w.Write([]byte(message.Subject))
}

// ListenAndServe start listening for client connections on given port
func main() {
	http.HandleFunc("/email", sendEmail)
	http.ListenAndServe(":8080", nil)
}

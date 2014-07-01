package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/mail"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/incoming_mail", incomingMail)

	http.Handle("/", r)
	fmt.Println("Server started. Listening...")
	http.ListenAndServe(":3000", nil)
}

func incomingMail(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	msg, err := mail.ReadMessage(r.Body)

	if err != nil {
		fmt.Printf("Error reading body: %v", err)
		return
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(msg.Body)

	fmt.Printf("Received mail: %v", buf.String())
}

func printAddrs(msg *mail.Message, filedName string) {
	addrs, err := msg.Header.AddressList(filedName)
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		fmt.Println("Addr:", addr)
	}
}

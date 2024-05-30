package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// httpserver 8080
// greeting message

const GreetingMessage = "Hello Dammmmnn!!!!!!!!!"

func main() {
	var m string

	mEnv := os.Getenv("GREETING_MESSAGE")

	if mEnv != "" {
		m = mEnv
	} else {
		m = GreetingMessage
	}

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	http.HandleFunc("/greeting", greeting(m))

	go func() {
		log.Fatalln(http.ListenAndServe(":8080", nil))
	}()

	log.Println("Server is ready to handle requests at :8080")

	<-ch

	log.Println("Shutting down...")
}

func greeting(m string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("%s", m)))
	}
}

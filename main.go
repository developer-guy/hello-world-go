package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const DefaultGreetingMessage = "Hello World!"

func main() {
	// start new http client server
	greeting := os.Getenv("GREETING")
	if greeting == "" {
		greeting = DefaultGreetingMessage
	}
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGKILL, syscall.SIGTERM)

	http.HandleFunc("/greetings", greetings(greeting))

	go func() {
		log.Fatalln(http.ListenAndServe(":8080", nil))
	}()

	fmt.Println("Application is started on port 8080")

	<-ch

	fmt.Println("Application is shutting down")
}

func greetings(gm string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Greetings %s!", gm)))
	}
}

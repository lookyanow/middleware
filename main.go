package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func live(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Alive!"))
}

func ready(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Ready!"))
}

func h1(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now())
	fmt.Println("h1")
	fmt.Println(time.Now())
}

func h2(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now())
	fmt.Println("h2")
	fmt.Println(time.Now())
}

func h3(w http.ResponseWriter, r *http.Request) {
	fmt.Println("h3")
}

func logger(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now())
		defer fmt.Println(time.Now())

		f(w, r)
	}
}
func logger2(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.NotFound(w, r)
			return
		}

		f(w, r)
	}
}

type config struct {
	env    string
	dbConn string
}

// handlers hang off a type that understands the environment/server details
func (c config) h8() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("h8 in " + c.env + " via " + c.dbConn)
	}

}

func main() {
	cfg := config{"production", "mysql"}
	// Handlers
	http.HandleFunc("/live", live)
	http.HandleFunc("/ready", ready)
	http.HandleFunc("/h1", h1)
	http.HandleFunc("/h2", h2)
	http.HandleFunc("/h3", logger(h3))
	http.HandleFunc("/h4", logger2(h3))
	http.HandleFunc("/h8", cfg.h8())

	listenAddress := ":8080"

	// log.Printf("listening on: %v", listenAddress)
	fmt.Printf("Listening on: %v", listenAddress)
	log.Fatal("Listener stoped with fatal error", http.ListenAndServe(listenAddress, nil))
}

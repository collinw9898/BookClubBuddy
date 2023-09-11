package main

import (
	"BookClubBuddy/web"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", web.ServeTemplate)

	log.Print("Starting server on port 3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

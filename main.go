package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/yottamusic/DeviceConfig/devapi"
)

func speakerListHandler(w http.ResponseWriter, r *http.Request) {
	devapi.SpeakerListHandler(w, r)
}

func setupRoutes() {
	// Handle Configuration Requests for Wi-FI and Speakers
	//http.HandleFunc("/wifi/api/v1/check", speakerListHandler)
	//http.HandleFunc("/wifi/api/v1/list", speakerListHandler)
	http.HandleFunc("/speakers/api/v1/list", speakerListHandler)
	//http.HandleFunc("/speakers/api/v1/config", speakerListHandler)

	// Handle Any Request
	port := flag.String("p", "80", "Port to Serve UI on")
	directory := flag.String("d", "./httpdocs", "Directory of Webpages to Serve")
	flag.Parse()

	fileServer := http.FileServer(http.Dir(*directory))
	http.Handle("/", fileServer)

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		panic(err)
	}
	log.Fatal(err)
}

func main() {
	setupRoutes()
}

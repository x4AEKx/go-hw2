package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	configPath := flag.String("config", "configs/config.json", "path to config file")
	flag.Parse()

	r := http.NewServeMux()
	r.HandleFunc("/status", statusHandler)

	err := http.ListenAndServe(":8081", r)
	if err != nil {
		log.Fatal(err)
	}
}

func statusHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)

	var data = []byte("123")

	w.Write(data)
}

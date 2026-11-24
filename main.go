package main

import (
	"log"
	"net/http"
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler(cfg))
	mux.HandleFunc("/healthz", healthzHandler)

	addr:= ":8080"
	log.Printf("starting hello-svc env=%s db_host=%s on %d", cfg.Env, cfg.DBHost, addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

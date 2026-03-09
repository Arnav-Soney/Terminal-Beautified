package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Config struct {
	PluginText string `json:"pluginText"`
	P10kText   string `json:"p10kText"`
	FontText   string `json:"fontText"`
}

func getConfig() Config {
	return Config{
		PluginText: PLUGIN_TEXT,
		P10kText:   P10K_TEXT,
		FontText:   FONT_TEXT,
	}
}

func configHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	config := getConfig()
	json.NewEncoder(w).Encode(config)
}

func main() {
	// Serve static files from current directory
	staticDir := "."
	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/", fs)
	http.HandleFunc("/api/config", configHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	log.Printf("Serving files from: %s", staticDir)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

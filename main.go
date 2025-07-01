package main

import (
	"log"
	"net/http"
	"encoding/json"
)
func main() {
	http.HandleFunc("/demo", demoHandler)
	log.Println("Server is starting...")
	er := http.ListenAndServe(":8080", nil)
	if er != nil {
		log.Fatal("Error starting server:", er)
	}
}

func demoHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%+v", r)

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := map[string]string{
		"message": "Hello, World!",
	}

	w.Header().Set("Content-Type", "application/json")

	// data, err := json.Marshal(response)
	// if err != nil {
	// 	http.Error(w, "Internal server error", http.StatusInternalServerError)
	// 	return
	// }

	// w.Write(data)
	json.NewEncoder(w).Encode(response)
}
package routeur

import (
	"fmt"
	"log"
	"net/http"
)

func RUN() {
	// Set up your other handlers
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/CreatChar", CreatHandler)
	http.HandleFunc("/CharList", ListHandler)

	// Serve static files from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Print statement indicating server is running
	fmt.Println("Server is running on :8080 http://localhost:8080")

	// Start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", nil)
}

func CreatHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "newChar", nil)
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "listChar", nil)
}

package main

import (
	"log"
	"net/http"
	"os"
	"fmt"
)

func outputHTML(w http.ResponseWriter, req *http.Request, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(w, req, file.Name(), fi.ModTime(), file)
}

// Homepage
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Homepage Hit!")
	outputHTML(w, r, "home.html")
}

func main() {
	log.Println("Trying to start server...")
	http.HandleFunc("/", HomePage)
	server_port := ":8080"
	log.Printf("Server is running at port %s\n", server_port)
	log.Fatal(http.ListenAndServe(server_port, nil))
}
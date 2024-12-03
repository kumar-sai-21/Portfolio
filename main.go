package main

import (
	"log"
	"net/http"
)



func aboutPage(w http.ResponseWriter, r *http.Request) {
	// Render the about html page
	http.ServeFile(w, r, "static/about.html")
}

func stylePage(w http.ResponseWriter, r *http.Request) {
	// Render the style css page
	http.ServeFile(w, r, "static/style.css")
}

func scriptPage(w http.ResponseWriter, r *http.Request) {
	// Render the script js page
	http.ServeFile(w, r, "static/script.js")
}


func main() {

	
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/style.css", stylePage)
	http.HandleFunc("/script.js", scriptPage)	



	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

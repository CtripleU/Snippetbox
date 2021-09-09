package main

import (
	"log"
	"net/http"
)

//Defining a home handler function to initialize a new servermux containing "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}


//Using the http.NewServeMux() function to initialize a new servemux, and registering the home function as the handler for the "/" URL pattern. 
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	
	//Register two more handler fumctions
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	
	log.Println("Starting server on :4000")
	err:= http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}


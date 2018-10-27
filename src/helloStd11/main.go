package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
)

var (
	Phrase = "Hello, Gopher! Also on gcloud ?"
)

/* Hander name is exported for test purposes. Import like "../" is not supported by dev_appserver.py*/
func Hello(w http.ResponseWriter, r *http.Request) {
	// log.Println("request on", r.RequestURI)
	// Not printing version during tests. Type varies with test method.
	/*
		if wType := reflect.TypeOf(w).String(); wType != "*httptest.ResponseRecorder" &&
			wType != "*http.response" {
			log.Println("running", runtime.Version())
		}
	*/
	log.Println("running", runtime.Version())
	fmt.Fprint(w, Phrase)
}

func main() {
	/* */
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	http.HandleFunc("/", Hello)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	/* */
	// appengine.Main()

}

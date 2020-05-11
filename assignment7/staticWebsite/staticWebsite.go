package staticWebsite

import (
	"log"
	"net/http"
)


func StaticFunc() {
	//Creates the handler to respond to HTTP request
	fs := http.FileServer(http.Dir("./static"))
	//Registers the file serveras the handler for all requests,
	http.Handle("/", fs)

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}


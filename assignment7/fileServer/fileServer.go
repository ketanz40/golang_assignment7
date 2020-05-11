package fileServer

import (
	"log"
	"net/http"
	"fmt"
)

func FileServerFunc() {
	dir := http.Dir(//My system's directory, which differs among other systems
		"C:\\Users\\joshi\\Documents")
	http.Handle("/", http.FileServer(dir))

	fmt.Println("Serving on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
package main

import (
    "fmt"
	"os"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	hostname := os.Getenv("HOSTNAME")
	fmt.Fprint(w, "Hello form container ", hostname, " !")
}

func main() {
	fmt.Println("Listening on 0.0.0.0:8080")
    http.HandleFunc("/", helloHandler)
    http.ListenAndServe(":8080", nil)
}

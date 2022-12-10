package conn;

import (
	"fmt"
	"net/http"
)


func getRoot(w http.ResponseWriter, r *http.Request) { 
	fmt.Printf("got / request\n")
}


func StartHttpServer() {
	
	fmt.Println("Starting HTTP server")

	http.HandleFunc("/", getRoot)

	go func() {
		http.ListenAndServe(":8000", nil)
	}()
}

package main

import (
	"./balbi"
		"log"
		"fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	balbi.ComplexCalc(0.00000001, 100000000)
	fmt.Fprintf(w, "Code Education Rocks")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Servidor ouvindo na porta 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

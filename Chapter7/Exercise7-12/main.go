// Exercise 7.12 Change the handler for /list to print its output as an HTML table, no text.
// you may find the html/template package useful.
package main

import (
	"log"
	"net/http"
)

func main() {
	db := database{m: map[string]int{"shoes": 50, "socks": 5}}
	http.HandleFunc("/update", db.updateItem)
	http.HandleFunc("/create", db.createItem)
	http.HandleFunc("/get", db.getItem)
	http.HandleFunc("/delete", db.deleteItem)
	http.HandleFunc("/list", db.listItems)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Exercise 7.11 Add aditional handlers so that clients can create, read, update and delete
// database entries. For example, a request of the form /update?item=socks&price=5 will update
// the price of an item in the inventory and report an error if the item does not exists or if
// the price is invalid. (Warning: this change introduces concurrent variable updates)
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

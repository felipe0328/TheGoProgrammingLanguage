package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"sync"
)

type database struct {
	sync.Mutex
	m map[string]int
}

func (db *database) createItem(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")

	if item == "" || price == "" {
		http.Error(w, "Please provide valid item and price elements", http.StatusBadRequest)
		return
	}

	db.Lock()
	_, ok := db.m[item]
	db.Unlock()

	if ok {
		msg := fmt.Sprintf("Element %s already exists", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	priceInt, err := strconv.Atoi(price)

	if err != nil {
		msg := fmt.Sprintf("Error parsing the price value: %v", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	db.Lock()
	db.m[item] = priceInt
	db.Unlock()

	fmt.Fprintf(w, "Created item %s:%d", item, db.m[item])
}

func (db *database) getItem(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")

	if item == "" {
		http.Error(w, "Please provide a valid item", http.StatusBadRequest)
		return
	}

	db.Lock()
	value, ok := db.m[item]
	db.Unlock()

	if !ok {
		msg := fmt.Sprintf("Element %s doesn't exists", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Item: %s\t Value:%d", item, value)
}

func (db *database) updateItem(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")

	if item == "" || price == "" {
		http.Error(w, "Please provide valid item and price elements", http.StatusBadRequest)
		return
	}

	db.Lock()
	_, ok := db.m[item]
	db.Unlock()

	if !ok {
		msg := fmt.Sprintf("Element %s doesn't exists", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	priceInt, err := strconv.Atoi(price)

	if err != nil {
		msg := fmt.Sprintf("Error parsing the price value: %v", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	db.Lock()
	db.m[item] = priceInt
	db.Unlock()

	fmt.Fprintf(w, "Updated item %s:%d", item, db.m[item])
}

func (db *database) deleteItem(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")

	if item == "" {
		http.Error(w, "Please provide a valid item", http.StatusBadRequest)
		return
	}

	db.Lock()
	_, ok := db.m[item]
	db.Unlock()

	if !ok {
		msg := fmt.Sprintf("Element %s doesn't exists", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	db.Lock()
	delete(db.m, item)
	db.Unlock()

	fmt.Fprintf(w, "Deleted item %s", item)
}

func (db *database) listItems(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("items").Parse(templateToUse))
	if err := t.Execute(w, db.m); err != nil {
		msg := fmt.Sprintf("Unable to show list %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
}

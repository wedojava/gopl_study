// go run gopl.io/ch1/fetch http://localhost:8000/list
// go run gopl.io/ch1/fetch http://localhost:8000/create?item=hat&price=30
// go run gopl.io/ch1/fetch http://localhost:8000/read?item=hat
// go run gopl.io/ch1/fetch http://localhost:8000/delete?item=hat
// go run gopl.io/ch1/fetch http://localhost:8000/list
// go run gopl.io/ch1/fetch http://localhost:8000/update?item=shoes&price=100
// go run gopl.io/ch1/fetch http://localhost:8000/list
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceGet := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceGet, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Parse price: %v.\n", err)
		return
	}
	if _, exist := db[item]; exist {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Item %s already exists.\n", item)
	} else {
		db[item] = dollars(price)
		fmt.Fprintf(w, "Successfully created item. %s: %s\n", item, db[item])
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceGet := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceGet, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Parse price: %v.\n", err)
		return
	}
	if _, ok := db[item]; ok {
		db[item] = dollars(price)
		fmt.Fprintf(w, "%s's price update to %s\n", item, db[item])
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "Item %s: %s\n", item, price)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		delete(db, item)
		fmt.Fprintf(w, "Successfully deleted item %s\n", item)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}

}

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

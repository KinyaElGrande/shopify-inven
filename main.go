package main

import (
	"net/http"
)

type inventory struct{}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := inventory{}
	generateHTML(w, data, "layout", "navbar", "index")
}

func createInventory(w http.ResponseWriter, r *http.Request){
	data := inventory{}
	generateHTML(w, data, "layout", "navbar", "create.inventory")
}

func viewInventory(w http.ResponseWriter, r *http.Request){
	data := inventory{}
	generateHTML(w, data, "layout", "navbar", "view.inventory")
}

func editInventory(w http.ResponseWriter, r *http.Request){
	data := inventory{}
	generateHTML(w, data, "layout", "navbar", "edit.inventory")
}

//Warehouses
func warehouses(w http.ResponseWriter, r *http.Request){
	data := inventory{}
	generateHTML(w, data, "layout", "navbar", "warehouses")
}

func createWarehouse(w http.ResponseWriter, r *http.Request){
	data := inventory{}
	generateHTML(w, data, "layout", "navbar", "create.warehouse")
}

func main() {
	mux := http.NewServeMux()

	//Serve static files
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", indexHandler)


	mux.HandleFunc("/inventory/new", createInventory)
	mux.HandleFunc("/inventory/view", viewInventory)
	mux.HandleFunc("/inventory/edit", editInventory)

	//warehouses
	mux.HandleFunc("/warehouses", warehouses)
	mux.HandleFunc("/create/warehouse", createWarehouse)

	http.ListenAndServe(":8000", mux)
}

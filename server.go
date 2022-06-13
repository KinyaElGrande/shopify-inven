package main

import (
	"mumbi/inven-logistics/controllers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	//Serve static files
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", controllers.IndexHandler)

	mux.HandleFunc("/inventory/new", controllers.NewInventory)
	mux.HandleFunc("/inventory/create", controllers.CreateInventory) //post

	mux.HandleFunc("/inventory/view", controllers.ViewInventory) //get
	mux.HandleFunc("/inventory/edit", controllers.EditInventory)
	mux.HandleFunc("/inventory/update", controllers.UpdateInventory)
	mux.HandleFunc("/inventory/delete", controllers.DeleteInventory)

	//warehouses
	mux.HandleFunc("/warehouses", controllers.Warehouses)
	mux.HandleFunc("/warehouse/new", controllers.NewWarehouse) //get request
	mux.HandleFunc("/warehouse/create", controllers.CreateWarehouse)

	http.ListenAndServe(":8000", mux)
}

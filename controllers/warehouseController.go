package controllers

import (
	"fmt"
	"mumbi/inven-logistics/data"
	"mumbi/inven-logistics/templates"
	"net/http"
	"strings"
)

//Warehouses
func Warehouses(w http.ResponseWriter, r *http.Request) {
	warehouses, _ := data.ListWareHouses()
	templates.GenerateHTML(w, &warehouses, "layout", "navbar", "warehouses")
}

func NewWarehouse(w http.ResponseWriter, r *http.Request) {
	data := data.Warehouse{}
	templates.GenerateHTML(w, data, "layout", "navbar", "create.warehouse")
}

func CreateWarehouse(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	city := r.PostFormValue("city")
	country := r.PostFormValue("country")
	manager := r.PostFormValue("manager")
	description := r.PostFormValue("description")
	code := strings.ToUpper(country[0:3] + city[0:4])

	warehouse := data.Warehouse{
		City:        city,
		Country:     country,
		Code:        code,
		Manager:     manager,
		Description: description,
	}

	_ = data.CreateWarehouse(warehouse)
	http.Redirect(w, r, "/warehouses", 302)
}
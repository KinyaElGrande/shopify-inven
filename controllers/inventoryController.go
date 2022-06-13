package controllers

import (
	"fmt"
	"mumbi/inven-logistics/data"
	"mumbi/inven-logistics/templates"
	"net/http"
	"strconv"
	"strings"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	inventories, _ := data.ListInventories()
	templates.GenerateHTML(w, &inventories, "layout", "navbar", "index")
}

func NewInventory(w http.ResponseWriter, r *http.Request) {
	warehouses, _ := data.ListWareHouses()
	templates.GenerateHTML(w, &warehouses, "layout", "navbar", "create.inventory")
}

//post create Inventory
func CreateInventory(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	name := r.PostFormValue("name")
	price, _ := strconv.Atoi(r.PostFormValue("price"))
	category := r.PostFormValue("category")
	quantity, _ := strconv.Atoi(r.PostFormValue("quantity"))
	manufacturer := r.PostFormValue("manufacturer")
	description := r.PostFormValue("description")

	inventory := data.Inventory{
		Name:         name,
		SKU:          name + manufacturer,
		Price:        price,
		Category:     category,
		Quantity:     quantity,
		Manufacturer: manufacturer,
		Description:  description,
		WarehouseID:  1,
	}

	_ = data.CreateInventory(inventory)

	http.Redirect(w, r, "/", 302)
}

func ViewInventory(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	id, _ := strconv.Atoi(vals.Get("id"))

	inventory, _ := data.GetInventory(id)
	templates.GenerateHTML(w, &inventory, "layout", "navbar", "view.inventory")
}

func EditInventory(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	id, _ := strconv.Atoi(vals.Get("id"))

	inventory, _ := data.GetInventory(id)
	warehouses, _ := data.ListWareHouses()

	type data map[string]interface{}

	resp := data{
		"inventoryID": id,
		"inventory":   inventory,
		"warehouses":  warehouses,
	}

	templates.GenerateHTML(w, &resp, "layout", "navbar", "edit.inventory")
}

func UpdateInventory(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	id, _ := strconv.Atoi(vals.Get("id"))

	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	name := r.PostFormValue("name")
	price, _ := strconv.Atoi(r.PostFormValue("price"))
	category := r.PostFormValue("category")
	quantity, _ := strconv.Atoi(r.PostFormValue("quantity"))
	manufacturer := r.PostFormValue("manufacturer")
	description := r.PostFormValue("description")
	warehouseId, _ := strconv.Atoi(r.PostFormValue("warehouse"))
	sku := strings.ToUpper(name[0:2] + manufacturer[0:2])

	inventory := data.Inventory{
		Name:         name,
		SKU:          sku,
		Price:        price,
		Category:     category,
		Quantity:     quantity,
		Manufacturer: manufacturer,
		Description:  description,
		WarehouseID:  warehouseId,
	}
	fmt.Println(id, inventory)

	_ = data.UpdateInventory(id, inventory)
	http.Redirect(w, r, "/", 302)
}

func DeleteInventory(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	id, _ := strconv.Atoi(vals.Get("id"))

	_ = data.DeleteInventory(id)
	http.Redirect(w, r, "/", 302)
}

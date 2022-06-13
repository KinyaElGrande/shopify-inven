package data

type Inventory struct {
	Name         string
	SKU          string
	Price        int
	Category     string
	Quantity     int
	Manufacturer string
	Description  string
	WarehouseID  int
}

type Warehouse struct {
	Location string
}

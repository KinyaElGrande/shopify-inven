package data

import (
	"context"
	"log"
	"time"
)

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

type FullInventory struct {
	ID           int
	Name         string
	SKU          string
	Category     string
	Price        int
	Quantity     int
	Manufacturer string
	Description  string
	Warehouse    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (inven *FullInventory) CreatedAtDate() string {
	return inven.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}
func (inven *FullInventory) UpdatedAtDate() string {
	return inven.UpdatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func GetInventory(id int) (inventory FullInventory, err error) {
	db := dbConn()
	defer db.Close()

	query := `SELECT  inventories.name, inventories.sku, inventories.category, inventories.price, inventories.quantity,
	inventories.manufacturer, inventories.description, warehouses.city, inventories.created_at, inventories.updated_at
	FROM inventories 
	JOIN warehouses ON inventories.warehouse_id = warehouses.id 
	WHERE inventories.id=?`

	row, err := db.Query(query, id)
	if err != nil {
		panic(err.Error())
	}

	for row.Next() {
		if err = row.Scan(&inventory.Name, &inventory.SKU, &inventory.Category, &inventory.Price,
			&inventory.Quantity, &inventory.Manufacturer, &inventory.Description, &inventory.Warehouse,
			&inventory.CreatedAt, &inventory.UpdatedAt); err != nil {
			log.Printf("Error %s when querying inventories", err)
		}
	}

	return inventory, nil
}

func ListInventories() (inventories []FullInventory, err error) {
	db := dbConn()
	defer db.Close()

	query := `SELECT inventories.id, inventories.name, inventories.sku, inventories.category, inventories.price, inventories.quantity,
	inventories.manufacturer, inventories.description, warehouses.city, inventories.created_at, inventories.updated_at
	FROM inventories 
	JOIN warehouses ON inventories.warehouse_id = warehouses.id`

	rows, _ := db.Query(query)

	inventory := FullInventory{}

	for rows.Next() {
		if err := rows.Scan(&inventory.ID, &inventory.Name, &inventory.SKU, &inventory.Category, &inventory.Price,
			&inventory.Quantity, &inventory.Manufacturer, &inventory.Description, &inventory.Warehouse,
			&inventory.CreatedAt, &inventory.UpdatedAt); err != nil {
			log.Printf("Error %s when querying inventories", err)
		}

		inventories = append(inventories, inventory)
	}

	return inventories, nil

}

func CreateInventory(inv Inventory) error {
	db := dbConn()
	defer db.Close()

	query := "INSERT INTO inventories(name,sku,price,category,quantity,manufacturer,description,warehouse_id) VALUES (?,?,?,?,?,?,?,?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, inv.Name, inv.SKU, inv.Price, inv.Category, inv.Quantity, inv.Manufacturer, inv.Description, inv.WarehouseID)
	if err != nil {
		log.Printf("Error %s when inserting row into inventories table", err)
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d Inventory created ", rows)
	return nil
}

func UpdateInventory(invID int, inv Inventory) error {
	db := dbConn()
	defer db.Close()

	query := "UPDATE inventories SET name=?,sku=?,price=?,category=?,quantity=?,manufacturer=?,description=?,warehouse_id=? WHERE id=?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing Inventory Update SQL statement", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, inv.Name, inv.SKU, inv.Price, inv.Category, inv.Quantity, inv.Manufacturer, inv.Description, inv.WarehouseID, invID)
	if err != nil {
		log.Printf("Error %s when inserting row into inventories table", err)
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d Inventory updated ", rows)
	return nil
}

func DeleteInventory(invID int) error {
	db := dbConn()
	defer db.Close()

	query := "DELETE FROM inventories WHERE id=?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, invID)
	if err != nil {
		log.Printf("Error %s when inserting row into inventories table", err)
		return err
	}
	return nil
}

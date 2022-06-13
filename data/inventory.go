package data

import (
	"context"
	"log"
	"time"
)

type Inventory struct {
	Name string
	SKU string
	Price int
	Category string
	Quantity int
	Manufacturer string
	Description string
	WarehouseID int
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




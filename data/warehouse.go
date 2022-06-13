package data

import (
	"context"
	"log"
	"time"
)

type Warehouse struct {
	City        string
	Country     string
	Code        string
	Manager     string
	Description string
}

func CreateWarehouse(house Warehouse) error {
	db := dbConn()
	defer db.Close()

	query := "INSERT INTO warehouses(city,country,code,manager,description) VALUES (?,?,?,?,?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, house.City, house.Country, house.Code, house.Manager, house.Description)
	if err != nil {
		log.Printf("Error %s when inserting row into warehouses table", err)
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

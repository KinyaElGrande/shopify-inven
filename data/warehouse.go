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

type FullWarehouse struct {
	ID          int
	City        string
	Country     string
	Code        string
	Manager     string
	Description string
	CreatedAt   time.Time
}

// format the CreatedAt date 
func (w *FullWarehouse) CreatedAtDate() string {
	return w.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func ListWareHouses() (warehouses []FullWarehouse, err error) {
	db := dbConn()
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM warehouses ORDER BY id DESC")

	wareHouse := FullWarehouse{}

	for rows.Next() {
		if err := rows.Scan(&wareHouse.ID, &wareHouse.City, &wareHouse.Country, &wareHouse.Code,
			&wareHouse.Manager, &wareHouse.Description, &wareHouse.CreatedAt); err != nil {
			log.Printf("Error %s when querying warehouses", err)
		}

		warehouses = append(warehouses, wareHouse)
	}

	return warehouses, nil
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

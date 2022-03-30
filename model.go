package main

import (
	"database/sql"
	"fmt"
)

// Datastore implements the Repository interface
type Datastore struct {
	db *sql.DB
}

type Inventory struct {
	Id          int `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Status      bool `json:"status"`
}

func (ds *Datastore) GetInventory(id int) (Inventory, error) {
	var inventory Inventory
	err := ds.db.QueryRow(`SELECT id, name, description,price,status FROM "inventory"
		WHERE id=$1`, id).Scan(&inventory.Id, &inventory.Name, &inventory.Description, &inventory.Price, &inventory.Status)
	return inventory, err
}

func (ds *Datastore) ListInventory() ([]Inventory, error) {
	var inventoryList []Inventory
	rows, err := ds.db.Query(`SELECT id, name, description,price,status FROM "inventory"`)
	for rows.Next() {
		inventory := Inventory{}
		err = rows.Scan(&inventory.Id, &inventory.Name, &inventory.Description, &inventory.Price, &inventory.Status)
		if err != nil {
			return nil, err
		}
		inventoryList = append(inventoryList, inventory)
	}
	return inventoryList, err
}

func (ds *Datastore) DeleteInventory(id int) error {
	_,err := ds.db.Query(`Delete from inventory where id=$1`,id)
	return err
}

func (ds *Datastore) CreateInventory(inventory Inventory) (Inventory, error) {
	err := ds.db.QueryRow(`INSERT INTO "inventory" (name, description, price, status) VALUES ($1, $2, $3, $4) RETURNING id`,
		inventory.Name, inventory.Description, inventory.Price, inventory.Status).Scan(&inventory.Id)
	return inventory, err
}

func (ds *Datastore) UpdateInventory(inventory Inventory) (Inventory, error) {
	panic("implement")
}

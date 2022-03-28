package main

import (
	"database/sql"
)

// Datastore implements the Repository interface
type Datastore struct {
	db *sql.DB
}

type Inventory struct {
	Id          int
	Name        string
	Description string
	Price       int
	Status      bool
}

func (ds *Datastore) GetInventory(id int) (Inventory, error) {
	var inventory Inventory
	err := ds.db.QueryRow(`SELECT id, name, description,price,status FROM "inventory"
		WHERE id=$1`, id).Scan(&inventory.Id, &inventory.Name, &inventory.Description, &inventory.Price, &inventory.Status)
	return inventory, err
}

func (ds *Datastore) ListInventory() ([]Inventory, error) {
	panic("implement")
}

func (ds *Datastore) DeleteInventory(id int) error {
	panic("implement")
}

func (ds *Datastore) CreateInventory(inventory Inventory) (Inventory, error) {
	panic("implement")
}

func (ds *Datastore) UpdateInventory(inventory Inventory) (Inventory, error) {
	panic("implement")
}

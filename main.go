package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func HandleGetInventory(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world\n"))
}

func main() {
	//database : hippo
	//host:hippo-primary.testing.svc
	//password:*+fVs0i<f@i[@<JM*KSuYn1B
	//port:5432
	//user: hippo
	// TODO: replace with the connection string
	db, err := sql.Open("pgx", "postgres://hippo:%2A%2BfVs0i%3Cf%40i%5B%40%3CJM%2AKSuYn1B@localhost:5432/hippo")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()
	var greeting string
	err = db.QueryRow("select 1").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
	datastore := Datastore{db}
	inventory, err := datastore.GetInventory(1)
	fmt.Println(inventory)

	router := mux.NewRouter()

	router.HandleFunc("/api/inventory/{id}", HandleGetInventory).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":8080")
}

package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"encoding/json"
	_ "github.com/jackc/pgx/v4/stdlib"
)
var db *sql.DB
func HandleGetInventory(w http.ResponseWriter, r *http.Request) {
	datastore := Datastore{db}
	inventory, err := datastore.GetInventory(1)
	fmt.Println(inventory)
	response,err:=json.Marshal(inventory)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func main() {
	//database : hippo
	//host:hippo-primary.testing.svc
	//password:*+fVs0i<f@i[@<JM*KSuYn1B
	//port:5432
	//user: hippo
	// TODO: replace with the connection string
	var err error
	db, err = sql.Open("pgx", "postgres://hippo:%2A%2BfVs0i%3Cf%40i%5B%40%3CJM%2AKSuYn1B@localhost:5432/hippo")
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
	

	router := mux.NewRouter()

	router.HandleFunc("/api/inventory/{id}", HandleGetInventory).Methods("GET")
	
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":8080")
}

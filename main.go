package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"encoding/json"
	"github.com/baijum/servicebinding/binding"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/urfave/negroni"
)

var db *sql.DB

func HandleGetInventory(w http.ResponseWriter, r *http.Request) {
	datastore := Datastore{db}
	vars := mux.Vars(r)
	id,_:= strconv.Atoi(vars["id"])
	inventory, err := datastore.GetInventory(id)
	fmt.Println(inventory)
	response, err := json.Marshal(inventory)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func HandleCreateInventory(w http.ResponseWriter, r *http.Request) {
	datastore := Datastore{db}
	var req Inventory
	json.NewDecoder(r.Body).Decode(&req)
	inventory, err := datastore.CreateInventory(req)
	response, err := json.Marshal(inventory)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(response))
}

func HandleListInventory(w http.ResponseWriter, r *http.Request) {
	datastore := Datastore{db}
	inventory, err := datastore.ListInventory()
	fmt.Println(inventory)
	response, err := json.Marshal(inventory)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func HandleDeleteInventory(w http.ResponseWriter, r *http.Request) {
	datastore := Datastore{db}
	vars := mux.Vars(r)
	id,_:= strconv.Atoi(vars["id"])
	err := datastore.DeleteInventory(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}


func main() {
	//database : hippo
	//host:hippo-primary.testing.svc
	//password:*+fVs0i<f@i[@<JM*KSuYn1B
	//port:5432
	//user: hippo
	// TODO: replace with the connection string
	var err error
	fmt.Fprintln(os.Stderr, "Starting of main")
	sb, err := binding.NewServiceBinding()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Could not read service bindings")
	}
	bindings, err := sb.Bindings("postgresql")
	fmt.Fprintln(os.Stderr,bindings)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Unable to find postgres binding")
	}
	connectionString := bindings[0]["pgbouncer-uri"]
	fmt.Println(connectionString)
	fmt.Fprintln(os.Stderr,connectionString)
	
	db, err = sql.Open("pgx", connectionString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	defer db.Close()
	var greeting string
	err = db.QueryRow("select 1").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	fmt.Println(greeting)

	router := mux.NewRouter()

	router.HandleFunc("/api/inventory/{id}", HandleGetInventory).Methods("GET")
	router.HandleFunc("/api/inventory", HandleCreateInventory).Methods("POST")
	router.HandleFunc("/api/inventory", HandleListInventory).Methods("GET")
	router.HandleFunc("/api/inventory/{id}", HandleDeleteInventory).Methods("DELETE")
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":8080")

}

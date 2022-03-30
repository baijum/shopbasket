package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"encoding/json"
	"github.com/nebhale/client-go/bindings"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/urfave/negroni"
	//"time"
	"io/ioutil"
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

func IOReadDir(root string) ([]string, error) {
    var files []string
    fileInfo, err := ioutil.ReadDir(root)
    if err != nil {
        return files, err
    }

    for _, file := range fileInfo {
        files = append(files, file.Name())
    }
    return files, nil
}
func main() {
	//database : hippo
	//host:hippo-primary.testing.svc
	//password:*+fVs0i<f@i[@<JM*KSuYn1B
	//port:5432
	//user: hippo
	// TODO: replace with the connection string
	//time.Sleep(1 * time.Minute)
	var err error
	// fmt.Fprintln(os.Stderr, "Starting of main")
	// sb, err := binding.NewServiceBinding()
	// if err != nil {
	// 	_, _ = fmt.Fprintln(os.Stderr, "Could not read service bindings")
	// }
	b := bindings.FromServiceBindingRoot()
	b = bindings.Filter(b, "postgresql")
	if len(b) != 1 {
		_, _ = fmt.Fprintf(os.Stderr, "Incorrect number of PostgreSQL drivers: %d\n", len(b))
		os.Exit(1)
	}
	// res,_:=IOReadDir("/bindings")
	// fmt.Fprintln(res)
	connectionString, ok := bindings.Get(b[0], "pgbouncer-uri")
	if !ok {
		_, _ = fmt.Fprintln(os.Stderr, "No pgbouncer-uri in binding")
		os.Exit(1)
	}
	// fmt.Println(res)
	// bindings, err := sb.Bindings("postgresql")
	// fmt.Fprintln(os.Stderr,bindings)
	// if err != nil {
	// 	_, _ = fmt.Fprintln(os.Stderr, "Unable to find postgres binding")
	// }
	// connectionString := bindings[0]["pgbouncer-uri"]
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

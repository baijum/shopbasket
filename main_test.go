package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gorilla/mux"
)

// NewTestServer helper for testing
func NewTestServer(path, method string, fn func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	hts := httptest.NewServer(func() *mux.Router {
		rt := mux.NewRouter()
		rt.HandleFunc(path, fn).Methods(method)
		return rt
	}())
	return hts
}

func TestIntegration(t *testing.T) {
	_, urt := Route()
	ts := httptest.NewServer(urt)
	defer ts.Close()
	dbtest, err := InitializeDB()
	fmt.Println(err)
	datastore := Datastore{dbtest}
	fmt.Println(err)
	inv := &Inventory{
		Name:        "Cars",
		Description: "Automobile",
		Price:       "4000",
		Status:      true,
	}
	resultInventory, err := datastore.CreateInventory(*inv)
	assert.Nil(t, err)
	_, err = datastore.GetInventory(resultInventory.Id)
	assert.Nil(t, err)
	list, err := datastore.ListInventory()
	fmt.Println(list)
	assert.NotNil(t, list)
	assert.Nil(t, err)
	err = datastore.DeleteInventory(resultInventory.Id)
	assert.Nil(t, err)
}

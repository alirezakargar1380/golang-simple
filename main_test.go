package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alirezakargar1380/test/calculator"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	total := calculator.Add(1, 2)
	assert.NotNil(t, total, "The `total` variable should not be nil")
	assert.Equalf(t, 3, total, "1 + 2 should equal 3")
}

func Router() *mux.Router {
	Router := mux.NewRouter()
	Router.HandleFunc("/", RootEndpoint).Methods("GET")
	return Router
}

func TestHomeRouter(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	res := httptest.NewRecorder()
	Router().ServeHTTP(res, req)
	fmt.Println(res.Code)
	fmt.Println(res.Body)
	fmt.Println(res)
}

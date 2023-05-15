package main

import (
	"fmt"
	"net/http"

	"github.com/agustfricke/crud-go/db"
	"github.com/agustfricke/crud-go/models"
	"github.com/agustfricke/crud-go/routes"
	"github.com/gorilla/mux"
)

func main() {

  db.DBConnection()
  db.DB.AutoMigrate(models.Item{})

  r := mux.NewRouter()

  r.HandleFunc("/", routes.HomeHandler)
  r.HandleFunc("/items/", routes.GetItems).Methods("GET")
  r.HandleFunc("/items/{id}/", routes.GetItem).Methods("GET")
  r.HandleFunc("/items/", routes.PostItem).Methods("POST")
  r.HandleFunc("/items/{id}/", routes.PutItem).Methods("PUT")
  r.HandleFunc("/items/{id}/", routes.DeleteItem).Methods("DELETE")

  http.ListenAndServe(":8000", r)
  fmt.Println("Server running on port 8000")

}

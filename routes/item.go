package routes

import (
	"encoding/json"
	"net/http"

	"github.com/agustfricke/crud-go/db"
	"github.com/agustfricke/crud-go/models"
	"github.com/gorilla/mux"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
  var items []models.Item
  db.DB.Find(&items)
  json.NewEncoder(w).Encode(&items)
}

func GetItem(w http.ResponseWriter, r *http.Request) { 
  params := mux.Vars(r)
  var item models.Item
  db.DB.First(&item, params["id"])
  if item.ID == 0 {
    w.WriteHeader(http.StatusNotFound)
    return
  }
  json.NewEncoder(w).Encode(&item)
}

func PostItem(w http.ResponseWriter, r *http.Request) {
  var item models.Item
  json.NewDecoder(r.Body).Decode(&item)
  itemCreated := db.DB.Create(&item)
  err := itemCreated.Error
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  json.NewEncoder(w).Encode(&item)
}

func PutItem(w http.ResponseWriter, r *http.Request) { 
  params := mux.Vars(r)
  var item models.Item
  db.DB.First(&item, params["id"])
  if item.ID == 0 {
    w.WriteHeader(http.StatusNotFound)
    return
  }
  json.NewDecoder(r.Body).Decode(&item)
  itemUpdated := db.DB.Save(&item)
  err := itemUpdated.Error
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  json.NewEncoder(w).Encode(&item)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  var item models.Item
  db.DB.First(&item, params["id"])
  if item.ID == 0 {
    w.WriteHeader(http.StatusNotFound)
    return
  }
  db.DB.Unscoped().Delete(&item)
  w.WriteHeader(http.StatusOK)
}

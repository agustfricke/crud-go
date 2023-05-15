package routes

import (
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	message := `REST with Go! 

GET /items/ - Get all items
GET /items/{id}/ - Get item by id
POST /items/ - Create item
PUT /items/{id}/ - Update item by id
DELETE /items/{id}/ - Delete item by id`

	w.Write([]byte(message))
}

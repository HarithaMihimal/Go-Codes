package api

import (
	"net/http"
	"github.com/google/mux"
	"github.com/google/uuid"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Item struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
}

type Server struct {
	*mux.Router

	shoppingItems []Item

}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
		shoppingItems: []Item{},
	}
 
	s.routes()
	
	return s
}


func (s *Server) routes() {
	s.Router.HandleFunc("/shopping-items", s.listShoppingItems()).Methods("GET")
	s.Router.HandleFunc("/shopping-items", s.createShoppingItem()).Methods("POST")
	s.Router.HandleFunc("/shopping-items/{id}", s.removeShoppingItems()).Methods("DELETE")
}

func (s *Server) listShoppingItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(w).Encode(s.shoppingItems); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
func (s *Server) createShoppingItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i Item
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) removeShoppingItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr, _ := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			
		}
		for i, item := range s.shoppingItems {
			if item.ID == id {
				s.shoppingItems = append(s.shoppingItems[:i], s.shoppingItems[i+1:]...)
				break
			}
		}
	}
}
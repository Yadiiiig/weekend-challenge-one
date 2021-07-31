package router

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type Collection struct {
	Router *mux.Router
	DB     *sqlx.DB
}

func InitRouter(db *sqlx.DB) *Collection {
	return &Collection{
		Router: mux.NewRouter().StrictSlash(true),
		DB:     db,
	}
}

func (c *Collection) InitRoutes() {
	c.Router.HandleFunc("/get_help", c.GetHelp).Methods("GET")
}

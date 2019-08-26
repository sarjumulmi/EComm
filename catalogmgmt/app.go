package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/sarjumulmi/ecomm/catalogmgmt/models"
	"github.com/sarjumulmi/ecomm/catalogmgmt/utils"
)

// App ...
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/products", a.getProducts).Methods("GET")
}

// Initialize DB
func (a *App) Initialize(user, pwd, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, pwd, dbname)
	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := models.GetProducts(a.DB)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, products)
}

// Run server
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/sarjumulmi/ecomm/catalogmgmt/models"
	"github.com/sarjumulmi/ecomm/catalogmgmt/utils"
)

func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := models.GetProducts(a.DB)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, products)
}

func (a *App) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID, err := strconv.Atoi(vars["productId"])
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "invalid product id")
		return
	}
	p := models.Product{ProductID: productID}
	if err := p.GetProduct(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.RespondWithError(w, http.StatusNotFound, "Product not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, p)
}

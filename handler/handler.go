package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/arjunmalhotra1/Grocery-Tech-Exercise/storage/krogermart"
	"github.com/arjunmalhotra1/Grocery-Tech-Exercise/store"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// TODO this has not been used yet.
const (
	getItemErr         = "Bad Request. Item for code %s couldn't be found"
	internalErr        = "Internal server error"
	noInventoryMessage = "No inventory entered yet."
)

type handler struct {
	Router *chi.Mux
	store  store.Store
}

func New() *handler {
	var kMart krogermart.KrogerMart
	krogerStore := store.NewStore(&kMart)

	var h handler
	h.store = krogerStore
	r := chi.NewMux()
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", http.StatusNotFound)
	})
	r.Get("/get-inventory", h.getInventory)
	r.Get("/item/{code}", h.getItem)
	r.Post("/item/", h.createItem)
	r.Delete("/item/{code}", h.deleteItem)
	h.Router = r
	return &h
}

func (h *handler) getInventory(res http.ResponseWriter, req *http.Request) {
	var inventory []store.Item = h.store.GetInventory()
	//inventory := h.store.GetInventory()
	if len(inventory) == 0 {
		render.JSON(res, req, noInventoryMessage)
		return
	}
	render.JSON(res, req, inventory)
}

func (h *handler) getItem(res http.ResponseWriter, req *http.Request) {
	pCode := chi.URLParam(req, "code")
	item, ifFound := h.store.Get(pCode)
	if !ifFound {
		render.Status(req, http.StatusNotFound)
		render.JSON(res, req, fmt.Sprintf(getItemErr, pCode))
		return

	}
	render.JSON(res, req, item)

}

func (h *handler) createItem(res http.ResponseWriter, req *http.Request) {
	var items []store.Item
	if err := json.NewDecoder(req.Body).Decode(&items); err != nil {
		fmt.Println("Json decoding error")
	}
	err := h.store.Post(items)
	if err != nil {
		log.Println("Error in Post request.")
		render.Status(req, http.StatusBadRequest)
		render.JSON(res, req, fmt.Sprintf("%+v", err))
		return
	}
	res.WriteHeader(http.StatusCreated)
	//render.JSON(res, req, fmt.Sprintf("%+v", items))
	render.JSON(res, req, items)
}

func (h *handler) deleteItem(res http.ResponseWriter, req *http.Request) {
	pCode := chi.URLParam(req, "code")
	h.store.Delete(pCode)
}

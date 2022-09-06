package store

import (
	"log"
)

type Item struct {
	Name        string  `json:"Name" validate:"required"`
	Price       float64 `json:"Price" validate:"required"`
	ProduceCode string  `json:"ProduceCode" validate:"required"`
}

type StoreAccessor interface {
	GetInventory() []Item
	GetItem(string) *Item
	CreateItem([]Item)
	DeleteItem(string)
}

type Store struct {
	s StoreAccessor
}

func NewStore(s StoreAccessor) Store {
	return Store{
		s: s,
	}
}

func (st Store) Get(pCode string) (*Item, bool) {
	log.Println("Get request for", pCode)
	item := st.s.GetItem(pCode)
	if item == nil {
		return nil, false
	}
	return item, true

}

func (st Store) Post(items []Item) error {
	err := isValidItem(items)
	if err != nil {
		return err
	}
	log.Println("Post request for", items)
	// This will catch an error if error while creating in the database.
	st.s.CreateItem(items)
	return nil
}

func (st Store) GetInventory() []Item {
	log.Println("Get all inventory request")
	inventory := st.s.GetInventory()
	return inventory
}

func (st Store) Delete(pCode string) {
	log.Println("Delete request for", pCode)
	st.s.DeleteItem(pCode)
}

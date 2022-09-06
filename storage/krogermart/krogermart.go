package krogermart

import (
	"strings"
	"sync"

	"github.com/arjunmalhotra1/Grocery-Tech-Exercise/store"
)

type KrogerMart []store.Item

var mu sync.Mutex

func (kr *KrogerMart) GetInventory() []store.Item {
	return *kr
}

func (kr *KrogerMart) GetItem(pCode string) *store.Item {
	// If present in cache then don't traverse.
	for _, v := range *kr {
		if strings.Compare(pCode, v.ProduceCode) == 0 {
			return &v

		}
	}
	return nil
}

func (kr *KrogerMart) CreateItem(items []store.Item) {
	mu.Lock()
	*kr = append(*kr, items...)
	mu.Unlock()
}

func (kr *KrogerMart) DeleteItem(pCode string) {
	mu.Lock()
	var index int = -1
	var items []store.Item
	for i, v := range *kr {
		if strings.Compare(pCode, v.ProduceCode) == 0 {
			index = i
			break
		}
	}
	if index != -1 {
		items = *kr
		items = append(items[:index], items[index+1:]...)
		*kr = items
	}
	mu.Unlock()
}

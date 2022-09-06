package krogermart

import (
	"testing"

	"github.com/arjunmalhotra1/Grocery-Tech-Exercise/store"
	"github.com/stretchr/testify/assert"
)

func TestGetItemEmpty(t *testing.T) {
	var testKrogerMart KrogerMart
	item := testKrogerMart.GetItem("12345")
	assert.Empty(t, item, "TestGetItem failed on empty store.")
}

func TestGetInventory(t *testing.T) {
	var testKrogerMart KrogerMart
	pCode := "1234-1234-1234-1234"
	i := store.Item{
		Name:        "Potato",
		Price:       12.24,
		ProduceCode: pCode,
	}
	i2 := store.Item{
		Name:        "Potato",
		Price:       12.24,
		ProduceCode: pCode,
	}
	testKrogerMart = append(testKrogerMart, i)
	testKrogerMart = append(testKrogerMart, i2)

	items := testKrogerMart.GetInventory()
	assert.Equal(t, len(items), 2, "TestGetInventory failed.")

}

func TestGetItemGood(t *testing.T) {
	var testKrogerMart KrogerMart
	pCode := "1234-1234-1234-1234"
	i := store.Item{
		Name:        "Potato",
		Price:       12.24,
		ProduceCode: pCode,
	}
	testKrogerMart = append(testKrogerMart, i)
	item := testKrogerMart.GetItem(pCode)
	assert.NotEmpty(t, item, "TestGetItemGood item not empty.")
	assert.Equal(t, *item, i, "TestGetItemGood failed.")
}

func TestDeleteItem(t *testing.T) {
	var testKrogerMart KrogerMart
	pCode := "1234-1234-1234-1234"
	i := store.Item{
		Name:        "Potato",
		Price:       12.24,
		ProduceCode: pCode,
	}
	testKrogerMart = append(testKrogerMart, i)
	testKrogerMart.DeleteItem(pCode)
	item := testKrogerMart.GetItem(pCode)
	assert.Empty(t, item, "TestDeleteItem item not deleted.")

}

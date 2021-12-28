package facade

import (
	"testing"
)

var dataStore *DataStore

func init()  {
	dataStore = New()
}

func TestDataStore_Create(t *testing.T) {
	id:= 1
	dataStore.Create(id)
}

func TestDataStore_Update(t *testing.T) {
	id:= 1
	dataStore.Update(id)
}

func TestDataStore_Delete(t *testing.T) {
	id :=1
	dataStore.Delete(id)
}

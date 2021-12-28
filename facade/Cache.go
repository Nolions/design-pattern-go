package facade

import (
	"fmt"
	"strconv"
)

type Cache struct {
}

func (c *Cache) Create(id int) {
	fmt.Printf("Create id is %s's data to Cache\n", strconv.Itoa(id))
}

func (c *Cache) Update(id int) {
	fmt.Printf("Update id is %s's data from Cache\n", strconv.Itoa(id))
}

func (c *Cache) Delete(id int) {
	fmt.Printf("Delete id is %s's data from Cache\n", strconv.Itoa(id))
}

package facade

import (
	"fmt"
	"strconv"
)

type DB struct {
}

func (d *DB) Create(id int) {
	fmt.Printf("Create id is %s's data to DB\n", strconv.Itoa(id))
}

func (d *DB) Update(id int) {
	fmt.Printf("Update id is %s's data from DB\n", strconv.Itoa(id))
}

func (d *DB) Delete(id int) {
	fmt.Printf("Delete id is %s's data from DB\n", strconv.Itoa(id))
}

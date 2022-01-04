package adapter

import "fmt"

type computer interface {
	insertIntoLightningPort()
}

type client struct {
}

func (c *client) insertLightningConnectorIntoComputer(com computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.insertIntoLightningPort()
}

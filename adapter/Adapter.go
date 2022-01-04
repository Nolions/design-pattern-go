package adapter

import "fmt"

type Adapter struct {
	win *windows
}

func (w *Adapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.win.insertIntoUSBPort()
}

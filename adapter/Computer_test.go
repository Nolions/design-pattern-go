package adapter

import "testing"

var c *client

func init() {
	c = &client{}
}

func TestMac(t *testing.T) {
	mac := &mac{}

	c.insertLightningConnectorIntoComputer(mac)
}

func TestWin(t *testing.T) {
	windowsMachine := &windows{}
	windowsMachineAdapter := &Adapter{
		win: windowsMachine,
	}

	c.insertLightningConnectorIntoComputer(windowsMachineAdapter)
}

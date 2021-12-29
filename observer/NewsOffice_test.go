package observer

import "testing"

func TestNotifyCustomerA(t *testing.T) {
	cA := &CustomerA{
		name: "Chen",
	}
	cB := &CustomerA{
		name: "Lin",
	}

	o := newNewsOffice()

	o.addCustomer(cA)
	o.addCustomer(cB)
	o.printingCompleted()
}

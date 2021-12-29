package observer

type NewsOffice struct {
	customers []Customer
}

func newNewsOffice() *NewsOffice {
	return &NewsOffice{}
}

func (n *NewsOffice) addCustomer(customer Customer) {
	n.customers = append(n.customers, customer)
}

func (n *NewsOffice) printingCompleted() {
	n.notifyAllCustomer()
}

func (n *NewsOffice) notifyAllCustomer()  {
	for _, customer := range n.customers {
		customer.update()
	}
}
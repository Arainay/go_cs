package chainOfResponsibility

import "fmt"

type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment is done")

	} else {
		fmt.Println("Cashier is getting money from the patient")
	}
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}

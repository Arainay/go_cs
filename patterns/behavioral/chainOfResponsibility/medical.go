package chainOfResponsibility

import "fmt"

type Medical struct {
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine is already given to the patient")
		m.next.execute(p)

		return
	}

	fmt.Println("Medical is giving medicine to the patient")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *Medical) setNext(next Department) {
	m.next = next
}

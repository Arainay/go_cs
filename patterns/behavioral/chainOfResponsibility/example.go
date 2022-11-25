package chainOfResponsibility

func Example() {
	cashier := &Cashier{}
	medical := &Medical{}
	doctor := &Doctor{}
	reception := &Reception{}
	patient := &Patient{name: "John"}

	medical.setNext(cashier)
	doctor.setNext(medical)
	reception.setNext(doctor)

	reception.execute(patient)
}

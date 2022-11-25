package chainOfResponsibility

type Department interface {
	execute(*Patient)
	setNext(Department)
}

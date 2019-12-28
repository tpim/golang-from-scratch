package object

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct{}

//Dock Type
func (p *GoProgrammer) WriteHelloWorld() string {
	return "OK"
}

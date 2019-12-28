package object

import (
	"testing"
)

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bod", 20}
	t.Log(e.String())
}

func TestInterface(t *testing.T) {
	var e Programmer = new(GoProgrammer)
	t.Log(e.WriteHelloWorld())
}

func say(n int) int {
	return n
}

func TestCustimerType(t *testing.T) {
	p := timeSpent(say)
	p(1)
}

package pointer

import (
	"fmt"
)

type myType struct {
	value int
}

func newMyType(value int) *myType {
	return &myType{ value }
}

func (obj myType) m1() {
	obj.value = 1
}

func (obj *myType) m2() {
	obj.value = 2
}

func UsePointer() {
	mt := newMyType(0)
	fmt.Printf("mt -> %d\n", mt.value)
	mt.m1()
	fmt.Printf("mt -> %d\n", mt.value)
	mt.m2()
	fmt.Printf("mt -> %d\n", mt.value)

}

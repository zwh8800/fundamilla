package pointer

import (
	"fmt"
)

// 指针的作用, 大概只剩下区别 ByVal 和 ByRef
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

// slice本身就是引用传递
func manSlice1(s []int) {
	s[0] = 100
}

// 所以不用传指针,
//func manSlice2(s *[]int) {
//	s[0] = 1000	//error
//}

func UsePointer() {
	mt := newMyType(0)
	fmt.Printf("mt -> %d\n", mt.value)
	mt.m1()
	fmt.Printf("mt -> %d\n", mt.value)
	mt.m2()
	fmt.Printf("mt -> %d\n", mt.value)


	// slice测试
	s := make([]int, 2)
	s[0] = 10	//正确的姿势是使用append
	s[1] = 20
	//s[2] = 30	//will panic
	s = append(s, 30)	//use append
	//除了append, 还有len, cap, copy函数可以用

	fmt.Printf("s[0] -> %d\n", s[0])
	manSlice1(s)
	fmt.Printf("s[0] -> %d\n", s[0])
	//manSlice2(s)
	fmt.Printf("s[0] -> %d\n", s[0])

	// 不是复制, 而是引用, 所以修改s, s2也会改变
	s2 := s[1:2]
	s[1] = 2000
	fmt.Printf("s2[0] -> %d\n", s2[0])	// -> 2000

	// 正确的姿势是使用copy
	//s3 := make([]int, 2)	//和下一句类似
	s3 := []int{0, 0}
	copy(s3, s[1:2])
	s[1] = 20000
	fmt.Printf("s2[0] -> %d\n", s2[0])	// -> 20000
	fmt.Printf("s3[0] -> %d\n", s3[0])	// -> 2000

}

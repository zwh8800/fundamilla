package usetype
import "fmt"

//type的几种用法

type t1 struct {
	myPrivateData int
}
func (obj *t1) getPrivateDataPlusOne() int {
	return obj.myPrivateData + 1
}

type t2 map[string]*t1

type t3 map[string] t1

type t4 chan t1

type t5 interface {
	getPrivateDataPlusOne() int
}

func useInterface(t t5) {
	fmt.Printf("t.p+1 -> %d\n", t.getPrivateDataPlusOne())

	fmt.Printf("t -> %#v\n", t)
}

func UseType() {
	mt1 := t1{ 1 }
	mt2 := t1{ 2 }
	mt3 := t1{ 3 }

	mapt1 := make(t2)
	mapt2 := make(t3)

	mapt1["h1"] = &mt1
	mapt1["h2"] = &mt2
	mapt1["h3"] = &mt3
	mapt2["h1"] = mt1
	mapt2["h2"] = mt2
	mapt2["h3"] = mt3

	//useInterface(mt1)
	//useInterface(mt3)
	useInterface(&mt3)
	useInterface(mapt1["h2"])
	fmt.Printf("mapt1 -> %#v\n", mapt1)
	fmt.Printf("mapt1[\"h2d\"] -> %#v\n", mapt1["h2d"])
	useInterface(mapt1["h2d"])
	//useInterface(mapt2["h3"])
	//useInterface(mapt2["h3d"])
}

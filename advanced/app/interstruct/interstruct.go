package interstruct

import "fmt"

type Customer struct {
	name    string
	address string
	bal     float64
}

type Rectangle struct {
	length, height float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.height
}

func Main() {
	var tS Customer
	tS.name = "Stefan Einarsson"
	tS.address = "Hellway"
	tS.bal = 123.123
	fmt.Printf("%s lives in %s at %f\n", tS.name, tS.address, tS.bal)

	rect1 := Rectangle.Area(Rectangle{5, 5})
	rect2 := Rectangle{10, 10}
	fmt.Println(rect1)
	fmt.Println(rect2.Area())
}

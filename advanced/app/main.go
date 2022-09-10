package main

import (
	definedType "example/project/definedType"
	generics "example/project/generics"
	interstruct "example/project/interstruct"
	stuff "example/project/mypackage"
	"fmt"
	"reflect"
)

func callDefinedType() {
	definedType.Main()
}
func callInterStruct() {
	interstruct.Main()
}

func callMyPackage() {
	fmt.Println("Hello", stuff.Name)
	intArr := []int{2, 3, 5, 7, 11}
	strArr := stuff.IntArrToStrArr(intArr)
	fmt.Println(strArr)
	fmt.Println(reflect.TypeOf(strArr))
}

func callMyConstraints() {
	fmt.Println(generics.GetSumGen(10, 15))
	fmt.Println(generics.GetSumGen(10.5, 15.9))
}

func main() {
	callMyPackage()
	callMyConstraints()
	callInterStruct()
	callDefinedType()
}

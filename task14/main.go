package main

import (
	"fmt"
	"reflect"
)

func printType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		rt := reflect.TypeOf(v)

		if rt.Kind() == reflect.Chan {
			fmt.Println("chan")
		} else {
			fmt.Println("unknown type")
		}
	}
}

func main() {
	printType(1)
	printType("str")
	printType(true)
	printType(make(chan int))
	printType(2.1)
}

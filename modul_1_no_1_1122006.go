package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	string1 := "this is string 1"           // menggunakan symbol :=
	var string2 string = "this is string 2" // menggunakan var
	var string3 = "this is string 3"        // menggunakan var tanpa mendefinisikan tipenya

	fmt.Println(string1, ", type: ", reflect.TypeOf(string1))
	fmt.Println(string2, ", type: ", reflect.TypeOf(string2))
	fmt.Println(string3, ", type: ", reflect.TypeOf(string3))

	fmt.Println(x, ", type: ", reflect.TypeOf(x))

	fmt.Println(y, ", type: ", reflect.TypeOf(y))
}

var DataFloat float32 = 20.5 // public, diawali huruf kapital
var dataFloat float32 = 20.5 // private, diawali huruf kecil

var x int = int(dataFloat)

var y string = strconv.Itoa(x)

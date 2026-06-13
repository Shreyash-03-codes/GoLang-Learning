package main

import "fmt"

func main(){
	fmt.Println("Data Types....")

	var isActive bool = true

	var char rune = 'C'

	var str string = "Shreyash"

	var percentage float32 = 65.6455
	var pi float64 = 3.142212342345433422

	var byte int8 = 127
	var short int16 = 31234
	var integer int32 = 234242342
	var long int64 = 324234151542123

	var comp1 complex64 = 12+1i
	var comp2 complex128 = 454+1212i

	fmt.Println(isActive)

    fmt.Println(char)
	 
	fmt.Println(str)
	
	fmt.Println(percentage)
	fmt.Println(pi)
	
	fmt.Println(byte)
	fmt.Println(short)
	fmt.Println(integer)
	fmt.Println(long)
	
	fmt.Println(comp1)
	fmt.Println(comp2)
	
}

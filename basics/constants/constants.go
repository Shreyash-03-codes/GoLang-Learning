package main

import "fmt"

const PI = 3.142212342341223

const(
	name = "Shreyash"
	age = 21
	country = "India"
)

const(
	_ = iota
	sunday
	monday
	tuesday
	wednesday
	thursday
	friday
	saturady
)



func main(){

	// var num int = 20
	// const n=num                  //   this is not allowed in golang

	fmt.Println(PI)

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(country)

	// name="sbg"             X cannt modify
	fmt.Println(name)      


	fmt.Println(sunday)
	fmt.Println(monday)
	fmt.Println(tuesday)
	fmt.Println(wednesday)
	fmt.Println(thursday)
	fmt.Println(friday)
	fmt.Println(saturady)


	

}

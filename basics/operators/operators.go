package main

import "fmt"

func main(){

	num1 := 10
	num2 :=5

	// Arithmetic operators

	add := num1+num2
	sub :=num1-num2
	mult := num1*num2
	div :=num1/num2
	mod := num1%num2


	fmt.Println("Addition + :",add)
	fmt.Println("Substraction - :",sub)
	fmt.Println("Multiplication * :",mult)
	fmt.Println("Division / :",div)
	fmt.Println("Modulo % :",mod)


	
	// Relational Operartors

	var gt bool = num1 > num2
	var gtEquals bool = num1 >= num2
	var lt bool= num1<num2
	var ltEquals bool= num1<=num2
	var equals bool= num1==num2
	var notEquals bool= num1 != num2

	fmt.Println("Greater Than > :",gt)
	fmt.Println("Greater Than Equals >= :",gtEquals)
	fmt.Println("Less Than < :",lt)
	fmt.Println("Less Than Equals <= :",ltEquals)
	fmt.Println("Equals To == :",equals)
	fmt.Println("No Equals != :",notEquals)


	// Logical Operators

	b1 := true
	b2 := false
	
	fmt.Println("Logical AND && :",b1 && b2)
	fmt.Println("Logical OR || :",b1 || b2)
	fmt.Println("Logical NOT ! :",!b2)

	// Bitswise Operators

	fmt.Println("Bitwise AND & :",num1 & num2)
	fmt.Println("Bitwise OR | :",num1 | num2)
	fmt.Println("Bitwise XOR ^ :",num1 ^ num2)
	fmt.Println("Right Shift >> :",num2 >> num1)
	fmt.Println("Left Shift << :",num1 << num2)
	

	// Assignment Operators

	var num int =1000

	fmt.Println(" = 1000 :",num)

	num+=100
	fmt.Println("+= 100 :",num)

	num-=120
	fmt.Println("-= 120 :",num)


	num*=78
	fmt.Println("*= 78 :",num)

	num/=25
	fmt.Println("/= 25 :",num)

	num%=99
	fmt.Println("%= 99 :",num)
}

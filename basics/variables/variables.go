package main

import "fmt"

var JUSTNO int

// NO :=20           // wrong

func main(){

	JUSTNO=10

	var num int
	var _s string 
	var ch1 rune
	var per float32
	var b bool
	// var 123num int                X wrong decalration
	// var if int                  X wrong decalration

	dd := "direct declration"

	fmt.Println(dd)

	num=50
	_s="string"
	ch1='a'
	per=64.65
	b=false


	fmt.Println(num)
	fmt.Println(per)
	fmt.Println(_s)
	fmt.Println(ch1)
	fmt.Println(b)
	fmt.Println(JUSTNO)
	// fmt.Println(NO)

}
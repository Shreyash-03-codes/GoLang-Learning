package main

import "fmt"


func fun(nums... int) {
	fmt.Println(nums)
}

func main(){

	fun(12,23,34)
	fun(12,23,3412,23)
	fun(12)
	fun()
	
}

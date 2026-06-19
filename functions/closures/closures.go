package main

import "fmt"

func increase() func() int{
	i:=0

	return func()(int){
		i++
		return i
	}
}


func main(){

	inc:=increase()

	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	
	
}


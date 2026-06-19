package main

import "fmt"

func main(){

	func(){
		fmt.Println("Greetings...")
	}()

	func(name string){

		fmt.Println("Hello ",name)
	}("Shreyash")

	add := func(num1,num2 int) (int){
		return num1+num2;
	}

	fmt.Println(add(10,20))
}

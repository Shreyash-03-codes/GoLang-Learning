package main

import "fmt"

func main(){

	var start int
	var end int

	fmt.Println("Enter The Start...")
	fmt.Scan(&start)

	fmt.Println("Enter The End..")
	fmt.Scan(&end)

	for fmt.Println("For Loop Start...") ; start <=end ; fmt.Println(start-1,"th Iteration...") {

		if(start%2==0){
			fmt.Println("The Number if Even....")
		} else {
			fmt.Println("The Number is Od....")
		}
		start++
	}


}

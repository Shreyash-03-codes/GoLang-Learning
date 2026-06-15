package main

import "fmt"

func main(){

	var num int

	fmt.Println("Enter The Number of Day from 0 to 6 .....")
	fmt.Scan(&num)

	switch(num){
		case 0,7 :
			fmt.Println("Sunday...")

		case 1 :
			fmt.Println("Monday...")

		case 2 :
			fmt.Println("Tuesday...")

		case 3 :
			fmt.Println("Wednesday...")

		case 4 :
			fmt.Println("Thursday...")

		case 5 :
			fmt.Println("Friday...")

		case 6 :
			fmt.Println("Saturday...")
		
		default :
			fmt.Println("No Day...")


		
	}


	
}

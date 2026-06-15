package main

import "fmt"

func main(){

	var num int 

	fmt.Println("Enter The Number...")
	fmt.Scan(&num)

	if fmt.Println("Hello"); (num % 2 == 0 && num > 0) {
		fmt.Println("The Number is Even...")
	} else if(num%2==1 && num>0) {
		fmt.Println("The Number is Odd....")
	} else if(num<0){
		fmt.Println("The Number is Negtive")
	} else{
		fmt.Println("The Number is zero")
	}


}

package main


import "fmt"

func getNumbers() (int,int) {

	var num1 int
	var num2 int

	fmt.Println("Enter The 2 Numbers...")
	fmt.Scan(&num1,&num2)

	return num1,num2
}

func add(num1 int,num2 int) (int) {
	return num1+num2;
}

func main(){

	// var num1 int
	// var num2 int
	num1,num2  := getNumbers()

	sum:= add(num1,num2)

	fmt.Println(sum)



}

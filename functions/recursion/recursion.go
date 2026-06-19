package main


import "fmt"


func fact(num int) (int) {
	if(num<=1){
		return 1;
	}

	return num*fact(num-1)
}

func sum(num int) (int) {
	if(num==0){
		return 0;
	}

	return num+sum(num-1)
}

func main(){

	fmt.Println(fact(5))
	fmt.Println(sum(10))
}

package main

import "fmt"

func main(){

	var num int
	var str string
	var ch rune
	var b bool
	var per float32

	fmt.Print("Hello ")
	fmt.Print("World...\n")

	// fmt.Scan(&num, &str, &ch, &b, &per)

	// fmt.Println(num," ",str," ",ch," ",b," ",per)

	fmt.Scanf("%d %s %c %t %f",&num,&str,&ch,&b,&per)

	fmt.Printf("%d %s %c %t %f",num,str,ch,b,per)
}

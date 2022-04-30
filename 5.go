package main


import "fmt"

var dec int = 11 //variable at global level


func main(){
	var dec int = 10 
	fmt.Println(dec)//functional level scope
	if anyVar := 9 ; anyVar>10{
		
		fmt.Println("x is greater than 10")
	}else{
		
		fmt.Println("x is smaller than 10")
	}
	
}
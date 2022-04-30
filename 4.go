package main


import "fmt"
import "math/rand"


func main(){
	j:=0
	//conventional for loop
	for i:=0; i<5;i++ {
		fmt.Println(i)
	}

	//for loop used as while
	for j < 5{ 
		fmt.Println(j)
		j++
	}

	for{
		fmt.Println(rand.Int())
	}
}
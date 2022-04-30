package main


import "fmt"
import "time"

func main(){
	var added int //function assigned to a variable added
	added =add(1,2)
	fmt.Println(added) // variable added used
	var swapX,swapY int = 5,7
	swap(swapX,swapY)
	fmt.Printf("SwapX=%d | SwapY = %d\n",swapX,swapY)
	swapX,swapY = swap(swapX,swapY)
	fmt.Printf("SwapX=%d | SwapY = %d\n",swapX,swapY)
	t := time.Now()
	fmt.Printf(t.String())

}

func add(i , j int) int { //opening of the braces should be on same line as of func declaration(if done viceversa ; are automated)
		return i + j
}

func swap(i, j int)(int,int){
	return j, i
}
package main
import  "fmt"

/* Pointer = direct access to memory

Changing *pointer = change the original variable without using the variable name itself

*/

func main(){
	x := 10
	p :=  &x //Store x vaiable path address
	fmt.Println(p)
	fmt.Println(*p) //Go to that address and get the value and print

	//Change its value without calling/reffering the variable
	*p = 20

	fmt.Println(x) 
}
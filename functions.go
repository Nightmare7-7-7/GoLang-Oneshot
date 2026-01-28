package main
import  "fmt"


/*func functionName(parameters) returnType {
        //code 
 }*/

func greet(say string) string {
	return say
}

//Multiple Return Values
func divide(a, b int) (int, bool) {
    if b == 0 {
        return 0, false
    }
    return a / b, true
}


func main(){
	fmt.Println(greet("Hello"))
	result, ok := divide(10, 2)
    if ok {
        fmt.Println("Result:", result,ok)
    } else {
        fmt.Println("Cannot divide")
    }

}
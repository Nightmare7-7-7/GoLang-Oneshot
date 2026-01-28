package main
import "fmt"
import "errors"

func divide(a int, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main(){
	//Basic error creation
	err := errors.New("Somethig went wrong try later")
	fmt.Println(err)

	// real case secnario

	result, err := divide(10, 0)
	
	if err != nil {
    fmt.Println("Error:", err)
       return
     }
    fmt.Println("Result:", result)
	
	

	
}
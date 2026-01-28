package main
import "fmt"

func main(){
	age := 19

    if age > 19 {
        fmt.Println("You are a full grown man")
    } else if age < 20 && age > 13 { //else if and else must be directly after } — no newline in between.
        fmt.Println("You are a teen")
    } else {
        fmt.Println("You are a child")
    }

	//SHORT VARIABLE DECLARATION IN IF

	if score := 32; score < 30 {
		fmt.Println("You failed in exams")
	}else {
		fmt.Println("You passed in exams")
	}

}
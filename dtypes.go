package main
import "fmt"

func main(){
	//vars
	var fullname string = "khuman poudel"
	var age int = 19
	var weight float64 = 180.36
	var isStaff bool = true

	fmt.Println("Information!!!!")
	fmt.Println("fullname:",fullname)
	fmt.Println("age:",age)
	fmt.Println("weight:",weight)
	fmt.Println("isStaff:",isStaff)

	//Types conversion

	fmt.Println("")
	fmt.Println("Types conversion")

	var x int = 10
	var y float64 = float64(x)  // convert int → float64
	fmt.Println(y)

	rollno := 1
	rollStr := fmt.Sprintf("%d", rollno)  // int → string
	fmt.Println("roll as string:", rollStr)
	fmt.Printf("Type of y: %T\n", y)
	fmt.Printf("Type of rollStr: %T\n", rollStr)

}
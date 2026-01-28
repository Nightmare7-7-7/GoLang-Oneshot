package main
import  "fmt"

func main(){
	//Basic slice creation
	var nums[]int

	//Method 2: make()
	//nums := make([]int, 3)
	//fmt.Println(nums) // [0 0 0]
	
	//nums := make([]int, 3, 5)
	//With length + capacity

	//append
	nums = append(nums,8)
	fmt.Println(nums)


}
package main
import "fmt"

func main(){
	var nums[3]int
	//Size must be known at compile time
	//Size cannot change
	nums[1] = 26
	fmt.Println(nums)

	nums2 := [4]int{2,4,6,8} //Or [...] = Let Go count size
	fmt.Println(nums2)

	//array length
	fmt.Println(len(nums2))

	//Multi-Dimensional Arrays

	matrix := [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
   }
   fmt.Println(matrix[1][2])




}
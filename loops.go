package main
import "fmt"

func main(){

	//Basic for loop
	for i :=0; i < 10; i++ {
		fmt.Println("I have counted",i,"times")
	}

	//Nested one

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 2; j++ {
			fmt.Printf("i=%d, j=%d\n", i, j)
    } 
   }

   //RANGE LOOP

   fruits := [4]string{"apple","banana","orange","pineapple"}  
   
   for index,value := range fruits{
	fmt.Println(index,value)
   }


	
}
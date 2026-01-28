package main
import  "fmt"

func main(){
	//Declare & Initialize Maps

	user := map[string]int{
		"Number": 987263666,
		"Age": 19,
	}

	delete(user, "Age") //To delete any key value pairs
	user["Roll"] = 1 //Add or Append new key value pairs

	//Print with range
	for key,val := range user{
		fmt.Println(key,val)
	}
	


}
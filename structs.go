package main
import  "fmt"


//A struct is a custom data type that groups related values together.(A group of different datatypes)

type User struct {
	_id int
	Name string
	Email string
}

type Info struct {
	User User
    City    string
    Country string
}





func main(){
	//Create Struct Instance

	var u User
	u._id = 1
	u.Name = "khuman Poudel"
	u.Email = "khumanpoudel@gmail.com"

	//u2 := User{"Ram", 25, "ram@gmail.com"} //shorthand

	fmt.Println(u)

	//Struct + Slice
	users := []User{
		{2,"karan","karan@gmail,com",},
		u,

	}
	fmt.Println(users)

	//Nested Struct 

	fullinfo := Info{
		User: User{
			_id:4,
			Name:"Bishnu Poudel",
			Email:"xdopa@gmail.com",
		},
		City:"walling",
		Country:"Nepal",
	}

	fmt.Println(fullinfo)



}
package main
import  "fmt"

type Human struct {
	name string
}

type Dog struct {
	name string
}

func (h Human) Greet(){ //struct + method is like attaching function inside struct 
	fmt.Println("You are a human and your name is",h.name)
}

func (d Dog) Greet(){
	fmt.Println("You are a dog and your name is",d.name)
}

type robust interface {
	Greet() 
}

func print(r robust){
	r.Greet()
}

func main(){
	h := Human{"Khuman"}
	d := Dog{"Jack"}

	//with interface
	print(h)
	print(d)

}
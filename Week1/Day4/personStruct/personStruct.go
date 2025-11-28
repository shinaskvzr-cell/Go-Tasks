package personStruct
import "fmt"

type Person struct{
	Name string
	Age int
	City string
}

func (p Person) PrintDetails(){
	fmt.Println("Name:",p.Name,"\nAge:",p.Age,"\nCity:",p.City)
}

func PersonStruct(){
	fmt.Println("\nCreate a 'person'struct with a method to print")

	p:=Person{Name:"Shinas",Age:21,City:"Malappuram"}
	p.PrintDetails()
}
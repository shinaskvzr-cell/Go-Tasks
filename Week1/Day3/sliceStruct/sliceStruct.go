package sliceStruct
import "fmt"

type Student struct{
	Name string
	Mark int
}

func SliceStruct(){
	s := []Student{
		{Name:"Shinas",Mark:90},
		{Name:"Anshad",Mark:99},
		{Name:"Yadhu",Mark:97},
		{Name:"Jasil",Mark:95},
	}
	fmt.Println("\nUse a slice to store students")
	fmt.Println("Display list using a loop")
	for _,v := range s{
		fmt.Println(v.Name,v.Mark)
	}
}
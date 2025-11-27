package studentStruct
import "fmt"

type Student struct {
	Name string
	Age int
}

func (s *Student) UpdateStudent(name string,age int){
	s.Name = name
	s.Age = age
}

func (s *Student) DeleteStudent () {
	s.Name = ""
	s.Age = 0
}


func StudentStruct(){

	fmt.Println("\nDefine 'Student' struct with Name , age")
	fmt.Println("Implement 'UpdateStudent' and Delete Student with pointer logic.")
	s := &Student{Name:"Shinas",Age:21}
	fmt.Println(s.Name)
	fmt.Println(s.Age)
	s.UpdateStudent("SHINAS MUHAMMED K",26)
	fmt.Println(s.Name)
	fmt.Println(s.Age)
}
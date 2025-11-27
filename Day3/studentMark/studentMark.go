package studentMark
import "fmt"

func StudentMark(){
	students := make(map[string]int)
	students["Shinas"]=90
	students["Anshad"]=99
	students["Yadhu"]=97
	students["Jasil"]=95
	
	fmt.Println("\nMap for student names and grades")
	for name,grade := range students{
		fmt.Println(name,grade)
	}
}
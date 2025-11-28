package removeDuplicate
import "fmt"

func RemoveDuplicate(){
	arr := []int{1,1,2,3,4,5,5,4,2,3,7}
	seen := make(map[int]bool)
	unique := []int{}

	fmt.Println("\nSlice functions to remove duplicates")
	for _,v := range arr{
		if !seen[v]{
			unique = append(unique,v)
			seen[v]=true
		}
	}
	fmt.Println("Unique values are :",unique)

}
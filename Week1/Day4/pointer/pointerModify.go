package pointer
import "fmt"

func Modify (x *int){
	*x =100
}

func PointerModify(){
	fmt.Println("\nFunction that modifies int using pointer")
	x := 10
	Modify(&x)
	fmt.Println(x)
}
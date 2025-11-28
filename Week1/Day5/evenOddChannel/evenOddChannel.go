package evenOddChannel
import "fmt"

func EvenOddChannel(){
	fmt.Println("\nGoroutines to print odd and even cocrrently")
	even := make(chan int)
	odd := make(chan int)

	go func(){
		for i:=0;i<10;i+=2{
			even <-i
		}
		close(even)
	}()

	go func(){
		for i:=1;i<10;i+=2{
			odd <-i
		}
		close(odd)
	}()

	for v := range even{
		fmt.Println("Even",v)
	}

	for v := range odd{
		fmt.Println("Odd",v)
	}
}
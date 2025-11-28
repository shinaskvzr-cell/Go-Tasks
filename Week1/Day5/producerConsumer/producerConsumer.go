package producerConsumer
import "fmt"

func ProducerConsumer(){
	fmt.Println("\nChannelbased producerconsumer")
	jobs := make(chan int)
	go func(){
		for i:=0;i<5;i++{
			fmt.Println("Producing",i)
			jobs <- i
		}
		close(jobs)
	}()

	for v := range jobs {
		fmt.Println("Consuming",v)
	}
	fmt.Println("All jobs processed")
}
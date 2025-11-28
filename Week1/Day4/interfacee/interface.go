package interfacee
import "fmt"

type Shape interface{
	Area() float64
	Perimeter() float64
}

type Circle struct{
	radius float64
}

type Rectangle struct{
	width float64
	height float64
}

func (c Circle) Area() float64{
	return 3.14*c.radius*c.radius
}

func (c Circle) Perimeter()float64{
	return 2*3.14*c.radius
}

func (r Rectangle) Area() float64{
	return r.width*r.height
}

func (r Rectangle) Perimeter() float64{
	return 2*(r.width+r.height)
}

func Interface(){
	fmt.Println("\nInterface with 'Shape'(Circle,Rectangle)")

	c := Circle{radius:10}
	r := Rectangle{width:10,height:10}
	fmt.Println("Area of the Circle is :",c.Area())
	fmt.Println("Perimeter of the circle is :",c.Perimeter())
	fmt.Println("Area of the Rectangle is :",r.Area())
	fmt.Println("Perimeter of the Rectangle is :",r.Perimeter())
}
package main
import "fmt"
import "math"

type shape interface {
    perimeter() float64
}
type rectangle struct {
    width, height float64
}
type circle struct {
    radius float64
}
func (r rectangle) perimeter() float64 {
    return 2*r.width + 2*r.height
}
func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}
func main() {
    fmt.Println("Enter the width of rectangle:")
    var x float64
    fmt.Scan(&x)
    fmt.Println("Enter the height of rectangle:")
    var y float64
    fmt.Scan(&y)
    fmt.Println("Enter the radius of circle:")
    var z float64
    fmt.Scan(&z)

    r := rectangle{width: x, height: y}
    c := circle{radius: z}

   	fmt.Println(r.perimeter())
   	fmt.Println(c.perimeter())
   	}
package main
import "fmt"
func main() { 
	example()
	fmt.Println("Finish")
	}
func example(){
	defer func() { 
		r := recover() 
		fmt.Println(r) }() 
		panic("Failure Panic")
}

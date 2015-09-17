package main
import "fmt"

func fibo(x int) int {
	if x == 0 {return x} else if x == 1{return 1}else {return fibo(x-1) + fibo (x-2)}
	}
func main() { 
	fmt.Println("Enter the number:")
	var x int
    fmt.Scan(&x)
	fmt.Println(fibo(x))
	}

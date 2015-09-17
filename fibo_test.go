package main
import "testing"
import "fmt"

func TestFibo1(t *testing.T) {
	if(fibo(1) != 1){
	 t.Error("Failed")
	 fmt.Println("The fibonacci function is not correct")
	}else{
	t.Log("Passed")
	fmt.Println("Fibonacci test one is successfully passed")
	}
 }

 func TestFibo2(t *testing.T) {
	if(fibo(6) != 8){
	 t.Error("Failed")
	 fmt.Println("The fibonacci function is not correct")
	}else{
	t.Log("Passed")
	fmt.Println("Fibonacci test two is successfully passed")
	}
 }


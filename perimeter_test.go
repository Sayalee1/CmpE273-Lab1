package main
import "fmt"
import "testing"

func Test_Rectperimeter1(t* testing.T){
	r := rectangle{width: 2,height: 3}
	if(r.perimeter() != 10){
	 t.Error("Failed")
	 fmt.Println("Rectangle perimeter test one is failed")
	}else{
	t.Log("Passed")
	fmt.Println("Rectangle perimeter test one is passed")
	fmt.Println(r.perimeter())
	}
}


func Test_Cirperimeter1(t* testing.T){
	c := circle{radius: 1}
	if(c.perimeter() < 6 || c.perimeter() > 7){
	 t.Error("Failed")
	 fmt.Println("Circle perimeter function test one is failed")
	}else{
	t.Log("Passed")
	fmt.Println("Circle perimeter test one is passed")
	fmt.Println(c.perimeter())
	}
}

func Test_Rectperimeter2(t* testing.T){
	r := rectangle{width: 5,height: 5}
	if(r.perimeter() != 20){
	 t.Error("Failed")
	 fmt.Println("Rectangle perimeter test two is failed")
	}else{
	t.Log("Passed")
	fmt.Println("Rectangle perimeter test two is passed")
	fmt.Println(r.perimeter())
	}
}


func Test_Cirperimeter2(t* testing.T){
	c := circle{radius: 10}
	if(c.perimeter() < 62 || c.perimeter() > 63){
	 t.Error("Failed")
	 fmt.Println("Circle perimeter function test two is failed")
	}else{
	t.Log("Passed")
	fmt.Println("Circle perimeter test two is passed")
	fmt.Println(c.perimeter())
	}
}

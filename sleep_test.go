package main

import "testing"
import "fmt"
import "time"


func Test_Sleep1(t* testing.T){	
	
	t1 := time.Now().Second()
	fmt.Println(t1)

	Sleep(10)
    
	t2 :=  time.Now().Second()
	fmt.Println(t2)

	if(t2-t1 < 10){
	 t.Error("Failed")
	 fmt.Println("The Sleep function has failed")
	}else{
	t.Log("Passed")
	fmt.Println("The sleep function passed the test one")
	}		
}

func Test_Sleep2(t* testing.T){
	t1 := time.Now().Second()
	fmt.Println(t1)

	Sleep(15)
    
	t2 :=  time.Now().Second()
	fmt.Println(t2)

	if(t2-t1  < 15){
	 t.Error("Failed")
	 fmt.Println("The Sleep function has failed")
	}else{
	t.Log("Passed")
	fmt.Println("The sleep function passed the test two")
	}
}


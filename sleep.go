package main
import "time"
import "fmt"
func main() {
	fmt.Println("Enter sleep time:")
	var n int
    fmt.Scan(&n) 
	fmt.Println("Sleep starting")
	Sleep(n)
	fmt.Println("Sleep Finish")
}
func Sleep(n int) {
	<-time.After(time.Second * time.Duration(n))
}
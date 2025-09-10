package main
import "fmt"
func main() {
	var i interface{} = "comma ok is an ok boolean type assertion -> variableName, ok := some value, okk error"
	number, ok := i.(int)
	j := i.(int) /*panic, type assertion without comma-ok if incorrect*/
	if ok {
		fmt.Println("number : ", number)
	} else {
		fmt.Println("not an int")
	}
	fmt.Println(j)
}

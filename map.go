package main
import "fmt"
func main() {
	/*map[KeyType]ValueType*/
	var chicken map[string]int
	/*panic: assignment to entry in nil map*/
	chicken = map[string]int{}
	chicken["januari"] = 50
	myMap := map[string]int{"banana":5, "apple":10}
	symbol := "BBCA"
	stocks := make(map[string]float64)
	stocks[symbol] = 125
	for k, vv := range chicken {
		fmt.Println(k, vv)
	}
	_, ok := myMap["fruit"]
	if ok {
		
	} else {
		fmt.Println("key not found")
	}
}

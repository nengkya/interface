package main
import "math" /*undefined math*/
import "fmt"

/*geometry interface has luas() function*/
type geometry interface{
	/*
	func functionName(parameter1 type1, parameter2 type2) returnType {
        function body
        return value
    }
	*/
	luas() float64
}

type persegi struct {
	sisi float64
}

/*persegi struct has luas method*/
func (persegiA persegi) luas() float64 {
	return math.Pow(persegiA.sisi, 2)
}

func main() {
	var bangunDatar geometry
	/*missing method luas*/
	bangunDatar = persegi{10}
	fmt.Println(bangunDatar.luas())
}

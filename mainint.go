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
	luas() int
}

type persegi struct {
	sisi float64
	soso int
}

/*
persegi struct has luas method
have luas() int
want luas() float64
cannot use persegiA.sisi (variable of type int) as float64 value in argument to math.Pow
cannot use persegi{…} (value of struct type persegi) as geometry value in assignment: persegi does not implement geometry (wrong type for method luas)
int refer to interface, not struct
*/
func (persegiA persegi) luas() int {
	/*func Pow(a, b float64) float64*/
	return int(math.Pow(persegiA.sisi, 2))
}

func main() {
	var bangunDatar geometry
	/*missing method luas*/
	bangunDatar = persegi{10, 1}
	fmt.Println(bangunDatar.luas())
}

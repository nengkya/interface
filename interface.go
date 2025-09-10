package main
import "math"
import "fmt"

/*geometry interface has luas() function*/
type geometry interface{
	luas() int
}

type persegi struct {
	sisi float64
	soso int
}

/*
persegi struct has luas method
int refer to interface, not struct
*/
func (persegiA persegi) luas() int {
	/*func Pow(a, b float64) float64*/
	return int(math.Pow(persegiA.sisi, 2))
}

func main() {
	var bangunDatar geometry
	bangunDatar = persegi{10, 1}
	fmt.Println(bangunDatar.luas())
}

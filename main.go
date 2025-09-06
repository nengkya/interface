/*
Reasons for using a method receiver instead of a function argument:

Code Organization and Readability:
Method receivers allow functions to be directly associated with a specific type, creating a more organized and object-oriented-like structure.
This improves code readability by clearly indicating which operations belong to which data structures.

Encapsulation:
Methods with receivers encapsulate logic and behavior directly within the type they operate on.
This promotes better design by keeping related functions and data together,
enhancing maintainability and reducing the potential for side effects. 

Cleaner Syntax:
Calling a method on a type instance using the dot notation (e.g., myObject.DoSomething()) is generally cleaner and more intuitive
than passing the object as a separate argument to a standalone function (e.g., DoSomething(myObject)).

Leveraging Type State:
Receivers provide direct access to the fields and state of the associated type,
eliminating the need to explicitly pass those values as individual arguments to the function.

Interface Implementation:
Receivers are fundamental for implementing interfaces in Go.
By defining methods with specific receivers, types can fulfill interface contracts, enabling polymorphism and flexible code design.

Disambiguation of Generic Function Names:
While Go does not support function overloading in the traditional sense,
using receivers allows for methods with the same name to exist across different types,
as the receiver type provides the necessary disambiguation.

Situations where function arguments might be preferred:
Standalone Utility Functions:
For functions that perform general operations not inherently tied to a specific type, using regular function arguments is appropriate.
Creating New Instances:
When a function's primary purpose is to construct and return a new instance of a type,
passing initial values as arguments is the standard approach (e.g., NewObject(arg1, arg2)).
*/

package main
import "math" /*undefined math*/
import "fmt"

type hitung interface{
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
	var bangunDatar hitung
	/*missing method luas*/
	bangunDatar = persegi{10}
	fmt.Println()
}

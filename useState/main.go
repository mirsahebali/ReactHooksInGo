// useState hook of react in go
package main

import "fmt"

func useState(initVal int) (func() int, func(int)) {
	// assigning to a temp variable which holds the state
	_val := initVal
	// getter function which returns the state of _val when we change initVal
	state := func() int { return _val }
	// dispatching setState which assigns finalval to _val which handles the return type of state()
	setState := func(finalVal int) {
		_val = finalVal
	}
	// Return both
	return state, setState
}

func main() {
	// using two varialbes
	num, setNum := useState(8)
	setNum(9)
	// using num as state()
	fmt.Println(num())
}

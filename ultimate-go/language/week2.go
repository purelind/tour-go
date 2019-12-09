//
// week2 module2 gendisplacefn.go
//

package main

import (
	"fmt"
	"math"
)

func main() {
	var acceleration, velocity, initial, elapseTime float64
	fmt.Printf("[ acceleration ] -> ")
	fmt.Scanf("%f", &acceleration)
	fmt.Printf("[ velocity ] -> ")
	fmt.Scanf("%f", &velocity)
	fmt.Printf("[ initial displacement ] -> ")
	fmt.Scanf("%f", &initial)
	fmt.Printf("[ time ] -> ")
	fmt.Scanf("%f", &elapseTime)

	fn := GenDisplaceFn(acceleration, velocity, initial)
	fmt.Printf("\nDisplacement after %.2f seconds -> %.2f", elapseTime, fn(elapseTime))
}

func GenDisplaceFn(acceleration, velocity, initial float64)  func (float64) float64 {
	return func (t float64) float64 {
		return  0.5 * acceleration  * math.Pow(t, 2) + velocity * t + initial
	}
}



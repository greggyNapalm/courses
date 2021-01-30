package main

import (
	"fmt"
	"math"
)

func ReadInputFloat(msg string) float64 {
	var rv float64
	fmt.Println("Please enter " + msg + ":")
	fmt.Scanf("%f", &rv)
	return rv
}

func GenDisplaceFn(a float64, v0 float64, s0 float64) func(t float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
}

func main() {
	a := ReadInputFloat("acceleration `a`")
	v0 := ReadInputFloat("initial velocity `vo`")
	s0 := ReadInputFloat("initial displacement `so`")
	displacementFn := GenDisplaceFn(a, v0, s0)
	t := ReadInputFloat("time `t`")
	fmt.Printf("displacement = %.2f for a=%.2f, v0=%.2f, s0=%.2f", displacementFn(t), a, v0, s0)
}

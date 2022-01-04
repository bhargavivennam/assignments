package main

import (
	"funcExample/calculator"
	"funcExample/phone"
)

func main() {
	calculator.Add(8, 2)
	calculator.Sub(8, 2)
	calculator.Mul(8, 2)
	calculator.Divide(8, 2)
	calculator.Mod(8, 2)
	phone.Ring()
}

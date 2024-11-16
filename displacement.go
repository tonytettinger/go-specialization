package main
import (
	"strconv"
	"fmt"
	"math"
	"log"
)

func SaveToFloatVar(message string, val *float64) {
	var temp string
	fmt.Println(message)
	fmt.Scanln(&temp)
	floatNum, err := strconv.ParseFloat(temp, 64)
	if err != nil {
		log.Fatal("Error! Please enter a valid float number!")
	}
	*val = floatNum
}

func GenDisplaceFn(acceleration float64, initialVelocity float64, initialDisplacement float64) func (t float64) float64  {
	return func(t float64) float64 {
		return 0.5 * acceleration * math.Pow(t,2) + initialVelocity * t + initialDisplacement
	}
}

func main() {
	var a,iV, iD, t float64
	var calculateDisplacement func(t float64) float64
	SaveToFloatVar("Enter acceleration", &a)
	SaveToFloatVar("Enter initial velocity", &iV)
	SaveToFloatVar("Enter initial displacement", &iD)

	calculateDisplacement = GenDisplaceFn(a,iV,iD)

	SaveToFloatVar("Enter time", &t)

	displacement := calculateDisplacement(t)
	fmt.Println("The displacement value is: ", displacement)


}
	
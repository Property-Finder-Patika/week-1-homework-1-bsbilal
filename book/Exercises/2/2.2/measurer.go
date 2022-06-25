package main

import (
	"fmt"
)

type Kilogram float64
type Pound float64
type Feet float64
type Meter float64
type Fahrenheit float64
type Celsius float64

func CalculateResults(t float64) {
	// temp prints
	fmt.Println("Temperature Conversion:")
	f := Fahrenheit(t)
	c := Celsius(t)
	fmt.Printf("%s equals %s, %s equals %s\n\n", f, FToC(f), c, CToF(c))
	// length prints
	fmt.Println("Length Conversion:")
	lF := Feet(t)
	lM := Meter(t)
	fmt.Printf("%s equals %s, %s equals %s\n\n",
		lF, FToM(lF), lM, MToF(lM))
	// weight prints

	fmt.Println("Weigth Conversion:")
	wK := Kilogram(t)
	wP := Pound(t)
	fmt.Printf("%s equals %s, %s equals %s\n\n",
		wK, KgToLb(wK), wP, LbToKg(wP))

}

func main() {
	fmt.Print("Please enter measure=> ")
	var enteredValue float64
	_, err := fmt.Scanf("%f", &enteredValue)
	if err != nil {
		fmt.Println(err)
	}
	CalculateResults(enteredValue)

}
func (w Kilogram) String() string {
	// Kilogram formatter
	return fmt.Sprintf("%g(KG)", w)
}
func (w Pound) String() string {
	// Kilogram formatter
	return fmt.Sprintf("%g(lbs)", w)
}
func (f Feet) String() string {
	// feet formatter
	return fmt.Sprintf("%g(ft)", f)
}
func (m Meter) String() string {
	// meter formatter
	return fmt.Sprintf("%g(m)", m)
}
func (c Celsius) String() string {
	// celsius formatter
	return fmt.Sprintf("%g(C)", c)
}
func (f Fahrenheit) String() string {
	// fahrenheit formatter
	return fmt.Sprintf("%g(F)", f)
}

func FToM(f Feet) Meter {
	//it converts to Feet to Meter
	return Meter(f * 0.3048)
}
func MToF(m Meter) Feet {
	// it converts celsius to fahrenheit
	return Feet(m * 3.28083)
}
func CToF(c Celsius) Fahrenheit {
	// It converts Celsius to Fahrenheit
	return Fahrenheit(c*9/5 + 32)
}
func FToC(f Fahrenheit) Celsius {
	// It Converts Fahrenheit to Celsius
	return Celsius((f - 32) * 5 / 9)
}

func KgToLb(kg Kilogram) Pound {
	// It converts Kilograms to Pounds
	return Pound(kg * 2.2)
}
func LbToKg(p Pound) Kilogram {
	// It Converts Pounds to Kilograms
	return Kilogram(p / 0.45)
}

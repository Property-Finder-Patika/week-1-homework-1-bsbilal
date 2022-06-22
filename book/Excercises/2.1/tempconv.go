package tempconv

type Celsius float64
type Kelvin float64

const (
	BoilingC, FreezingC, AbsoluteZeroC Celsius = 100, 0, -273.15
	BoilingK, FreezingK, AbsoluteZeroK Kelvin  = 373.15, 273.15, 0
)

func CToK(c Celsius) Kelvin {
	// It converts Celsius to Kelvin
	return Kelvin(c + 273.15)
}
func KToC(k Kelvin) Celsius {
	// It Converts Kelvin to Celsius
	return Celsius(k - 273.15)
}

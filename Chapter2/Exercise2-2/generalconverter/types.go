package generalconverter

import "fmt"

type (
	Celsius    float64
	Fahrenheit float64
	Meters     float64
	Feets      float64
	Pounds     float64
	Kilogams   float64
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (m Meters) String() string     { return fmt.Sprintf("%g m", m) }
func (f Feets) String() string      { return fmt.Sprintf("%g ft", f) }
func (p Pounds) String() string     { return fmt.Sprintf("%g pounds", p) }
func (k Kilogams) String() string   { return fmt.Sprintf("%g kg", k) }

package generalconverter

func (c Celsius) ToF() Fahrenheit   { return Fahrenheit(c*9/5 + 32) }
func (f Fahrenheit) ToC() Celsius   { return Celsius((f - 32) * 5 / 9) }
func (m Meters) ToFeet() Feets      { return Feets(m * 3.281) }
func (f Feets) ToM() Meters         { return Meters(f / 3.281) }
func (k Kilogams) ToPounds() Pounds { return Pounds(k * 2.205) }
func (p Pounds) ToKg() Kilogams     { return Kilogams(p / 2.205) }

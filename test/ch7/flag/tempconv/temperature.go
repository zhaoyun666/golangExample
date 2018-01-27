package tempconv

import (
    "fmt"
    "flag"
)

type Celsius float64
type Fahrenheit float64

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Scanf(s, "%f%s", &value, &unit)
	fmt.Println("-------------------", unit)
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
    f := &celsiusFlag{value}
    flag.CommandLine.Var(f, name, usage)
    fmt.Println(name, usage, value.String())
    f.Set(value.String())
    return &f.Celsius
}

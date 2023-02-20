package main

import (
	"conv/conv"
	"flag"
	"fmt"
)

var C float64
var F float64
var K float64
var out string

func init() {
	flag.Float64Var(&C, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&F, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&K, "K", 0.0, "temperatur i kelvin")
	flag.StringVar(&out, "out", "", "temperaturskala for resultat")
}

func main() {
	flag.Parse()

	var res float64
	if K != 0 {
		if out == "C" {
			res = conv.KelvinToCelsius(K)
			fmt.Printf("%.2fK is %.2fF\n", K, res)
		} else if out == "F" {
			res = conv.KelvinToFahrenheit(K)
			fmt.Printf("%.2fK is %.2f°C\n", K, res)
		}
	}

	if C != 0 {
		if out == "F" {
			res = conv.CelcsiusToFahrenheit(C)
			fmt.Printf("%.2fC is %.2fF\n", C, res)
		} else if out == "K" {
			res = conv.CelcsiusToKelvin(C)
			fmt.Printf("%.2fC is %.2fK\n", C, res)
		}
	}

	if F != 0 {
		if out == "K" {
			res = conv.FahrenheitToKelvin(F)
			fmt.Printf("%.2fF is %.2fK\n", F, res)
		} else if out == "C" {
			res = conv.FarhenheitToCelsius(F)
			fmt.Printf("%.2fF is %.2f°C\n", F, res)
		}
	}

}

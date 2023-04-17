package main

import "fmt"

func ConvertTemperature(temp float64, fromScale string, toScale string) (float64, error) {
	var result float64

	if fromScale == "C" && toScale == "F" {
		result = temp*9/5 + 32
	} else if fromScale == "C" && toScale == "K" {
		result = temp + 273.15
	} else if fromScale == "F" && toScale == "C" {
		result = (temp - 32) * 5 / 9
	} else if fromScale == "F" && toScale == "K" {
		result = (temp-32)*5/9 + 273.15
	} else if fromScale == "K" && toScale == "C" {
		result = temp - 273.15
	} else if fromScale == "K" && toScale == "F" {
		result = (temp-273.15)*9/5 + 32
	} else {
		return 0, fmt.Errorf("Escala inválida.")
	}
	return 0, nil
}

func main() {
	r, err := ConvertTemperature(10, "C", "K")
	fmt.Println(r, err)

}

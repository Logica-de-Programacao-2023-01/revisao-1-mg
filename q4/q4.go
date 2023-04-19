package main

import "fmt"

func CalculateFinalPrice(basePrice float64, state string, taxCode int) (float64, error) {
	var result float64

	if state == "SP" && taxCode == 1 {
		result = basePrice + basePrice*0.05 + basePrice*0.10
	} else if state == "SP" && taxCode == 2 {
		result = basePrice + basePrice*0.10 + basePrice*0.10
	} else if state == "SP" && taxCode == 3 {
		result = basePrice + basePrice*0.15 + basePrice*0.10
	} else if state == "RJ" && taxCode == 1 {
		result = basePrice + basePrice*0.05 + basePrice*0.15
	} else if state == "RJ" && taxCode == 2 {
		result = basePrice + basePrice*0.10 + basePrice*0.15
	} else if state == "RJ" && taxCode == 3 {
		result = basePrice + basePrice*0.15 + basePrice*0.15
	} else if state == "MG" && taxCode == 1 {
		result = basePrice + basePrice*0.05 + basePrice*0.20
	} else if state == "MG" && taxCode == 2 {
		result = basePrice + basePrice*0.10 + basePrice*0.20
	} else if state == "MG" && taxCode == 3 {
		result = basePrice + basePrice*0.15 + basePrice*0.20
	} else if state == "ES" && taxCode == 1 {
		result = basePrice + basePrice*0.05 + basePrice*0.25
	} else if state == "ES" && taxCode == 2 {
		result = basePrice + basePrice*0.10 + basePrice*0.25
	} else if state == "ES" && taxCode == 3 {
		result = basePrice + basePrice*0.15 + basePrice*0.25
	} else if state == "OUTROS" && taxCode == 1 {
		result = basePrice + basePrice*0.05 + basePrice*0.30
	} else if state == "OUTROS" && taxCode == 2 {
		result = basePrice + basePrice*0.10 + basePrice*0.30
	} else if state == "OUTROS" && taxCode == 3 {
		result = basePrice + basePrice*0.15 + basePrice*0.30
	} else if basePrice <= 0 {
		return 0, fmt.Errorf("preço base inválido")
	} else if taxCode <= 0 && taxCode <= 4 {
		return 0, fmt.Errorf("imposto não encontrado")
	}

	return result, nil
}

func main() {
	r, err := CalculateFinalPrice(1000, "RJ", 3)
	fmt.Println(r, err)
}

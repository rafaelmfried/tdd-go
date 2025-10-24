package soma

import "fmt"

func Soma(numeros []int) int {
	result := 0
	for _, numero := range numeros {
		result += numero
	}
	return result
}

func SomTudo(numeros ...[]int) map[string]int {
	result := map[string]int{}
	for _, slice := range numeros {
		result[fmt.Sprintf("%v", slice)] = Soma(slice)
	}
	return result
}
package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ImportFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Não foi possível ler o arquivo")
		return ""
	}
	return string(content)
}

func TransformData(data, sep string) []float64 {
	results := []float64{}
	lines := strings.Split(data, sep)
	for _, v := range lines {
		value, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Println("Não foi possível converter o dado", v)
		}
		results = append(results, value)
	}
	return results
}

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	var (
		a         int
		b         int
		operation string
	)

	inputs := strings.Split(os.Args[1], " ")

	a, _ = strconv.Atoi(inputs[0])
	operation = inputs[1]
	b, _ = strconv.Atoi(inputs[2])

	var result int

	switch operation {
	case "+":
		result = a + b
		break
	case "-":
		result = a - b
		break
	case "*":
		result = a * b
		break
	case "/":
		result = a / b
		break
	case "**":
		res := math.Pow(float64(a), float64(b))
		result = int(res)
		break
	default:
		fmt.Printf("Invalid Operation: %s", operation)
	}

	fmt.Println(result)

}

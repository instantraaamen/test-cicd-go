package main

import (
	"fmt"
	"test-cicd-go/operations"
)

func main() {
	a, b := 20, 5

	fmt.Printf("加算: %d + %d = %d\n", a, b, operations.Add(a, b))
	fmt.Printf("減算: %d - %d = %d\n", a, b, operations.Sub(a, b))
	fmt.Printf("乗算: %d * %d = %d\n", a, b, operations.Mul(a, b))

	result, err := operations.Div(a, b)
	if err != nil {
		fmt.Println("エラー:", err)
	} else {
		fmt.Printf("除算: %d / %d = %d\n", a, b, result)
	}
}

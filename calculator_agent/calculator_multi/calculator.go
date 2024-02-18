package calculator_multi

import (
	"sync"
)

func Calculate(a, b int, op string, stack *[]int, wg *sync.WaitGroup) {
	defer wg.Done()
	switch op {
	case "+":
		*stack = append(*stack, a+b)
	case "-":
		*stack = append(*stack, a-b)
	case "*":
		*stack = append(*stack, a*b)
	case "/":
		if b == 0 {
			panic("division by zero")
		}
		*stack = append(*stack, a/b)
	}
}

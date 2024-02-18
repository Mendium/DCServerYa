package calculator_multi

func Calculate(a, b int, op string, stack *[]int) {
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

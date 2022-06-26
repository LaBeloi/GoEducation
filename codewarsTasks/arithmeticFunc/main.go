package main

func test(a, b int, oper string) int {
	switch oper {
	case "add":
		return a + b
	case "subtract":
		return a - b
	case "multiply":
		return a * b
	case "divide":
		return a / b
	default:
		return a + b
	}
}

func main() {
	println(test(5, 3, "multiply"))
}

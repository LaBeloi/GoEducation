package main

func cubes(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i * i * i
	}
	return sum
}

func main() {
	println(cubes(2))
}

package main

func main() {
	println("Hello World")
	println(If(true, add(1, 2), 3))
}

func add(a, b int) int {
	return a + b
}

func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	} else {
		return falseVal
	}
}

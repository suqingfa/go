package main

func main() {
	println("Hello World")
	println(If(true, add(1, 2), 3))
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}

func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	} else {
		return falseVal
	}
}

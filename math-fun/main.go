package main

// primera cosita que hago!
func DetermineDivisors(x int) []int {
	divisors := []int{}
	for i := 1; i < x; i++ {
		isDivisor := x%i == 0
		if isDivisor {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

// segunda cosa que se apoya en la primera!!
func isPerfect(x int) bool {
	if x <= 0 { // por definición tiene que ser un entero positivo (mayor a cero), ver https://en.wikipedia.org/wiki/Perfect_number
		return false
	}
	divisors := DetermineDivisors(x)
	return sumIntSlice(divisors) == x
}

// funcion auxiliar para sumar todos los elementos de slice de enteros (reducir el slice a un número via suma matemática)
func sumIntSlice(ints []int) int {
	ac := 0
	for _, i := range ints {
		ac += i
	}
	return ac
}

// tercera cosa que se apoya tmb en la primera
func AreFriends(x int, y int) bool {
	if x <= 0 || y <= 0 { // por definición tiene que ser un entero positivo (mayor a cero), ver https://en.wikipedia.org/wiki/Amicable_numbers
		return false
	}
	xDivisors := DetermineDivisors(x)
	yDivisors := DetermineDivisors(y)
	return sumIntSlice(xDivisors) == y && sumIntSlice(yDivisors) == x
}

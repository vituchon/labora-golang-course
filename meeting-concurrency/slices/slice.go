package slices

func SumNotUsingChannel(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func SumUsingChannel(s []int) int {
	c := make(chan int)
	go sumInChannel(s[:len(s)/2], c)
	go sumInChannel(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	return x + y
}

func sumInChannel(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func RotateRight(text string) string {
	size := len(text)
	rotatedText := ""
	for fromIndex := range text {
		toIndex := (size - 1 + fromIndex) % size
		char := string(text[toIndex])
		rotatedText = rotatedText + char
	}
	return rotatedText
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func RotateLeft(text string) string {
	size := len(text)
	rotatedText := ""
	for i := range text {
		char := string(text[(size+1+i)%size])
		rotatedText = rotatedText + char
	}
	return rotatedText
}

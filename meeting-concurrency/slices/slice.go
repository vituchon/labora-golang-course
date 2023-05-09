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

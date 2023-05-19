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

func RotateRightTimes(text string, times int) string {
	for i := 1; i <= times; i++ {
		text = RotateRight(text)
	}
	return text
}

func RotateRight(text string) string { // graphical idea: https://jamboard.google.com/d/1h5aAO-wiPI3TO5QIChnGQEaAPlT8dU79k3ZFckh2NW8/viewer?mtt=v7wp3cnbz3an&f=0
	size := len(text)
	rotatedText := ""
	for fromIndex := range text {
		toIndex := (size - 1 + fromIndex) % size
		char := string(text[toIndex])
		rotatedText = rotatedText + char
	}
	return rotatedText
}

func RotateRightVersion2(text string) string { // graphical idea: https://jamboard.google.com/d/1h5aAO-wiPI3TO5QIChnGQEaAPlT8dU79k3ZFckh2NW8/viewer?mtt=v7wp3cnbz3an&f=0
	size := len(text)
	var rotatedChars []rune = make([]rune, size)
	for posicionNueva := range text {
		posicionOriginal := (size - 1 + posicionNueva) % size
		char := rune(text[posicionOriginal])
		rotatedChars[posicionNueva] = char
	}
	return string(rotatedChars) // hay una compatibilidad entre []rune y string, tmb entre []byte y string, y de esta forma podemos transformar de un tipo a otro con "facilidad"
}

func RotateRightVersion3(text string) string {
	size := len(text)
	if size > 1 {
		primerCaracterDelStringRotado := string(text[size-1])
		elRestoDeLosCaracteresDelStringRotato := text[0 : size-1]
		return primerCaracterDelStringRotado + elRestoDeLosCaracteresDelStringRotato
	}
	return text
}

func RotateRightVersion4(textOriginal string) string { // textoOrigial = abc
	size := len(textOriginal)
	if size > 1 {
		primerCaracterDelStringRotado := string(textOriginal[size-1])
		textRotado := primerCaracterDelStringRotado // c
		for i := 0; i < size-1; i++ {
			textRotado = textRotado + string(textOriginal[i]) // 1: ca, 2: cab
		}
		return textRotado
	}
	return textOriginal
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

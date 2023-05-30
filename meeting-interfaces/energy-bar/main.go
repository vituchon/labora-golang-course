package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	animateBarGoAndBackUsingTwoForLoops()
	animateBarGoAnbBackUsingFunctions()
	animateBarGoAnbBackUsingInterfaces()
}

// SOLUCION 1	👷

func animateBarGoAndBackUsingTwoForLoops() {
	const col = 30
	barFormat := fmt.Sprintf("[%%-%vs]", col)
	for i := 0; i < col; i++ {
		fmt.Print("\033[H\033[2J")
		fmt.Printf(barFormat, strings.Repeat("=", i))
		time.Sleep(30 * time.Millisecond)
	}

	for i := col; i >= 0; i-- {
		fmt.Print("\033[H\033[2J")
		fmt.Printf(barFormat, strings.Repeat("=", i))
		time.Sleep(30 * time.Millisecond)
	}
	fmt.Println()
}

// SOLUCION 2 🤓
// Muy importante es primero ver la solución 1 y ver ambos ciclos for y detectar que partes varian y que partes son iguales!

func animateBarGoAnbBackUsingInterfaces() {
	const col = 30
	barFormat := fmt.Sprintf("[%%-%vs]", col)
	var sequencers []Sequencer = []Sequencer{
		&IncSequence{Higher: col, Current: 0},
		&DecSequence{Lower: 0, Current: col},
	}
	for _, sequencer := range sequencers {
		for sequencer.HasNext() {
			i := sequencer.Next()
			fmt.Print("\033[H\033[2J")
			fmt.Printf(barFormat, strings.Repeat("=", i))
			time.Sleep(30 * time.Millisecond)
		}
	}
}

type Sequencer interface {
	Next() int
	HasNext() bool
}

// esto podria ser una estructura que sirva para tanto incrementar o decrementar
type BaseSequence struct {
	Limit   int // el limit sirve de límite inferior para sequencia de decremento y como límite superior para sequencia de incremento
	Current int
}

type IncSequence struct {
	Higher  int
	Current int
}

func (is *IncSequence) Next() int {
	is.Current++
	return is.Current
}

func (is IncSequence) HasNext() bool {
	return is.Current < is.Higher
}

type DecSequence struct {
	Lower   int
	Current int
}

func (is *DecSequence) Next() int {
	is.Current--
	return is.Current
}

func (is DecSequence) HasNext() bool {
	return is.Current > is.Lower
}

// SOLUCION 3 😎 Hack level

func animateBarGoAnbBackUsingFunctions() {
	const col = 30
	var patterns []string = []string{"[%%-%vs]", "[%%%vs]"} // CLAVE: leer la documentación sirve para resolver problemas, acá https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/#string-or-byte-slice explican como justificar los valores a transformar en string!
	var strGeneratorFuncs []func(i int) string = []func(i int) string{
		func(i int) string {
			return strings.Repeat("=", i) + ">"
		},

		func(i int) string {
			return "<" + strings.Repeat("=", i)
		},
	}

	for j := 0; j < 2; j++ {
		bar := fmt.Sprintf(patterns[j], col)
		strGeneratorFunc := strGeneratorFuncs[j]
		for i := 0; i < col; i++ {
			fmt.Print("\033[H\033[2J")
			fmt.Printf(bar, strGeneratorFunc(i))
			time.Sleep(30 * time.Millisecond)
		}
	}

}

package main

import (
	"testing"
)

// para correr los benchmark
// go test -benchmem -bench=. github.com/vituchon/labora-golang-course/meeting-benchmark
// go test -bench=.  ./ // todos los benchs
// go test -bench=BenchmarkPrint . // los que sigan el patron
// se puede agregar flag -run=none (para que no ejecute tests!) => go test -run=none  -bench .  ./
// https://apuntes.de/golang/pruebas-benchmark/#gsc.tab=0
// https://www.golinuxcloud.com/golang-benchmark/

func BenchmarkPrintWhenIsSaturdaySwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printWhenIsSaturdaySwitch()
	}
}

func BenchmarkPrintWhenIsSaturdayIf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printWhenIsSaturdayIf()
	}
}

package main

/**
Con a&&uda de "graficar" una tabla de resultados para cada caso, hago una primera implementación

    x              | s1 | s2     | s3         | s4      | s5
0    < x <= 50     | x  | 0      | 0          | 0       | 0
50   < x <= 100    | 50 | x - 50 | 0          | 0       | 0
100  < x <= 700    | 50 | 50     | x - 100    | 0       | 0
700  < x <= 1500   | 50 | 50     | 600        | x - 700 | 0
1500 < x <= +inf   | 50 | 50     | 600        | 800     | x - 1500
*/
func SegmentarValorPorRangos(x int) (s1 int, s2 int, s3 int, s4 int, s5 int) {
	if x > 1500 {
		s1 = 50
		s2 = 50
		s3 = 600
		s4 = 800
		s5 = x - 1500
	}
	if x > 700 && x <= 1500 {
		s1 = 50
		s2 = 50
		s3 = 600
		s4 = x - 700
	}
	if x > 100 && x <= 700 {
		s1 = 50
		s2 = 50
		s3 = x - 100
	}
	if x > 50 && x <= 100 {
		s1 = 50
		s2 = x - 50
	}
	if x > 0 && x <= 50 {
		s1 = x - 0
	}
	return s1, s2, s3, s4, s5
}

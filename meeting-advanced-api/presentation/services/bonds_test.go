package services

import "testing"

var ids []int = []int{8} // a person id that actually has some bonds in the underlying postgres DB

// you may run this benchmark meeting-advanced-api$ go test -bench=Benchmark -run=none ./presentation/services/

func BenchmarkGetBondsOfUsingRepositories(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetBondsByPersonIdsUsingRepositories(ids)
	}
}

func BenchmarkGetBondsOfUsingDirectDBAccess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetBondsByPersonIdsDirectDBAccess(ids)
	}
}

// goos: linux
// goarch: amd64
// pkg: github.com/vituchon/labora-golang-course/meeting-advanced-api/presentation/services
// cpu: Intel(R) Core(TM) i7-3630QM CPU @ 2.40GHz
// BenchmarkGetBondsOfUsingRepositories-8     	     979	   1168552 ns/op
// BenchmarkGetBondsOfUsingDirectDBAccess-8   	    2139	    765467 ns/op <-- WINNER BY => 2139 / 979 = 2.1848825332
// PASS
// ok  	github.com/vituchon/labora-golang-course/meeting-advanced-api/presentation/services	4.862s

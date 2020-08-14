package main

import (
	"fmt"
	"strconv"
	"testing"
)

var num = 123456
var numstr = "123456"

func BenchmarkAtoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x, err := strconv.Atoi(numstr)
		if x != num || err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkStrconvParseInt(b *testing.B) {
	num64 := int64(num)
	for i := 0; i < b.N; i++ {
		x, err := strconv.ParseInt(numstr, 10, 64)
		if x != num64 || err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkFmtSscan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x int
		n, err := fmt.Sscanf(numstr, "%d", &x)
		if n != 1 || x != num || err != nil {
			b.Error(err)
		}
	}
}

// go test -bench=. atoi_test.go

/*

goos: darwin
goarch: amd64
BenchmarkStrconvParseInt-4       34717491               30.0 ns/op
BenchmarkAtoi-4                 100000000               10.6 ns/op
BenchmarkFmtSscan-4               1920675              614.0 ns/op
PASS
ok      command-line-arguments  4.842s

*/

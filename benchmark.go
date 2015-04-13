package benchmark

import (
	"flag"
	"fmt"
	"testing"
)

type DataStructure interface {
	Add(index int, element interface{})
	Get(index int) interface{}
}

var Elements int

func init() {
	flag.IntVar(&Elements, "elements", 100, "Amount of elements used for the benchmarks")
}

func AddElements(b *testing.B, d DataStructure) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < Elements; j++ {
			d.Add(j, "test")
		}
	}
}

func GetElements(b *testing.B, d DataStructure) {
	for i := 0; i < Elements; i++ {
		d.Add(i, fmt.Sprintf("Number %d", i))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < Elements; j++ {
			_ = d.Get(j)
		}
	}
}

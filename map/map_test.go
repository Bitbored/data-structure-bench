package aMap

import (
	"github.com/bitbored/data-structure-bench"
	"testing"
)

func BenchmarkAddElements(b *testing.B) {
	m := aMap(make(map[int]interface{}))

	benchmark.AddElements(b, m)
}

func BenchmarkGetElements(b *testing.B) {
	m := aMap(make(map[int]interface{}))

	benchmark.AddElements(b, m)

	b.ResetTimer()

	benchmark.GetElements(b, m)
}

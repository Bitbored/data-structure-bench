package slice

import (
	"github.com/bitbored/data-structure-bench"
	"testing"
)

func BenchmarkAddElements(b *testing.B) {
	s := slice(make([]interface{}, benchmark.Elements, benchmark.Elements))

	benchmark.AddElements(b, s)
}

func BenchmarkGetElements(b *testing.B) {
	s := slice(make([]interface{}, benchmark.Elements, benchmark.Elements))

	benchmark.AddElements(b, s)

	b.ResetTimer()

	benchmark.GetElements(b, s)
}

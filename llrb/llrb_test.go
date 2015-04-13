package llrb

import (
	"github.com/bitbored/data-structure-bench"
	"github.com/petar/GoLLRB/llrb"
	"testing"
)

func BenchmarkAddElements(b *testing.B) {
	tree := (*LLRB)(llrb.New())

	benchmark.AddElements(b, tree)
}

func BenchmarkGetElements(b *testing.B) {
	tree := (*LLRB)(llrb.New())

	benchmark.AddElements(b, tree)

	b.ResetTimer()

	benchmark.GetElements(b, tree)
}

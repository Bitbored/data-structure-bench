package bst

import (
	"fmt"
	"github.com/bitbored/data-structure-bench"
	"math"
	"testing"
)

func TestSingleElement(t *testing.T) {
	tree := new(BST)

	tree.Add(3, "test")

	if s := tree.Get(3).(string); s != "test" {
		t.Errorf("Found %s, expected to find test", s)
	}
}

func TestOverwriteElement(t *testing.T) {
	tree := new(BST)

	tree.Add(3, "hello")
	tree.Add(3, "test")

	if s := tree.Get(3).(string); s != "test" {
		t.Errorf("Found %s, expected to find test", s)
	}
}

func insert(i int, diff float64, tree *BST) {
	tree.Add(i, fmt.Sprintf("Number %d", i))

	prevDiff := math.Floor(diff + 0.5)
	if prevDiff == 0 {
		return
	}

	nextDiff := diff / 2.0
	insert(int(float64(i)-prevDiff), nextDiff, tree)
	insert(int(float64(i)+prevDiff), nextDiff, tree)

}

func insertBestCase(tree *BST) {
	insert(benchmark.Elements/2, float64(benchmark.Elements)/4.0, tree)
}

func TestBestCase(t *testing.T) {
	tree := new(BST)

	insertBestCase(tree)

	for i := 0; i < benchmark.Elements; i++ {
		expected := fmt.Sprintf("Number %d", i)

		s := tree.Get(i)
		if s == nil {
			t.Errorf("%s not found", expected)
			return
		}
		if s.(string) != expected {
			t.Errorf("Found %s, expected %s", s, expected)
		}
	}

}

func BenchmarkAddElements(b *testing.B) {
	tree := new(BST)

	benchmark.AddElements(b, tree)
}

func BenchmarkAddElementsBestCase(b *testing.B) {
	tree := new(BST)

	for i := 0; i < b.N; i++ {
		insertBestCase(tree)
	}
}

func BenchmarkGetElementsBestCase(b *testing.B) {
	tree := new(BST)

	insertBestCase(tree)

	b.ResetTimer()

	benchmark.GetElements(b, tree)
}

func BenchmarkGetElementsWorstCase(b *testing.B) {
	tree := new(BST)

	benchmark.AddElements(b, tree)

	b.ResetTimer()

	benchmark.GetElements(b, tree)
}

func BenchmarkFindElementsWorstCaseReverse(b *testing.B) {
	tree := new(BST)

	for i := benchmark.Elements - 1; i >= 0; i-- {
		tree.Add(i, fmt.Sprintf("Number %d", i))
	}

	b.ResetTimer()

	benchmark.GetElements(b, tree)
}

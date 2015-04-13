package llrb

import (
	"github.com/petar/GoLLRB/llrb"
)

type index int

func (i index) Less(than llrb.Item) bool {
	switch than.(type) {
	case index:
		return i < than.(index)
	case *data:
		return i < than.(*data).index
	}
	return false
}

type data struct {
	index   index
	element interface{}
}

func (d *data) Less(than llrb.Item) bool {
	switch than.(type) {
	case index:
		return d.index < than.(index)
	case *data:
		return d.index < than.(*data).index
	}
	return false
}

type LLRB llrb.LLRB

func (l *LLRB) Add(i int, element interface{}) {
	tree := (*llrb.LLRB)(l)
	tree.ReplaceOrInsert(&data{index(i), element})
}

func (l *LLRB) Get(i int) interface{} {
	tree := (*llrb.LLRB)(l)
	return tree.Get(index(i))
}

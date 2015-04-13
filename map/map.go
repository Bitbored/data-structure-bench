package aMap

type aMap map[int]interface{}

func (m aMap) Add(index int, element interface{}) {
	m[index] = element
}

func (m aMap) Get(index int) interface{} {
	return m[index]
}

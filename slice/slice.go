package slice

type slice []interface{}

func (s slice) Add(index int, element interface{}) {
	if len(s) <= index {
		t := make([]interface{}, (len(s)+1)*2, (len(s)+1)*2)
		copy(t, s)
		s = t
	}
	s[index] = element
}

func (s slice) Get(index int) interface{} {
	return s[index]
}

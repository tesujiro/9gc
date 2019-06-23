package parser

type Map struct {
	keys []string
	vals []interface{}
}

func NewMap() *Map {
	return &Map{
		keys: make([]string, 0),
		vals: make([]interface{}, 0),
	}
}

func (m *Map) Put(key string, val interface{}) {
	m.keys = append(m.keys, key)
	m.vals = append(m.vals, val)
}

func (m *Map) Get(key string) *interface{} {
	for i := len(m.keys) - 1; i >= 0; i-- {
		if m.keys[i] == key {
			return &m.vals[i]
		}
	}
	return nil
}

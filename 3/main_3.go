package main

type StringIntMap map[string]int

func NewStringIntMap() StringIntMap {
	return make(StringIntMap)
}

func (m StringIntMap) Add(key string, value int) {
	m[key] = value
}

func (m StringIntMap) Remove(key string) {
	delete(m, key)
}

func (m StringIntMap) Copy() StringIntMap {
	newMap := make(StringIntMap, len(m))
	for key, value := range m {
		newMap[key] = value
	}
	return newMap
}

func (m StringIntMap) Exists(key string) bool {
	_, exists := m[key]
	return exists
}

func (m StringIntMap) Get(key string) (int, bool) {
	value, exists := m[key]
	return value, exists
}
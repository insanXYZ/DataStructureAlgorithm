package main

type HashTable struct {
	Arrays [10]string
}

func (h *HashTable) hash(s string) (i int) {
	for _, v := range []byte(s) {
		i += int(v)
	}
	i = i % 10
	return
}

func (h *HashTable) insert(s string) {
	index := h.hash(s)
	h.Arrays[index] = s
}

func main() {
	h := HashTable{}
	h.insert("insan")
	h.insert("Bob")
}

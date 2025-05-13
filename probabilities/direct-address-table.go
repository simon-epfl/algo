package main

type Table struct {
	size int
	data map[int]int
}

// worst case en temps = O(1) pour toutes les opérations. super !

func (h *Table) DirectAddressSearch(key int) int {
	return h.data[key]
}

func (h *Table) DirectAddressInsert(key, value int) {
	h.data[key] = value
}

func (h *Table) DirectAddressDelete(key int) {
	delete(h.data, key)
}

// U est l'univers de clés
// K est le nombre de clés vraiment utilisées
// mais problème, en space on a du O(|U|) !

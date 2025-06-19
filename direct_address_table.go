package main

type CustomDirectTable struct {
	data map[int]int
}

// worst case en temps = O(1) pour toutes les opérations. super !

func (h *CustomDirectTable) DirectAddressSearch(key int) int {
	return h.data[key]
}

func (h *CustomDirectTable) DirectAddressInsert(key, value int) {
	h.data[key] = value
}

func (h *CustomDirectTable) DirectAddressDelete(key int) {
	delete(h.data, key)
}

// U est l'univers de clés
// K est le nombre de clés vraiment utilisées
// mais problème, en space on a du O(|Univers|) !

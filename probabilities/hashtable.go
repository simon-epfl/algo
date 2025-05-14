package main

type HashTable struct {
	size int
	data map[int]int
}

const ERR_NOT_FOUND = -1

// on veut en espace O(K), avec K le nombre de clés utilisées
// et pas O(|U|) avec U l'univers de clés

// on veut implémenter la recherche, l'insertion et la suppression en O(1) en termes de temps, en moyenne

func (h *HashTable) Hash(key int) int {
	return key % h.size
}

func (h *HashTable) HashSearch(key int) int {
	index := h.Hash(key)
	if value, ok := h.data[index]; ok {
		return value
	}
	return ERR_NOT_FOUND
}

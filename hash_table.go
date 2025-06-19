package main

type CustomHashTable struct {
	size int
	data []*CustomHashTableItem
}

type CustomHashTableItem struct {
	key   int
	value int
	next  *CustomHashTableItem
}

// on veut en espace O(K), avec K le nombre de clés utilisées
// et pas O(|U|) avec U l'univers de clés

// on veut implémenter la recherche, l'insertion et la suppression en O(1) en termes de temps, en moyenne

func NewCustomHashTable(size int) *CustomHashTable {
	return &CustomHashTable{
		size: size,
		data: make([]*CustomHashTableItem, size),
	}
}

func (h *CustomHashTable) Hash(key int) int {
	return key % h.size
}

func (h *CustomHashTable) HashSearch(key int) int {
	index := h.Hash(key)
	row := h.data[index]
	if row == nil {
		panic("hash table is empty or key not found")
	}

	for item := row; item != nil; item = item.next {
		if item.key == key {
			return item.value
		}
	}

	panic("key not found in hash table")

}

func (h *CustomHashTable) HashInsert(key, value int) {
	index := h.Hash(key)
	row := h.data[index]

	item := &CustomHashTableItem{
		key:   key,
		value: value,
		next:  nil,
	}

	if row != nil {
		h.data[index] = item
	} else {
		// on insère en tête de liste
		item.next = row
		h.data[index] = item
	}
}

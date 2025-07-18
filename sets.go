package main

type CustomDisjointSet[T comparable] struct {
	parent map[T]T
	rank   map[T]int
}

func NewCustomDisjointSet[T comparable]() *CustomDisjointSet[T] {
	return &CustomDisjointSet[T]{
		parent: make(map[T]T),
		rank:   make(map[T]int),
	}
}

func (sets *CustomDisjointSet[T]) MakeSet(x T) {
	sets.parent[x] = x // le root est son propre parent
	sets.rank[x] = 0
}

func (sets *CustomDisjointSet[T]) Find(x T) T {
	if _, ok := sets.parent[x]; !ok {
		panic("Element not found in disjoint set!")
	}
	if sets.parent[x] != x { // on remonte l'arbre jusqu'au root
		sets.parent[x] = sets.Find(sets.parent[x]) // path compression! on maj le parent au fur et à mesure
	}
	return sets.parent[x]
}

func (sets *CustomDisjointSet[T]) Union(u, v T) {
	representantU := sets.Find(u)
	representantV := sets.Find(v)
	if representantU == representantV {
		return
	}
	if sets.rank[representantU] < sets.rank[representantV] {
		sets.parent[representantU] = representantV
	} else if sets.rank[representantU] > sets.rank[representantV] {
		sets.parent[representantV] = representantU
	} else {
		sets.parent[representantV] = representantU
		sets.rank[representantU]++
	}
}

package main

import (
	sh "algo/shared"
	"sort"
)

func kruskal(vertices []*sh.Vertex, edges []*sh.Edge) []*sh.Edge {
	// Tri des arêtes par poids croissant
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	dsets := NewCustomDisjointSet[*sh.Vertex]()
	for _, vertex := range vertices {
		dsets.MakeSet(vertex)
	}

	minimumSpanningTree := make([]*sh.Edge, 0)

	for _, edge := range edges {
		if dsets.Find(edge.Origin) != dsets.Find(edge.Destination) {
			dsets.Union(edge.Origin, edge.Destination)
			minimumSpanningTree = append(minimumSpanningTree, edge)
		}
	}

	return minimumSpanningTree
}

func runKruskal() {

	// sommets
	a := &sh.Vertex{Name: "a"}
	b := &sh.Vertex{Name: "b"}
	c := &sh.Vertex{Name: "c"}
	d := &sh.Vertex{Name: "d"}
	e := &sh.Vertex{Name: "e"}
	f := &sh.Vertex{Name: "f"}
	g := &sh.Vertex{Name: "g"}
	h := &sh.Vertex{Name: "h"}
	i := &sh.Vertex{Name: "i"}

	vertices := []*sh.Vertex{a, b, c, d, e, f, g, h, i}

	edges := []*sh.Edge{
		{Weight: 10, Origin: a, Destination: b}, // a → b
		{Weight: 12, Origin: a, Destination: c}, // a → c
		{Weight: 9, Origin: b, Destination: c},  // b → c
		{Weight: 8, Origin: b, Destination: d},  // b → d
		{Weight: 1, Origin: c, Destination: f},  // c → f
		{Weight: 3, Origin: c, Destination: e},  // c → e
		{Weight: 7, Origin: d, Destination: e},  // d → e
		{Weight: 5, Origin: d, Destination: h},  // d → h
		{Weight: 3, Origin: e, Destination: f},  // e → f
		{Weight: 9, Origin: h, Destination: g},  // h → g
		{Weight: 2, Origin: g, Destination: i},  // g → i
		{Weight: 11, Origin: h, Destination: i}, // h → i
		{Weight: 6, Origin: h, Destination: f},  // h → f
		{Weight: 8, Origin: d, Destination: g},  // d → g
	}

	minimumSpanningTree := kruskal(vertices, edges)
	for _, edge := range minimumSpanningTree {
		sh.PrintEdge(edge)
	}
}

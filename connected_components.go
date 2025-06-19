package main

import (
	sh "algo/shared"
	"fmt"
)

func connectedComponents() {

	a := &sh.Vertex{Name: "a"}
	b := &sh.Vertex{Name: "b"}
	c := &sh.Vertex{Name: "c"}

	vertices := []*sh.Vertex{a, b, c}

	edges := []*sh.Edge{
		{Weight: 1, Origin: a, Destination: b}, // a → b
		{Weight: 1, Origin: b, Destination: c}, // b → c
	}

	dsets := NewCustomDisjointSet[*sh.Vertex]()
	for _, vertex := range vertices {
		dsets.MakeSet(vertex)
	}
	for _, edge := range edges {
		if dsets.Find(edge.Origin) != dsets.Find(edge.Destination) {
			dsets.Union(edge.Origin, edge.Destination)
		}
	}

	groups := make(map[*sh.Vertex][]*sh.Vertex)
	for _, v := range vertices {
		rep := dsets.Find(v)
		groups[rep] = append(groups[rep], v)
	}

	for rep, members := range groups {
		fmt.Printf("Représentant : %v\n", rep.Name)
		fmt.Println("Membres :")
		for _, m := range members {
			if m != rep {
				fmt.Printf(" - %v\n", m.Name)
			}
		}
		fmt.Println()
	}

}

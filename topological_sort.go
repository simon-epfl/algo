package main

import (
	sh "algo/shared"
	"fmt"
	"sort"
)

func topologicalSort() {

	fmt.Println("DFS")

	// https://moodle.epfl.ch/pluginfile.php/3434625/mod_resource/content/1/algorithms%20I%20-%20Lecture%2013.pdf
	// slide 43

	s := &sh.Vertex{Name: "undershorts"}
	a := &sh.Vertex{Name: "pants"}
	b := &sh.Vertex{Name: "shoes"}
	c := &sh.Vertex{Name: "socks"}
	d := &sh.Vertex{Name: "watch"}
	e := &sh.Vertex{Name: "belt"}
	f := &sh.Vertex{Name: "shirt"}
	g := &sh.Vertex{Name: "tie"}
	h := &sh.Vertex{Name: "jacket"}

	vertices := []*sh.Vertex{f, s, a, b, c, d, e, g, h}

	edges := []*sh.Edge{
		{Weight: 1, Origin: s, Destination: a}, // undershorts → pants
		{Weight: 1, Origin: s, Destination: b}, // undershorts → shoes
		{Weight: 1, Origin: a, Destination: e}, // pants → belt
		{Weight: 1, Origin: a, Destination: b}, // pants → shoes
		{Weight: 1, Origin: c, Destination: b}, // socks → shoes
		{Weight: 1, Origin: f, Destination: e}, // shirt → belt
		{Weight: 1, Origin: f, Destination: g}, // shirt → tie
		{Weight: 1, Origin: g, Destination: h}, // tie → jacket
		{Weight: 1, Origin: e, Destination: h}, // belt → jacket
	}

	sh.InitializeSingleSource(vertices, s)

	adjacents := make(map[*sh.Vertex][]*sh.Vertex)
	for _, edge := range edges {
		adjacents[edge.Origin] = append(adjacents[edge.Origin], edge.Destination)
	}

	for _, vertex := range vertices {
		vertex.Color = sh.ColorWhite
		vertex.DiscoveryTime = 0
		vertex.FinishTime = 0
	}

	time := 0
	for _, vertex := range vertices {
		if vertex.Color == sh.ColorWhite {
			dfsVisit(vertex, adjacents, &time)
		}
	}

	fmt.Println("Topological sorting order:")
	// on trie par ordre de finish time décroissant
	sort.Slice(vertices, func(i, j int) bool {
		return vertices[i].FinishTime > vertices[j].FinishTime
	})
	for _, vertex := range vertices {
		fmt.Printf("%s (discovery: %d, finish: %d)\n", vertex.Name, vertex.DiscoveryTime, vertex.FinishTime)
	}
	fmt.Println()

}

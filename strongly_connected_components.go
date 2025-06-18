package main

import (
	sh "algo/shared"
	"fmt"
	"sort"
)

func sccDfsVisit(vertex *sh.Vertex, adjacents map[*sh.Vertex][]*sh.Vertex, time *int) {
	*time++
	vertex.DiscoveryTime = *time
	vertex.Color = sh.ColorGray
	fmt.Printf("Visiting vertex %s\n", vertex.Name)

	for _, adjacent := range adjacents[vertex] {
		if adjacent.Color == sh.ColorWhite {
			adjacent.Predecessor = vertex
			sccDfsVisit(adjacent, adjacents, time)
		}
	}
	vertex.Color = sh.ColorBlack
	*time++
	vertex.FinishTime = *time
}

func sccDfs(vertices []*sh.Vertex, edges []*sh.Edge) {

	adjacents := make(map[*sh.Vertex][]*sh.Vertex)
	for _, edge := range edges {
		adjacents[edge.Origin] = append(adjacents[edge.Origin], edge.Destination)
	}

	for _, vertex := range vertices {
		vertex.Color = sh.ColorWhite
		vertex.DiscoveryTime = 0
		vertex.FinishTime = 0
		vertex.Predecessor = nil
	}

	time := 0
	for _, vertex := range vertices {
		if vertex.Color == sh.ColorWhite {
			sccDfsVisit(vertex, adjacents, &time)

			fmt.Printf("=== Arbre DFS racine %s ===\n", vertex.Name)
			printTree(vertex, vertices)
			fmt.Println()
		}
	}
}

func stronglyConnectedComponents() {

	// https://moodle.epfl.ch/pluginfile.php/3434625/mod_resource/content/1/algorithms%20I%20-%20Lecture%2013.pdf
	// slide 43

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
	j := &sh.Vertex{Name: "j"}

	vertices := []*sh.Vertex{a, b, c, d, e, f, g, h, i, j}

	edges := []*sh.Edge{
		{Weight: 1, Origin: a, Destination: b}, // a → b
		{Weight: 1, Origin: a, Destination: f}, // a → f
		{Weight: 1, Origin: b, Destination: c}, // b → c
		{Weight: 1, Origin: b, Destination: g}, // b → g
		{Weight: 1, Origin: c, Destination: d}, // c → d
		{Weight: 1, Origin: c, Destination: h}, // c → h
		{Weight: 1, Origin: d, Destination: e}, // d → e
		{Weight: 1, Origin: d, Destination: j}, // d → j
		{Weight: 1, Origin: d, Destination: i}, // d → i
		{Weight: 1, Origin: d, Destination: h}, // d → h
		{Weight: 1, Origin: e, Destination: j}, // e → j
		{Weight: 1, Origin: f, Destination: g}, // f → g
		{Weight: 1, Origin: g, Destination: h}, // g → h
		{Weight: 1, Origin: g, Destination: a}, // g → a
		{Weight: 1, Origin: h, Destination: c}, // h → c
		{Weight: 1, Origin: i, Destination: j}, // i → j
		{Weight: 1, Origin: j, Destination: e}, // j → e
	}

	sccDfs(vertices, edges)
	fmt.Println("Parcours DFS terminé, maintenant on va transposer le graphe")

	verticesTransposed := make([]*sh.Vertex, len(vertices))
	copy(verticesTransposed, vertices)
	edgesTransposed := make([]*sh.Edge, len(edges))
	for k, edge := range edges {
		edgesTransposed[k] = &sh.Edge{
			Weight:      edge.Weight,
			Origin:      verticesTransposed[indexOf(vertices, edge.Destination)],
			Destination: verticesTransposed[indexOf(vertices, edge.Origin)],
		}
	}

	// on doit changer l'ordre des sommets pour le DFS sur le graphe transposé
	// on trie les sommets par ordre de finish time décroissant
	sort.Slice(verticesTransposed, func(i, j int) bool {
		return verticesTransposed[i].FinishTime > verticesTransposed[j].FinishTime
	})

	sccDfs(verticesTransposed, edgesTransposed)
	fmt.Println("Parcours DFS terminé sur le graphe transposé")

	fmt.Println("Composantes fortement connexes :")
	for _, vertex := range verticesTransposed {
		if vertex.Color == sh.ColorBlack { // on affiche seulement les sommets qui ont été visités dans la deuxième passe
			fmt.Printf("Composante fortement connexe : %s (temps de découverte = %d, temps de finition = %d)\n", vertex.Name, vertex.DiscoveryTime, vertex.FinishTime)
		}
	}
	fmt.Println("Composantes fortement connexes terminées")

}

func indexOf(slice []*sh.Vertex, v *sh.Vertex) int {
	for i, u := range slice {
		if u == v {
			return i
		}
	}
	return -1
}

func printTree(u *sh.Vertex, all []*sh.Vertex) {
	var recurse func(v *sh.Vertex, indent string)
	recurse = func(v *sh.Vertex, indent string) {
		fmt.Printf("%s- %s (d=%d, f=%d)\n", indent, v.Name, v.DiscoveryTime, v.FinishTime)
		for _, w := range all {
			if w.Predecessor == v {
				recurse(w, indent+"  ")
			}
		}
	}
	recurse(u, "")
}

package main

import (
	sh "algo/shared"
	"fmt"
)

/*
 * Tree edge (u→v) : on découvre v pour la première fois (couleur blanche) en explorant u.
 *
 * Back edge (u→v) : on voit un arc vers un sommet déjà en cours de visite (couleur grise).
 * Autrement dit, un retour vers un ancêtre dans la pile d’appel.
 *
 * Forward edge (u→v) : on voit un arc vers un descendant déjà complètement visité (couleur noire),
 * c’est-à-dire un nœud qui était dans notre sous-arbre DFS, mais qu’on a déjà fini de traiter.
 *
 * Cross edge (u→v) : on voit un arc vers un sommet déjà complètement visité (couleur noire) qui n’est ni ancêtre,
 * ni descendant de u dans l’arbre DFS (autre branche ou composante).
 */

func dfsVisit(vertex *sh.Vertex, adjacents map[*sh.Vertex][]*sh.Vertex, time *int) {
	*time++
	vertex.DiscoveryTime = *time
	vertex.Color = sh.ColorGray
	fmt.Printf("Visiting vertex %s\n", vertex.Name)

	for _, adjacent := range adjacents[vertex] {
		if adjacent.Color == sh.ColorWhite {
			dfsVisit(adjacent, adjacents, time)
		}
	}
	vertex.Color = sh.ColorBlack
	*time++
	vertex.FinishTime = *time
}

func dfs(vertices []*sh.Vertex, edges []*sh.Edge) {

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
}

func runDfs() {

	fmt.Println("DFS")

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

	vertices := []*sh.Vertex{b, a, c, d, e, f, g, h}

	// arêtes orientées (poids = 1 partout)
	edges := []*sh.Edge{
		{Weight: 1, Origin: a, Destination: g}, // a → g
		{Weight: 1, Origin: a, Destination: h}, // a → h

		{Weight: 1, Origin: b, Destination: a}, // b → a
		{Weight: 1, Origin: b, Destination: c}, // b → c
		{Weight: 1, Origin: b, Destination: g}, // b → g

		{Weight: 1, Origin: c, Destination: d}, // c → d
		{Weight: 1, Origin: c, Destination: g}, // c → g

		{Weight: 1, Origin: d, Destination: g}, // d → g

		{Weight: 1, Origin: e, Destination: d}, // e → d
		{Weight: 1, Origin: e, Destination: f}, // e → f
		{Weight: 1, Origin: e, Destination: c}, // e → c

		{Weight: 1, Origin: g, Destination: h}, // g → h

		{Weight: 1, Origin: f, Destination: g}, // f → g
	}

	sh.InitializeSingleSource(vertices, b)

	dfs(vertices, edges)

	fmt.Println("Temps de découverte et de finition :")
	for _, vertex := range vertices {
		fmt.Printf("Sommet %s : temps de découverte = %d, temps de finition = %d\n", vertex.Name, vertex.DiscoveryTime, vertex.FinishTime)
	}
	fmt.Println("Parcours DFS terminé")

}

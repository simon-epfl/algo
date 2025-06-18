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

func dfs() {

	fmt.Println("DFS")

	// https://moodle.epfl.ch/pluginfile.php/3434625/mod_resource/content/1/algorithms%20I%20-%20Lecture%2013.pdf
	// slide 43

	s := &sh.Vertex{Name: "s"}
	a := &sh.Vertex{Name: "a"}
	b := &sh.Vertex{Name: "b"}
	c := &sh.Vertex{Name: "c"}
	d := &sh.Vertex{Name: "d"}
	e := &sh.Vertex{Name: "e"}
	f := &sh.Vertex{Name: "f"}
	g := &sh.Vertex{Name: "g"}
	h := &sh.Vertex{Name: "h"}

	vertices := []*sh.Vertex{s, a, b, c, d, e, f, g, h}

	edges := []*sh.Edge{
		{Weight: 1, Origin: s, Destination: c}, // s → c
		{Weight: 1, Origin: s, Destination: a}, // s → a
		{Weight: 1, Origin: a, Destination: d}, // a → d
		{Weight: 1, Origin: b, Destination: d}, // b → d
		{Weight: 1, Origin: b, Destination: a}, // b → a
		{Weight: 1, Origin: c, Destination: d}, // c → d
		{Weight: 1, Origin: d, Destination: b}, // d → b
		{Weight: 1, Origin: c, Destination: f}, // c → f
		{Weight: 1, Origin: f, Destination: e}, // f → e
		{Weight: 1, Origin: f, Destination: h}, // f → h
		{Weight: 1, Origin: f, Destination: g}, // f → g
		{Weight: 1, Origin: g, Destination: f}, // g → f
		{Weight: 1, Origin: h, Destination: g}, // h → g
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

	fmt.Println("Temps de découverte et de finition :")
	for _, vertex := range vertices {
		fmt.Printf("Sommet %s : temps de découverte = %d, temps de finition = %d\n", vertex.Name, vertex.DiscoveryTime, vertex.FinishTime)
	}
	fmt.Println("Parcours DFS terminé")

}

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	name     string
	distance int
	previous *Vertex

	isInfinity bool
}

type Edge struct {
	weight      int
	origin      *Vertex
	destination *Vertex
}

/*
 * Relax prend deux noeuds, et un poids.
 * Si le noeud V de destination de notre edge
 * peut être atteint avec une distance plus courte en utilisant le current edge
 * alors on l'utilise
 */
func relax(edge *Edge) bool {
	fmt.Printf("Tentative de relaxation de %s vers %s avec un poids de %d\n", edge.origin.name, edge.destination.name, edge.weight)
	if edge.origin.isInfinity {
		return false
	}
	if edge.destination.distance > edge.origin.distance+edge.weight {
		fmt.Printf("On a trouvé un chemin plus court de %s vers %s avec un poids de %d\n", edge.origin.name, edge.destination.name, edge.weight)
		edge.destination.distance = edge.origin.distance + edge.weight
		edge.destination.previous = edge.origin
		edge.destination.isInfinity = false
		return true
	}
	return false
}

func initializeSingleSource(vertices []*Vertex, source *Vertex) {
	for _, vertex := range vertices {
		vertex.distance = math.MaxInt
		vertex.previous = nil
		vertex.isInfinity = true
	}
	source.distance = 0
	source.isInfinity = false
}

func hasNoNegativeCycles(edges []*Edge) bool {
	for _, edge := range edges {

		// on est censé avoir une solution optimale mais on trouve encore une meilleure solution
		// ==> negative cycle
		if edge.destination.distance > edge.origin.distance+edge.weight {
			return false
		}
	}
	return true
}

func main() {

	fmt.Println("Bellman-Ford Algorithm")

	s, t, x, y, z := &Vertex{name: "s"}, &Vertex{name: "t"}, &Vertex{name: "x"}, &Vertex{name: "y"}, &Vertex{name: "z"}
	vertices := []*Vertex{s, t, x, y, z}

	initializeSingleSource(vertices, z)

	edges := []*Edge{
		{weight: 5, origin: t, destination: x},
		{weight: 8, origin: t, destination: y},
		{weight: -4, origin: t, destination: z},
		{weight: -2, origin: x, destination: t},
		{weight: -3, origin: y, destination: x},
		{weight: 9, origin: y, destination: z},
		{weight: 7, origin: z, destination: x},
		{weight: 2, origin: z, destination: s},
		{weight: 6, origin: s, destination: t},
		{weight: 7, origin: s, destination: y},
	}

	hasAtLeastAnUpdateLastCycle := false
	// on a besoin de faire n-1 iterations
	// parce qu'en fait le shortest path d'un graphe avec n sommets
	// dans le pire des cas, est composé de n-1 edges (on veut pas de cycle)

	// imaginer le pire cas est une droite comme S-->B-->C-->D-->E-->T
	// et on commence par (E, T) et on finit par (S, B)
	// donc pour la première itération, on aura rien update sauf la dernière
	// (parce que (E, T) comme on connaît pas E de toute façon, on peut pas update T)
	for i := 0; i < len(vertices)-1; i++ {
		for _, edge := range edges {
			update := relax(edge)
			if update {
				hasAtLeastAnUpdateLastCycle = true
			}
		}
		if !hasAtLeastAnUpdateLastCycle {
			break
		}
		hasAtLeastAnUpdateLastCycle = false
	}

	if !hasNoNegativeCycles(edges) {
		fmt.Println("On a un negative cycle!")
	}

	fmt.Printf("Distance de z à s est %d\n", s.distance)
	fmt.Printf("Distance de z à t est %d\n", t.distance)
	fmt.Printf("Distance de z à x est %d\n", x.distance)
	fmt.Printf("Distance de z à y est %d\n", y.distance)

}

package main

import (
	sh "algo/shared"
	"fmt"
)

func hasNoNegativeCycles(edges []*sh.Edge) bool {
	for _, edge := range edges {

		// on est censé avoir une solution optimale mais on trouve encore une meilleure solution
		// ==> negative cycle
		if edge.Destination.Distance > edge.Origin.Distance+edge.Weight {
			return false
		}
	}
	return true
}

func bellmanford() {

	fmt.Println("Bellman-Ford Algorithm")

	s, t, x, y, z := &sh.Vertex{Name: "s"}, &sh.Vertex{Name: "t"}, &sh.Vertex{Name: "x"}, &sh.Vertex{Name: "y"}, &sh.Vertex{Name: "z"}
	vertices := []*sh.Vertex{s, t, x, y, z}

	sh.InitializeSingleSource(vertices, z)

	edges := []*sh.Edge{
		{Weight: 5, Origin: t, Destination: x},
		{Weight: 8, Origin: t, Destination: y},
		{Weight: -4, Origin: t, Destination: z},
		{Weight: -2, Origin: x, Destination: t},
		{Weight: -3, Origin: y, Destination: x},
		{Weight: 9, Origin: y, Destination: z},
		{Weight: 7, Origin: z, Destination: x},
		{Weight: 2, Origin: z, Destination: s},
		{Weight: 6, Origin: s, Destination: t},
		{Weight: 7, Origin: s, Destination: y},
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
			update := sh.Relax(edge)
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

	fmt.Printf("Distance de z à s est %d\n", s.Distance)
	fmt.Printf("Distance de z à t est %d\n", t.Distance)
	fmt.Printf("Distance de z à x est %d\n", x.Distance)
	fmt.Printf("Distance de z à y est %d\n", y.Distance)

}

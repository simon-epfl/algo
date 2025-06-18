package shared

import (
	"fmt"
	"math"
)

type Color int

const (
	ColorWhite = iota
	ColorGray
	ColorBlack
)

type Vertex struct {
	Name     string
	Distance int
	Previous *Vertex

	IsInfinity bool

	// DFS
	DiscoveryTime int
	FinishTime    int
	Color         Color
}

type Edge struct {
	Weight      int
	Origin      *Vertex
	Destination *Vertex
}

/*
 * Relax prend deux noeuds, et un poids.
 * Si le noeud V de destination de notre edge
 * peut être atteint avec une distance plus courte en utilisant le current edge
 * alors on l'utilise
 */
func Relax(edge *Edge) bool {
	fmt.Printf("Tentative de relaxation de %s vers %s avec un poids de %d\n", edge.Origin.Name, edge.Destination.Name, edge.Weight)
	if edge.Origin.IsInfinity {
		return false
	}
	if edge.Destination.Distance > edge.Origin.Distance+edge.Weight {
		fmt.Printf("On a trouvé un chemin plus court de %s vers %s avec un poids de %d\n", edge.Origin.Name, edge.Destination.Name, edge.Weight)
		edge.Destination.Distance = edge.Origin.Distance + edge.Weight
		edge.Destination.Previous = edge.Origin
		edge.Destination.IsInfinity = false
		return true
	}
	return false
}

func InitializeSingleSource(vertices []*Vertex, source *Vertex) {
	for _, vertex := range vertices {
		vertex.Distance = math.MaxInt
		vertex.Previous = nil
		vertex.IsInfinity = true
	}
	source.Distance = 0
	source.IsInfinity = false
}

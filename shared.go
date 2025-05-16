package main

import "fmt"

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

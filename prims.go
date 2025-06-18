package main

import (
	sh "algo/shared"
	"container/heap"
	"fmt"
	"math"
)

func runPrims() {
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

	adjacents := make(map[*sh.Vertex][]*sh.Vertex)

	for _, edge := range edges {
		adjacents[edge.Origin] = append(adjacents[edge.Origin], edge.Destination)
		adjacents[edge.Destination] = append(adjacents[edge.Destination], edge.Origin)
	}

	initialVertex := a

	queue := sh.PriorityQueue(make([]*sh.PriorityItem, 0, len(vertices)))
	heap.Init(&queue)

	// Map pour garder une référence aux items dans la queue
	itemMap := make(map[*sh.Vertex]*sh.PriorityItem)

	for _, vertex := range vertices {
		item := &sh.PriorityItem{
			Value:    vertex,
			Priority: math.MaxInt,
		}
		if vertex == initialVertex {
			item.Priority = 0
		}
		vertex.Color = sh.ColorGray // dans la queue
		heap.Push(&queue, item)
		itemMap[vertex] = item
	}

	// on veut stocker les edges du graphe minimum
	minimumSpanningTreeEdges := make([]*sh.Edge, 0)

	// va garder en mémoire la meilleure arête pour chaque sommet
	// celle qui est connectée au graphe minimum
	// et la plus proche
	bestEdge := make(map[*sh.Vertex]*sh.Edge)

	for len(queue) > 0 {
		// On récupère le sommet avec la plus petite distance
		item := heap.Pop(&queue).(*sh.PriorityItem)
		vertex := item.Value.(*sh.Vertex)

		// On ignore les sommets déjà traités (peuvent être présents plusieurs fois dans la queue)
		if vertex.Color == sh.ColorBlack {
			continue
		}

		vertex.Color = sh.ColorBlack // Marquer le sommet comme traité

		fmt.Printf("On traite le sommet %s avec une priorité de %d\n", vertex.Name, item.Priority)

		if e := bestEdge[vertex]; e != nil {
			minimumSpanningTreeEdges = append(minimumSpanningTreeEdges, bestEdge[vertex])
		}

		for _, neighbor := range adjacents[vertex] {
			if neighbor.Color == sh.ColorBlack {
				fmt.Printf("Le voisin %s de %s est déjà traité, on l'ignore\n", neighbor.Name, vertex.Name)
				continue
			}
			var edge *sh.Edge
			for _, e := range edges {
				if (e.Origin == vertex && e.Destination == neighbor) || (e.Destination == vertex && e.Origin == neighbor) {
					edge = e
					break
				}
			}
			neigh := itemMap[neighbor]
			fmt.Printf("On a un voisin (%s) de %s sur un chemin de poids %d (priorité %d)\n", neighbor.Name, vertex.Name, edge.Weight, neigh.Priority)
			if (neighbor.Color == sh.ColorGray && edge != nil) && (edge.Weight < neigh.Priority) {
				fmt.Printf("Le poids de l'arête %s --(%d)--> %s est inférieur à la priorité actuelle %d\n", vertex.Name, edge.Weight, neighbor.Name, neigh.Priority)
				// Mettre à jour la priorité du voisin
				fmt.Printf("Mise à jour de la priorité de %s avec le poids %d\n", neighbor.Name, edge.Weight)
				queue.Update(neigh, neigh.Value, edge.Weight)
				bestEdge[neighbor] = edge
			}
		}
	}

	fmt.Println("Arêtes du graphe minimum :")
	for _, edge := range minimumSpanningTreeEdges {
		fmt.Printf("%s --(%d)--> %s\n", edge.Origin.Name, edge.Weight, edge.Destination.Name)
	}
	fmt.Println("Graphe minimum construit avec Prim's algorithm.")

}

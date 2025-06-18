package main

import (
	sh "algo/shared"
	"fmt"
	"math"
)

// en fait cette structure va stocker toutes nos étapes
// pour reconstruire le chemin augmentant
// on en a besoin parce qu'une fois qu'on a trouvé ce chemin,
// faut calculer le bottleneck et mettre à jour les flux
type PathStep struct {
	edge     *sh.Edge
	backward bool
}

func findAugmentingPath(source, puits *sh.Vertex, adjacents map[*sh.Vertex][]*sh.Edge) []PathStep {
	parent := make(map[*sh.Vertex]PathStep)
	visited := map[*sh.Vertex]bool{source: true}
	queue := getEmptyCustomQueue()
	queue.Enqueue(CustomQueueItem{
		Content: source,
	})

	for !queue.IsEmpty() {
		u := queue.Dequeue().Content.(*sh.Vertex)
		for _, edge := range adjacents[u] {
			var vertex *sh.Vertex
			var back bool

			// forward ?
			if edge.Origin == u && edge.Flow < edge.Capacity {
				vertex, back = edge.Destination, false
				// backward ?
			} else if edge.Destination == u && edge.Flow > 0 {
				vertex, back = edge.Origin, true
			} else {
				continue
			}

			if visited[vertex] {
				continue
			}
			visited[vertex] = true
			parent[vertex] = PathStep{edge: edge, backward: back}

			if vertex == puits {
				// on reconstruit le chemin
				var path []PathStep
				for current := puits; current != source; {
					step := parent[current]
					path = append([]PathStep{step}, path...)
					if step.backward {
						current = step.edge.Destination
					} else {
						current = step.edge.Origin
					}
				}
				return path
			}
			queue.Enqueue(CustomQueueItem{
				Content: vertex,
			})
		}
	}
	return nil
}

func fordFulkerson() {

	source := &sh.Vertex{Name: "s"}
	u := &sh.Vertex{Name: "u"}
	puits := &sh.Vertex{Name: "t"}

	edges := []*sh.Edge{
		{Capacity: 5, Origin: source, Destination: u}, // s → u
		{Capacity: 3, Origin: u, Destination: puits},  // u → t
	}

	adjacents := make(map[*sh.Vertex][]*sh.Edge)
	for _, edge := range edges {
		edge.Flow = 0
		adjacents[edge.Origin] = append(adjacents[edge.Origin], edge)
		adjacents[edge.Destination] = append(adjacents[edge.Destination], edge)
	}

	maxFlow := 0
	for {
		path := findAugmentingPath(source, puits, adjacents)
		if path == nil {
			break
		}
		// Cool ! on a trouvé un nouveau chemin augmentant
		// trouver le goulot pour savoir de combien on peut augmenter le flux
		bottleneck := math.MaxInt
		for _, step := range path {
			if !step.backward {
				if res := step.edge.Capacity - step.edge.Flow; res < bottleneck {
					bottleneck = res
				}
			} else {
				if step.edge.Flow < bottleneck {
					bottleneck = step.edge.Flow
				}
			}
		}
		// on a le bottleneck, on peut mettre à jour les flux
		for _, step := range path {
			if !step.backward {
				step.edge.Flow += bottleneck
			} else {
				step.edge.Flow -= bottleneck
			}
		}
		maxFlow += bottleneck
	}

	fmt.Printf("Le flux maximum est de %d\n", maxFlow)
}

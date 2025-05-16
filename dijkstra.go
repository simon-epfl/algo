package main

import (
	"fmt"
)

func dijkstra() {

	fmt.Println("Dijkstra Algorithm")

	s, t, x, y, z := &Vertex{name: "s"}, &Vertex{name: "t"}, &Vertex{name: "x"}, &Vertex{name: "y"}, &Vertex{name: "z"}
	vertices := []*Vertex{s, t, x, y, z}

	initializeSingleSource(vertices, z)

	edges := []*Edge{
		{weight: 6, origin: t, destination: x},
		{weight: 2, origin: t, destination: y},
		{weight: 3, origin: s, destination: t},
		{weight: 5, origin: s, destination: y},
		{weight: 6, origin: y, destination: z},
		{weight: 4, origin: y, destination: x},
		{weight: 2, origin: x, destination: z},
		{weight: 7, origin: z, destination: x},
		{weight: 3, origin: z, destination: s},
	}

	h := &PriorityQueue{}

	for _, vertex := range vertices {
		h.Push(&PriorityItem{value: vertex, priority: vertex.distance})
	}

	for h.Len() > 0 {
		item := h.Pop().(*PriorityItem)
		vertex := item.value.(*Vertex)
		fmt.Printf("On traite le sommet %s avec une distance de %d\n", vertex.name, vertex.distance)

		for _, edge := range edges {
			if edge.origin == vertex {
				relax(edge)
				h.update(item, vertex, edge.destination.distance)
			}
		}
	}

	fmt.Println("Distances finales :")
	for _, vertex := range vertices {
		fmt.Printf("Sommet %s : distance = %d\n", vertex.name, vertex.distance)
	}

}

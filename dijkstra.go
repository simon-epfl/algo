package main

import (
	sh "algo/shared"
	"fmt"
)

func dijkstra() {

	fmt.Println("Dijkstra Algorithm")

	s, t, x, y, z := &sh.Vertex{Name: "s"}, &sh.Vertex{Name: "t"}, &sh.Vertex{Name: "x"}, &sh.Vertex{Name: "y"}, &sh.Vertex{Name: "z"}
	vertices := []*sh.Vertex{s, t, x, y, z}

	sh.InitializeSingleSource(vertices, z)

	edges := []*sh.Edge{
		{Weight: 6, Origin: t, Destination: x},
		{Weight: 2, Origin: t, Destination: y},
		{Weight: 3, Origin: s, Destination: t},
		{Weight: 5, Origin: s, Destination: y},
		{Weight: 6, Origin: y, Destination: z},
		{Weight: 4, Origin: y, Destination: x},
		{Weight: 2, Origin: x, Destination: z},
		{Weight: 7, Origin: z, Destination: x},
		{Weight: 3, Origin: z, Destination: s},
	}

	h := &sh.PriorityQueue{}

	for _, vertex := range vertices {
		h.Push(&sh.PriorityItem{Value: vertex, Priority: vertex.Distance})
	}

	for h.Len() > 0 {
		item := h.Pop().(*sh.PriorityItem)
		vertex := item.Value.(*sh.Vertex)
		fmt.Printf("On traite le sommet %s avec une distance de %d\n", vertex.Name, vertex.Distance)

		for _, edge := range edges {
			if edge.Origin == vertex {
				sh.Relax(edge)
				h.Update(item, vertex, edge.Destination.Distance)
			}
		}
	}

	fmt.Println("Distances finales :")
	for _, vertex := range vertices {
		fmt.Printf("Sommet %s : distance = %d\n", vertex.Name, vertex.Distance)
	}

}

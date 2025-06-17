package main

import (
	sh "algo/shared"
	"fmt"
)

func bfs() {

	fmt.Println("BFS")

	// https://moodle.epfl.ch/pluginfile.php/3434625/mod_resource/content/1/algorithms%20I%20-%20Lecture%2013.pdf
	// slide 41

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
		// de la partie gauche
		{Weight: 1, Origin: s, Destination: c}, // s → c
		{Weight: 1, Origin: s, Destination: a}, // s → a
		{Weight: 1, Origin: a, Destination: d}, // a → d
		{Weight: 1, Origin: b, Destination: d}, // b → d
		{Weight: 1, Origin: b, Destination: a}, // b → a
		{Weight: 1, Origin: c, Destination: d}, // c → d
		{Weight: 1, Origin: d, Destination: b}, // d → b

		// de la partie droite
		{Weight: 1, Origin: c, Destination: f}, // c → f
		{Weight: 1, Origin: f, Destination: e}, // f → e
		{Weight: 1, Origin: f, Destination: h}, // f → h

		{Weight: 1, Origin: f, Destination: g}, // f → g
		{Weight: 1, Origin: g, Destination: f}, // g → f

		{Weight: 1, Origin: h, Destination: g}, // h → g
	}

	for _, v := range vertices {
		v.Distance = 999999
	}

	queue := getEmptyCustomQueue()

	queue.Enqueue(CustomQueueItem{
		Content: s,
	})

	s.Distance = 0

	// on veut calculer la distance entre s et tous les vertices

	for !queue.IsEmpty() {
		item := queue.Dequeue()
		vertex := item.Content.(*sh.Vertex)
		fmt.Printf("On traite le sommet %s (distance de %d)\n", vertex.Name, vertex.Distance)

		// trouver tous les voisins du vertex courant
		for _, edge := range edges {
			if edge.Origin == vertex {
				// on a trouvé un voisin!
				neighbor := edge.Destination
				fmt.Printf("On a trouvé %s --> %s (distance de %d)\n", vertex.Name, neighbor.Name, neighbor.Distance)

				if neighbor.Distance == 999999 { // si on n'a pas encore visité ce voisin
					neighbor.Distance = vertex.Distance + 1
					queue.Enqueue(CustomQueueItem{
						Content: neighbor,
					})
					fmt.Printf("On sauvegarde source %s --> %s (distance de %d)\n", s.Name, neighbor.Name, neighbor.Distance)
				} else {
					fmt.Printf("On a déjà visité le voisin %s (distance de %d), on ne fait rien\n", neighbor.Name, neighbor.Distance)
				}
			}
		}

		fmt.Printf("Queue actuelle : ")
		for i := queue.head; i != queue.tail; i = (i + 1) % MAX_CUSTOM_QUEUE_SIZE {
			item := queue.items[i]
			vertex := item.Content.(*sh.Vertex)
			fmt.Printf("%s (distance de %d) ", vertex.Name, vertex.Distance)
		}
		fmt.Println()
	}

	fmt.Println("Distances finales :")
	for _, vertex := range vertices {
		fmt.Printf("Sommet %s : distance = %d\n", vertex.Name, vertex.Distance)
	}
	fmt.Println("Parcours BFS terminé")

}

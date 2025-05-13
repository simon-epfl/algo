package main

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

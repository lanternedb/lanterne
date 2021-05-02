package model

import "math"

type LoadQuery struct {
	Seed      Vertex
	Degree    int
	MinWeight float32
	MaxWeight float32
}

func NeighborQuery(seed Vertex, degree int) LoadQuery {
	return LoadQuery{
		Seed: seed,
		Degree: degree,
		MinWeight: -math.MaxFloat32,
		MaxWeight: math.MaxFloat32,
	}
}

func AdjacentQuery(seed Vertex) LoadQuery {
	return NeighborQuery(seed, 1)
}

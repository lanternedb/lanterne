package model

type NeighborQuery struct {
	Seed      Vertex
	Degree    int
	MinWeight float32
	MaxWeight float32
}


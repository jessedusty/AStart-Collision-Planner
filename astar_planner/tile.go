package astar_planner

import (
	"github.com/beefsack/go-astar"
	"math"
)

type Tile struct {
	Dims  int
	Coord []int
	Value float64
	Grid  *TileGrid
}

func (t *Tile) EuclideanDistance(other *Tile) float64 {
	if other == nil {
		return math.Inf(1)
	}
	sum := float64(0)
	for i, v := range t.Coord {
		sum += math.Pow(float64(other.Coord[i])-float64(v), 2)
	}
	return math.Sqrt(sum)
}

func (t *Tile) ManhattanDistance(other *Tile) float64 {
	if other == nil {
		return math.Inf(1)
	}
	sum := float64(0)
	for i, v := range t.Coord {
		sum += math.Abs(float64(other.Coord[i]) - float64(v))
	}

	return sum
}

func (t *Tile) PathNeighbors() []astar.Pather {
	var retVal []astar.Pather
	for i := 0; i < t.Dims; i++ {
		curr := CloneCoord(t.Coord)
		curr[i] += 1
		node := t.Grid.Get(curr)
		if node != nil {
			retVal = append(retVal, node)
		}

		curr = CloneCoord(t.Coord)
		curr[i] -= 1
		node = t.Grid.Get(curr)
		if node != nil {
			retVal = append(retVal, node)
		}
	}
	return retVal
}

func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	return to.(*Tile).Value
}

func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	return t.EuclideanDistance(to.(*Tile))
}

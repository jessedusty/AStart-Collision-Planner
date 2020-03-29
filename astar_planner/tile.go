package astar_planner

import (
	"fmt"
	"github.com/beefsack/go-astar"
	"math"
	"strconv"
	"strings"
)

type Tile struct {
	Dims         int
	Coord        []int
	Value        float64
	Grid         *TileGrid
	onlyForwards bool
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
		if node != nil && node.Value < 50 {
			retVal = append(retVal, node)
		}
		if !t.onlyForwards {
			curr = CloneCoord(t.Coord)
			curr[i] -= 1
			node = t.Grid.Get(curr)
			if node != nil && node.Value < 50 {
				retVal = append(retVal, node)
			}
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

func (t *Tile) String() string {
	coords := make([]string, len(t.Coord)+1)
	for i := range t.Coord {
		coords[i] = strconv.Itoa(t.Coord[i])
	}

	coords[len(coords)-1] = fmt.Sprintf("%g", t.Value)

	return "(" + strings.Join(coords, ", ") + ")"
}

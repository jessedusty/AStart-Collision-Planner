package astar_planner

import (
	"fmt"
	"strings"
)

func CoordTo1D(extents []int, coord []int) int {
	if len(extents) != len(coord) {
		panic("Coord and extents are not the same length")
	}

	index := coord[0]
	for n := 1; n < len(extents); n++ {
		index = index*extents[n] + coord[n]
	}

	return index
}

func InsideExtents(extents []int, coord []int) bool {
	for i := range coord {
		if coord[i] >= extents[i] || coord[i] < 0 {
			return false
		}
	}
	return true
}

func CloneCoord(coord []int) []int {
	newCoord := make([]int, len(coord))
	copy(newCoord, coord)
	return newCoord
}

func ConvertTileSliceToIntSlice(tiles []*Tile) [][]int {
	retVal := make([][]int, len(tiles))
	for i, v := range tiles {
		retVal[i] = v.Coord
	}
	return retVal
}

func TilePathAsString(tiles []*Tile) string {
	tileStr := make([]string, len(tiles))
	for i := range tiles {
		tileStr[i] = tiles[i].String()
	}
	return fmt.Sprintf("[%s]", strings.Join(tileStr, ", "))
}

func intSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

package astar_planner

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

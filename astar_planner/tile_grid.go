package astar_planner

import (
	"ee631_midterm/msgs/nav_msgs"
	"github.com/beefsack/go-astar"
)

type TileGrid struct {
	Provider  Provider
	Extents   []int
	TileCache []*Tile
}

func NewOccupancyTileGrid(grid *nav_msgs.OccupancyGrid) *TileGrid {
	op := OccupancyProvider{SourceMap: grid}
	return NewTileGrid(&op)
}

func NewTileGrid(p Provider) *TileGrid {
	tilesNeeded := 1
	for _, i := range p.GetExtents() {
		tilesNeeded *= i
	}

	retVal := TileGrid{
		Provider:  p,
		Extents:   p.GetExtents(),
		TileCache: make([]*Tile, tilesNeeded),
	}
	return &retVal
}

func (tg *TileGrid) Get(coord []int) *Tile {
	if !InsideExtents(tg.Extents, coord) {
		return nil
	}
	index := CoordTo1D(tg.Extents, coord)
	if tg.TileCache[index] == nil {
		tg.TileCache[index] = &Tile{
			Dims:  len(tg.Extents),
			Coord: coord,
			Value: tg.Provider.Get(coord),
			Grid:  tg,
		}
	}
	return tg.TileCache[index]
}

func (tg *TileGrid) Plan(start []int, end []int) (path []*Tile, distance float64, found bool) {
	p, distance, found := astar.Path(tg.Get(start), tg.Get(end))
	path = make([]*Tile, len(p))
	for i, v := range p {
		path[i] = v.(*Tile)
	}
	return
}

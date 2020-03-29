package astar_planner

import (
	"ee631_midterm/msgs/nav_msgs"
	"github.com/beefsack/go-astar"
	"github.com/disintegration/imaging"
	"golang.org/x/image/colornames"
	"image"
)

type TileGrid struct {
	Provider  Provider
	Extents   []int
	TileCache []*Tile
	oneWay    bool
	sparse    bool
}

func NewOccupancyTileGrid(grid *nav_msgs.OccupancyGrid) *TileGrid {
	op := NewOccupancyProvider(grid)
	return NewTileGrid(op, false, false)
}

func NewTileGrid(p Provider, oneWay bool, sparse bool) *TileGrid {
	tilesNeeded := 1
	for _, i := range p.GetExtents() {
		tilesNeeded *= i
	}
	var tileCache []*Tile
	if !sparse {
		tileCache = make([]*Tile, tilesNeeded)
	}

	retVal := TileGrid{
		Provider:  p,
		Extents:   p.GetExtents(),
		TileCache: tileCache,
		oneWay:    oneWay,
		sparse:    sparse,
	}
	return &retVal
}

func (tg *TileGrid) Get(coord []int) *Tile {
	if !InsideExtents(tg.Extents, coord) {
		return nil
	}

	if tg.sparse {
		// Look for existing tile
		for i := range tg.TileCache {
			if intSliceEqual(tg.TileCache[i].Coord, coord) {
				return tg.TileCache[i]
			}
		}

		// Create new tile
		value := tg.Provider.Get(coord)
		newTile := &Tile{
			Dims:  len(tg.Extents),
			Coord: coord,
			Value: value,
			Grid:  tg,
		}
		tg.TileCache = append(tg.TileCache, newTile)
		return newTile
	} else {
		index := CoordTo1D(tg.Extents, coord)
		if tg.TileCache[index] == nil {

			value := tg.Provider.Get(coord)
			tg.TileCache[index] = &Tile{
				Dims:  len(tg.Extents),
				Coord: coord,
				Value: value,
				Grid:  tg,
			}
		}
		return tg.TileCache[index]
	}
}

func (tg *TileGrid) Plan(start []int, end []int) (path []*Tile, distance float64, found bool) {
	p, distance, found := astar.Path(tg.Get(start), tg.Get(end))
	path = make([]*Tile, len(p))
	// Reverse the plan (not quite sure why it's backwards)
	for i := range p {
		path[i] = p[len(p)-1-i].(*Tile)
	}
	return
}

func (tg *TileGrid) TwoDimensionalDebugOutput() {
	if len(tg.Extents) != 2 {
		return
	}

	highCost := colornames.Red
	lowCost := colornames.Black
	image := image.NewRGBA(image.Rect(0, 0, tg.Extents[0], tg.Extents[1]))

	for _, tile := range tg.TileCache {
		if tile == nil {
			continue
		}
		if tile.Value > 10 {
			image.Set(tile.Coord[0], tile.Coord[1], highCost)
		} else {
			image.Set(tile.Coord[0], tile.Coord[1], lowCost)
		}
	}

	imaging.Save(image, "debugOutput.png")
}

func (tg *TileGrid) CostmapOutput() {
	if len(tg.Extents) != 2 {
		return
	}

	image := image.NewRGBA(image.Rect(0, 0, tg.Extents[0], tg.Extents[1]))

	for i := 0; i < tg.Extents[0]; i++ {
		for j := 0; j < tg.Extents[1]; j++ {
			tile := tg.Get([]int{i, j})
			if tile.Value > 10 {
				image.Set(tile.Coord[0], tile.Coord[1], colornames.Black)
			} else {
				image.Set(tile.Coord[0], tile.Coord[1], colornames.White)
			}

		}

	}

	imaging.Save(image, "costmap.png")
}

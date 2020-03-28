package astar_planner

import (
	"ee631_midterm/msgs/nav_msgs"
	"image"
	"math"
)

type Provider interface {
	Get(coord []int) float64
	GetExtents() []int
}

type OccupancyProvider struct {
	SourceMap *nav_msgs.OccupancyGrid
}

func (o *OccupancyProvider) Get(coord []int) float64 {
	return float64(o.SourceMap.Data[CoordTo1D(o.GetExtents(), coord)])
}

func (o *OccupancyProvider) GetExtents() []int {
	return []int{int(o.SourceMap.Info.Width), int(o.SourceMap.Info.Height)}
}

type ImageProvider struct {
	Image image.Image
}

func (i ImageProvider) Get(coord []int) float64 {
	r, _, _, _ := i.Image.At(coord[0], coord[1]).RGBA()
	retVal := float64(r)
	if retVal < 128 {
		retVal = math.Inf(1)
	}
	return retVal
}

func (i ImageProvider) GetExtents() []int {
	return []int{i.Image.Bounds().Max.X, i.Image.Bounds().Max.Y}
}

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
	image     *image.Gray
}

func NewOccupancyProvider(sourceMap *nav_msgs.OccupancyGrid) Provider {
	img := image.NewGray(image.Rect(0, 0, int(sourceMap.Info.Width), int(sourceMap.Info.Height)))
	for i := range sourceMap.Data {
		img.Pix[i] = uint8(sourceMap.Data[i])
	}
	//imaging.Save(img, "costmap_out.png")
	return &InvertedGreyImageProvider{Image: img}
}

func (o *OccupancyProvider) Get(coord []int) float64 {
	value := o.SourceMap.Data[CoordTo1D(o.GetExtents(), coord)]
	return float64(value)
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

type GreyImageProvider struct {
	Image image.Image
}

func (i GreyImageProvider) Get(coord []int) float64 {
	img := i.Image.(*image.Gray)
	r := img.GrayAt(coord[0], coord[1]).Y
	retVal := float64(r)
	if retVal < 128 {
		retVal = math.Inf(1)
	} else {
		retVal = 0
	}
	return retVal
}

func (i GreyImageProvider) GetExtents() []int {
	return []int{i.Image.Bounds().Max.X, i.Image.Bounds().Max.Y}
}

type InvertedGreyImageProvider struct {
	Image image.Image
}

func (i InvertedGreyImageProvider) Get(coord []int) float64 {
	img := i.Image.(*image.Gray)
	r := img.GrayAt(coord[0], coord[1]).Y
	retVal := float64(r)
	if retVal > 50 {
		retVal = math.Inf(1)
	} else {
		retVal = 0
	}
	return retVal
}

func (i InvertedGreyImageProvider) GetExtents() []int {
	return []int{i.Image.Bounds().Max.X, i.Image.Bounds().Max.Y}
}

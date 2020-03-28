package grid_planner

import (
	"ee631_midterm/dstarlite/grid"
	"log"
	"math"
)

const safetyMargin = 10

type RobotPlan struct {
	XYPath         []grid.Coord
	pointDistances []float64
	length         float64
}

func (rp *RobotPlan) ReCalc() {
	rp.length = 0
	if len(rp.XYPath) <= 1 {
		return
	}

	rp.pointDistances = make([]float64, len(rp.XYPath))

	prev := rp.XYPath[0]
	for i, c := range rp.XYPath {
		rp.length += prev.Dist(c)
		rp.pointDistances[i] = rp.length
		prev = c
	}

	return
}

func (rp *RobotPlan) InterpolateFromDistance(dist float64) *grid.Coord {
	if dist < 0 {
		return nil
	}

	for i, n := range rp.pointDistances {
		if n > dist {
			return &rp.XYPath[i]
		}
	}

	return &rp.XYPath[len(rp.XYPath)-1]
}

func (rp *RobotPlan) CollidesWith(otherPlan *RobotPlan) (collisions []bool) {
	log.Printf("CollidesWith() Start")
	collisions = make([]bool, int(math.Ceil(rp.length)))
	for i, _ := range collisions {
		me := rp.InterpolateFromDistance(float64(i))
		otherGuy := otherPlan.InterpolateFromDistance(float64(i))
		collisions[i] = me.Dist(*otherGuy) < safetyMargin
	}
	log.Printf("CollidesWith() Finish")
	return
}

func NewRobotPlan(path []grid.Coord) *RobotPlan {
	retVal := RobotPlan{
		XYPath: path,
	}
	retVal.ReCalc()
	return &retVal
}

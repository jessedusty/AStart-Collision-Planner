package grid_planner

import (
	"ee631_midterm/astar_planner"
	"log"
	"math"
)

const safetyMargin = 10

type RobotPlan struct {
	XYPath []*astar_planner.Tile
	// Possibly remove these?
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
		rp.length += prev.EuclideanDistance(c)
		rp.pointDistances[i] = rp.length
		prev = c
	}

	return
}

func (rp *RobotPlan) InterpolateFromDistance(dist float64) *astar_planner.Tile {
	if dist < 0 {
		return nil
	}

	for i, n := range rp.pointDistances {
		if n > dist {
			return rp.XYPath[i]
		}
	}

	return rp.XYPath[len(rp.XYPath)-1]
}

func (rp *RobotPlan) FindIndexForDistance(dist float64) int {
	if dist < 0 {
		return 0
	}

	for i, n := range rp.pointDistances {
		if n > dist {
			return i
		}
	}

	return len(rp.XYPath) - 1
}

//func (rp *RobotPlan) CollidesWithAtDistance(otherPlan *RobotPlan, distance int) bool {
//	if distance > len(rp.XYPath) || distance > len(otherPlan.XYPath) {
//		return false
//	}
//
//	closeTo := rp.XYPath[distance].EuclideanDistance(otherPlan.XYPath[distance])
//	return closeTo < safetyMargin
//}

func (rp *RobotPlan) CollidesWith(otherPlan *RobotPlan) (collisions []bool) {
	log.Printf("CollidesWith() Start")
	collisions = make([]bool, int(math.Ceil(rp.length)))
	for i, _ := range collisions {
		me := rp.InterpolateFromDistance(float64(i))
		otherGuy := otherPlan.InterpolateFromDistance(float64(i))
		collisions[i] = me.EuclideanDistance(otherGuy) < safetyMargin
	}
	log.Printf("CollidesWith() Finish")
	return
}

func NewRobotPlan(path []*astar_planner.Tile) *RobotPlan {
	retVal := RobotPlan{
		XYPath: path,
	}
	retVal.ReCalc()
	return &retVal
}

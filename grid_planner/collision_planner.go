package grid_planner

import (
	"ee631_midterm/dstarlite/grid"
	"ee631_midterm/msgs/nav_msgs"
	"fmt"
	"log"
	"math"
	"sync"
	"time"
)

type CollisionPlanner struct {
	Num               int
	grid              *grid.Data
	robots            []*RobotPlan
	collisionPlan     []grid.Coord
	RobotPositionChan chan NewPos
	Planners          []BasicPlanner
}

func (cp *CollisionPlanner) InputLoop(pathInputs []chan []grid.Coord) {
	for {
		var wg sync.WaitGroup
		for i := 0; i < cp.Num; i++ {
			wg.Add(1)
			go func(i int) {
				log.Printf("CollisionPlanner::InputLoop() Waiting for plan %d", i)
				plan := <-pathInputs[i]
				log.Printf("CollisionPlanner::InputLoop() Got plan %d", i)
				cp.robots[i] = NewRobotPlan(plan)
				log.Printf("CollisionPlanner::InputLoop() Updating plan %d", i)
				wg.Done()
			}(i)
		}
		wg.Wait()
		log.Printf("CollisionPlanner::InputLoop(): Got all plans creating collision grid")
		cp.CreateCollisionGrid()
		fmt.Printf("%v", cp.grid.GetMapData())

		log.Printf("CollisionPlanner::InputLoop(): Done creating collision grid")
		log.Printf("CollisionPlanner::InputLoop(): Planning")
		cp.Plan()
		log.Printf("CollisionPlanner::InputLoop(): Done Planning %v", cp.collisionPlan)
		cp.PlayPlan()
	}
}

func (cp *CollisionPlanner) PlayPlan() {

	paths := make([]*nav_msgs.Path, cp.Num)
	for i, planner := range cp.Planners {
		paths[i] = planner.getPath()
	}

	for _, pathPosition := range cp.collisionPlan {
		for robotID, _ := range pathPosition {
			pathIndex := cp.robots[robotID].FindIndexForDistance(float64(pathPosition[robotID]))
			cp.RobotPositionChan <- NewPos{
				Coord:   paths[robotID].Poses[pathIndex],
				RobotId: robotID,
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func NewCollisionPlanner(numRobots int, robotPositionChan chan NewPos, planners []BasicPlanner) *CollisionPlanner {
	// TODO: Input robot plan
	retVal := CollisionPlanner{Num: numRobots, robots: make([]*RobotPlan, numRobots), RobotPositionChan: robotPositionChan, Planners: planners}
	return &retVal
}

func initCoordWithValue(dims, defaultValue int) grid.Coord {
	retVal := make(grid.Coord, dims)
	for i, _ := range retVal {
		retVal[i] = defaultValue
	}
	return retVal
}

func (cp *CollisionPlanner) CreateCollisionGrid() {
	maxDimensions := make([]int, len(cp.robots))
	for i, robot := range cp.robots {
		robot.ReCalc()
		maxDimensions[i] = int(math.Ceil(robot.length))
	}

	cp.grid = grid.NewGrid(maxDimensions, initCoordWithValue(cp.Num, 0), maxDimensions, true)

	// TODO: Test me
	for i, myRobot := range cp.robots {
		for j, otherRobot := range cp.robots {
			if i == j { // Same robot condition
				continue
			}
			log.Printf("CollisionPlanner::CreateCollisionGrid(): Finding all collisions between %d and %d", i, j)
			mask := initCoordWithValue(cp.Num, 1)
			mask[i] = 0
			mask[j] = 0
			myRobotCollisions := myRobot.CollidesWith(otherRobot)
			for distIndex, collides := range myRobotCollisions {
				if distIndex > int(myRobot.length) || distIndex > int(otherRobot.length) {
					log.Println("CollisionPlanner::CreateCollisionGrid(): break early")
					break
				}
				if collides {
					log.Printf("CollisionPlanner::CreateCollisionGrid(): Found a collision at %d", distIndex)
					fillStart := initCoordWithValue(cp.Num, 0)
					fillStart[i] = distIndex
					fillStart[j] = distIndex
					cp.grid.Fill(-1, fillStart, mask)
				}
			}
		}
	}
	log.Printf("CollisionPlanner::CreateCollisionGrid(): Finished creating collision grid")
}

func (cp *CollisionPlanner) Plan() {
	cp.collisionPlan = cp.grid.Plan()
}

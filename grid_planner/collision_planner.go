package grid_planner

import (
	"ee631_midterm/astar_planner"
	"ee631_midterm/msgs/nav_msgs"
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"sync"
	"time"
)

type CollisionPlanner struct {
	Num               int
	grid              *astar_planner.TileGrid
	robots            []*RobotPlan
	collisionPlan     []*astar_planner.Tile
	RobotPositionChan chan NewPos
	Planners          []BasicPlanner
}

func (cp *CollisionPlanner) InputLoop(pathInputs []chan []*astar_planner.Tile) {
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

		log.Printf("CollisionPlanner::InputLoop(): Done creating collision grid")
		log.Printf("CollisionPlanner::InputLoop(): Planning")
		cp.Plan()
		log.Printf("CollisionPlanner::InputLoop(): Done Planning")
		cp.PlayPlan()
	}
}

func (cp *CollisionPlanner) PlayPlan() {

	paths := make([]*nav_msgs.Path, cp.Num)
	for i, planner := range cp.Planners {
		paths[i] = planner.getPath()
	}

	for _, pathTile := range cp.collisionPlan {
		for robotID := range pathTile.Coord {
			pathIndex := cp.robots[robotID].FindIndexForDistance(float64(pathTile.Coord[robotID]))
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

func initCoordWithValue(dims, defaultValue int) []int {
	retVal := make([]int, dims)
	for i, _ := range retVal {
		retVal[i] = defaultValue
	}
	return retVal
}

func (cp *CollisionPlanner) Get(coord []int) float64 {
	//start := time.Now()
	collisionCost := math.Inf(1)
	freeCost := float64(0)

	current := astar_planner.Tile{Coord: coord}
	end := astar_planner.Tile{Coord: cp.GetExtents()}
	if current.ManhattanDistance(&end) < 2 {
		return freeCost
	}

	for i := range coord {
		sourceRobotPoint := cp.robots[i].XYPath[coord[i]]
		for j := range coord {
			if i == j {
				continue
			}

			// If I'm finished no cost
			if cp.GetExtents()[i]-coord[i] < 2 {
				continue
			}

			// If other is finished no cost
			if cp.GetExtents()[j]-coord[j] < 2 {
				continue
			}
			otherRobotPoint := cp.robots[j].XYPath[coord[j]]
			if sourceRobotPoint.EuclideanDistance(otherRobotPoint) < safetyMargin {
				//log.Printf("Collision of %d and %d at coord %v in %d us", i, j, coord, time.Since(start).Nanoseconds())
				return collisionCost
			}
		}
	}

	return freeCost
}

func (cp *CollisionPlanner) GetExtents() []int {
	retVal := make([]int, len(cp.robots))
	for i := range cp.robots {
		retVal[i] = len(cp.robots[i].XYPath)
	}
	return retVal
}

func (cp *CollisionPlanner) CreateCollisionGrid() {
	maxDimensions := make([]int, len(cp.robots))
	for i, robot := range cp.robots {
		robot.ReCalc()
		maxDimensions[i] = int(math.Ceil(robot.length))
	}

	cp.grid = astar_planner.NewTileGrid(cp, true, true)
}

func (cp *CollisionPlanner) Plan() {
	goal := cp.GetExtents()
	for i := range goal {
		goal[i] -= 1
	}
	plan, _, found := cp.grid.Plan(initCoordWithValue(cp.Num, 0), goal)
	if !found {
		log.Print("No collision path found")
	}
	//cp.grid.CostmapOutput()
	cp.grid.TwoDimensionalDebugOutput()
	pathSaver(plan)
	cp.collisionPlan = plan
}

func pathSaver(path []*astar_planner.Tile) {
	ints := astar_planner.ConvertTileSliceToIntSlice(path)
	output, _ := json.Marshal(ints)
	err := ioutil.WriteFile("collisionpath.json", output, 0755)
	if err != nil {
		panic(err)
	}
}

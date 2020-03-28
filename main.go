package main

//go:generate gengo msg std_msgs/String
import (
	"ee631_midterm/dstarlite/grid"
	"ee631_midterm/grid_planner"
	"ee631_midterm/msgs/geometry_msgs"
	"ee631_midterm/msgs/map_msgs"
	"ee631_midterm/msgs/nav_msgs"
	"fmt"
	"github.com/fetchrobotics/rosgo/ros"
	"github.com/pkg/profile"
	"os"
)

const (
	NUM_ROBOTS = 2
)

func mapChannelSplitter(input chan *nav_msgs.OccupancyGrid, outputs []chan *nav_msgs.OccupancyGrid) {
	for in := range input {
		for i, _ := range outputs {
			outputs[i] <- in
		}
	}
}

func goalChannelSplitter(input chan *geometry_msgs.PoseStamped, outputs []chan *geometry_msgs.PoseStamped) {
	for in := range input {
		for i, _ := range outputs {
			outputs[i] <- in
		}
	}
}

func pathChannelSplitter(input chan *nav_msgs.Path, outputs []chan *nav_msgs.Path) {
	for in := range input {
		for i, _ := range outputs {
			outputs[i] <- in
		}
	}
}

func main() {
	defer profile.Start(profile.ProfilePath(".")).Stop()

	node, err := ros.NewNode("/gridmapper", os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer node.Shutdown()

	// Publishes new robot positions
	positionChan := make(chan grid_planner.NewPos)
	go grid_planner.RobotPositionUpdateLoop(positionChan, NUM_ROBOTS, node)

	goalChan := make(chan *geometry_msgs.PoseStamped)
	mapChan := make(chan *nav_msgs.OccupancyGrid)
	plannerMapInputs := make([]chan *nav_msgs.OccupancyGrid, NUM_ROBOTS)
	goalChanInputs := make([]chan *geometry_msgs.PoseStamped, NUM_ROBOTS)
	level1Planners := make([]grid_planner.BasicPlanner, NUM_ROBOTS)
	gridPathChanOutputs := make([]chan []grid.Coord, NUM_ROBOTS)

	startPositions := []grid.Coord{
		grid.Coord{21, 177},
		grid.Coord{21, 177 - 25},
		grid.Coord{21, 177 - 50},
	}

	for robotID := 0; robotID < NUM_ROBOTS; robotID++ {
		level1Planners[robotID] = grid_planner.BasicPlanner{
			Start:             startPositions[robotID],
			Goal:              grid.Coord{82, 121},
			GoalChan:          make(chan *geometry_msgs.PoseStamped),
			MapChan:           make(chan *nav_msgs.OccupancyGrid),
			MapUpdateChan:     make(chan *map_msgs.OccupancyGridUpdate),
			PathChan:          make(chan *nav_msgs.Path),
			GridPathChan:      make(chan []grid.Coord),
			RobotPositionChan: positionChan,
			RobotId:           robotID,
		}

		plannerMapInputs[robotID] = level1Planners[robotID].MapChan
		goalChanInputs[robotID] = level1Planners[robotID].GoalChan

		go level1Planners[robotID].PlanGenerator()
		go level1Planners[robotID].PathPublisher(node, level1Planners[robotID].PathChan)

		gridPathChanOutputs[robotID] = level1Planners[robotID].GridPathChan
	}

	go mapChannelSplitter(mapChan, plannerMapInputs)
	go goalChannelSplitter(goalChan, goalChanInputs)

	collisionPlanner := grid_planner.NewCollisionPlanner(NUM_ROBOTS, positionChan, level1Planners)
	go collisionPlanner.InputLoop(gridPathChanOutputs)

	node.NewSubscriber("/move_base_simple/goal", geometry_msgs.MsgPoseStamped, func(msg *geometry_msgs.PoseStamped) {
		goalChan <- msg
	})

	node.NewSubscriber("/global_costmap/costmap/costmap", nav_msgs.MsgOccupancyGrid, func(msg *nav_msgs.OccupancyGrid) {
		mapChan <- msg
	})

	node.NewSubscriber("/global_costmap/costmap/costmap_updates", map_msgs.MsgOccupancyGridUpdate, func(msg *map_msgs.OccupancyGridUpdate) {
		//p.MapUpdateChan<-msg
	})

	node.Spin()
}

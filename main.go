package main

//go:generate gengo msg std_msgs/String
import (
	"ee631_midterm/astar_planner"
	"ee631_midterm/grid_planner"
	"ee631_midterm/msgs/geometry_msgs"
	"ee631_midterm/msgs/map_msgs"
	"ee631_midterm/msgs/nav_msgs"
	"fmt"
	"github.com/fetchrobotics/rosgo/ros"
	"os"
)

const (
	NUM_ROBOTS = 4
)

func main() {
	//defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()

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
	gridPathChanOutputs := make([]chan []*astar_planner.Tile, NUM_ROBOTS)

	startPositions := [][]int{
		{21, 177},
		{21, 177 - 25},
		{21, 177 - 50},
		{88, 178},
		{88 + 25, 178},
		{88 + 50, 178},
		{88 + 75, 178},
	}

	for robotID := 0; robotID < NUM_ROBOTS; robotID++ {
		level1Planners[robotID] = grid_planner.BasicPlanner{
			Grid:              &astar_planner.TileGrid{},
			Start:             startPositions[robotID],
			Goal:              []int{82, 121},
			GoalChan:          make(chan *geometry_msgs.PoseStamped),
			MapChan:           make(chan *nav_msgs.OccupancyGrid),
			MapUpdateChan:     make(chan *map_msgs.OccupancyGridUpdate),
			PathChan:          make(chan *nav_msgs.Path),
			GridPathChan:      make(chan []*astar_planner.Tile),
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

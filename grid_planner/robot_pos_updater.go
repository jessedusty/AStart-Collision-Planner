package grid_planner

import (
	"ee631_midterm/msgs/geometry_msgs"
	"fmt"
	"github.com/fetchrobotics/rosgo/ros"
)

type NewPos struct {
	Coord   geometry_msgs.PoseStamped
	RobotId int
}

func RobotPositionUpdateLoop(newPos chan NewPos, numRobots int, node ros.Node) {
	pubs := make([]ros.Publisher, numRobots)
	for i, _ := range pubs {
		poseTopic := fmt.Sprintf("/robot%d/pose", i)
		pubs[i] = node.NewPublisher(poseTopic, geometry_msgs.MsgPoseStamped)
	}

	for pos := range newPos {
		pubs[pos.RobotId].Publish(&pos.Coord)
	}
}

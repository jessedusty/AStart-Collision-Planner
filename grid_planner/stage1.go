package grid_planner

import (
	"ee631_midterm/dstarlite/grid"
	"ee631_midterm/msgs/geometry_msgs"
	"ee631_midterm/msgs/map_msgs"
	"ee631_midterm/msgs/nav_msgs"
	"ee631_midterm/msgs/std_msgs"
	"fmt"
	"github.com/fetchrobotics/rosgo/ros"
	"log"
	"time"
)

type BasicPlanner struct {
	lastMap           *nav_msgs.OccupancyGrid
	Start, Goal       grid.Coord
	Grid              *grid.Data
	GoalChan          chan *geometry_msgs.PoseStamped
	MapChan           chan *nav_msgs.OccupancyGrid
	MapUpdateChan     chan *map_msgs.OccupancyGridUpdate
	PathChan          chan *nav_msgs.Path
	GridPathChan      chan []grid.Coord
	RobotPositionChan chan NewPos
	Path              []grid.Coord
	MapResolution     float64
	MapOrigin         []float64
	RobotId           int
}

func (p *BasicPlanner) PlanGenerator() {
	for {
		select {
		case newMap := <-p.MapChan:
			log.Printf(debugStr, p.RobotId, "PathGenerator", "Got new map")
			p.newMap(newMap)
			p.rePlan()
			log.Printf(debugStr, p.RobotId, "PathGenerator", "Push PathChan")
			p.PathChan <- p.getPath()
			log.Printf(debugStr, p.RobotId, "PathGenerator", "Push GridPathChan")
			p.GridPathChan <- p.Path
			log.Printf(debugStr, p.RobotId, "PathGenerator", "Finished processing new map")
			//go p.FollowPath()
		case mapUpdate := <-p.MapUpdateChan:
			log.Printf(debugStr, p.RobotId, "PathGenerator", "Got map update")
			p.mapUpdate(mapUpdate)
			log.Printf(debugStr, p.RobotId, "PathGenerator", "Finished processing map update")
		case goalUpdate := <-p.GoalChan:
			log.Printf(debugStr, p.RobotId, "PathGenerator", "Got goal update")
			p.Goal = p.ConvertPointToCoord(goalUpdate.Pose.Position)
			log.Printf("Got new goal (%d, %d)", p.Goal[0], p.Goal[1])
			p.updateGoal()
			p.rePlan()
			log.Printf(debugStr, p.RobotId, "PathGenerator", "Push PathChan")
			p.PathChan <- p.getPath()
			log.Printf(debugStr, p.RobotId, "PathGenerator", "Push GridPathChan")
			p.GridPathChan <- p.Path
			log.Printf(debugStr, p.RobotId, "PathGenerator", "Finished processing goal update")
			//go p.FollowPath()
		}
	}
}

func (p *BasicPlanner) PathPublisher(n ros.Node, pathChan chan *nav_msgs.Path) {

	pathTopic := fmt.Sprintf("/robot%d/path", p.RobotId)
	pub := n.NewPublisher(pathTopic, nav_msgs.MsgPath)

	for path := range pathChan {
		log.Printf("ROBOT %d ::: PathPublisher(): Publishing path with length %d", p.RobotId, len(path.Poses))
		pub.Publish(path)
		log.Printf("ROBOT %d ::: PathPublisher(): Finished publishing path", p.RobotId)

	}
}

func (p *BasicPlanner) FollowPath() {
	path := p.getPath()
	for _, pose := range path.Poses {
		p.RobotPositionChan <- NewPos{Coord: pose, RobotId: p.RobotId}
		time.Sleep(time.Millisecond * 10)
		//log.Printf("ROBOT %d ::: Following path, point %d", p.RobotId, i)
	}
}

//const (
//	Costmap2dLethalObstacle            = 254
//	Costmap2dNoInformation             = 255
//	Costmap2dInscribedInflatedObstacle = 253
//	Costmap2dFreeSpace                 = 0
//)

const (
	Costmap2dLethalObstacle            = 50
	Costmap2dNoInformation             = 50
	Costmap2dInscribedInflatedObstacle = 50
	Costmap2dFreeSpace                 = 0
)

func (p *BasicPlanner) mapUpdate(m *map_msgs.OccupancyGridUpdate) {
	// TODO: Add me in
	log.Println("TODO: Map updating not yet supported")
}

const debugStr = "ROBOT %d ::: %s(): %s"

func (p *BasicPlanner) rePlan() {
	log.Printf(debugStr, p.RobotId, "rePlan", "Start")
	p.Path = p.Grid.Plan()
	if p.Path == nil {
		log.Printf(debugStr, p.RobotId, "rePlan", "Path not found")
	} else {
		log.Printf(debugStr, p.RobotId, "rePlan", "Finish")
	}
}

func (p *BasicPlanner) updateGoal() {
	log.Printf(debugStr, p.RobotId, "updateGoal", "Start")
	p.Grid = grid.NewGrid([]int{int(p.lastMap.Info.Width), int(p.lastMap.Info.Height)}, p.Start, p.Goal, true)
	for row := 0; row < int(p.lastMap.Info.Height); row++ {
		for col := 0; col < int(p.lastMap.Info.Width); col++ {
			value := p.lastMap.Data[row*int(p.lastMap.Info.Width)+col]
			if int(value) > Costmap2dInscribedInflatedObstacle {
				p.Grid.Set(grid.Coord{col, row}, -1)
			}
		}
	}
	log.Printf(debugStr, p.RobotId, "updateGoal", "Finish")
}

func (p *BasicPlanner) newMap(m *nav_msgs.OccupancyGrid) {
	//p.Start = grid.Coord{1, 1} // FIXME: Actually use start and end
	//p.Goal = grid.Coord{100, 100}
	p.MapResolution = float64(m.Info.Resolution)
	p.MapOrigin = []float64{m.Info.Origin.Position.X, m.Info.Origin.Position.Y}
	p.lastMap = m
	p.updateGoal()
}

func (p *BasicPlanner) getPath() *nav_msgs.Path {

	if p.Path == nil {
		return nil
	}

	retVal := nav_msgs.Path{
		Header: std_msgs.Header{
			Stamp:   ros.Now(),
			FrameId: "map",
		},
		Poses: make([]geometry_msgs.PoseStamped, len(p.Path)),
	}

	for i, point := range p.Path {
		retVal.Poses[i] = geometry_msgs.PoseStamped{
			Header: std_msgs.Header{
				Seq:     uint32(i),
				Stamp:   ros.Now(),
				FrameId: "/map",
			},
			Pose: geometry_msgs.Pose{
				Position: geometry_msgs.Point{
					X: float64(point[0])*p.MapResolution + p.MapOrigin[0],
					Y: float64(point[1])*p.MapResolution + p.MapOrigin[1],
					Z: 0,
				},
				Orientation: geometry_msgs.Quaternion{},
			},
		}
	}

	return &retVal
}

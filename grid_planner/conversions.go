package grid_planner

import (
	"ee631_midterm/dstarlite/grid"
	"ee631_midterm/msgs/geometry_msgs"
)

func (bp BasicPlanner) ConvertPointToCoord(point geometry_msgs.Point) grid.Coord {
	return grid.Coord{
		int((point.X - bp.MapOrigin[0]) / bp.MapResolution),
		int((point.Y - bp.MapOrigin[1]) / bp.MapResolution),
	}
}

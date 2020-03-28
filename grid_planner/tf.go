package grid_planner

import (
	"ee631_midterm/msgs/geometry_msgs"
	"ee631_midterm/msgs/std_msgs"
	"ee631_midterm/msgs/tf2_msgs"
	"github.com/fetchrobotics/rosgo/ros"
)

func TFMessageForPose(stamped *geometry_msgs.PoseStamped) tf2_msgs.TFMessage {
	return tf2_msgs.TFMessage{
		Transforms: []geometry_msgs.TransformStamped{
			{Header: std_msgs.Header{
				Seq:     0,
				Stamp:   ros.Now(),
				FrameId: "/map",
			},
				ChildFrameId: "/robot", Transform: geometry_msgs.Transform{
					Translation: geometry_msgs.Vector3{
						X: stamped.Pose.Position.X,
						Y: stamped.Pose.Position.Y,
						Z: stamped.Pose.Position.Z,
					},
				}}}}
}

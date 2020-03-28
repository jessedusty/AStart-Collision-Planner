// Automatically generated from the message definition "geometry_msgs/PoseWithCovarianceStamped.msg"
package geometry_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgPoseWithCovarianceStamped struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPoseWithCovarianceStamped) Text() string {
	return t.text
}

func (t *_MsgPoseWithCovarianceStamped) Name() string {
	return t.name
}

func (t *_MsgPoseWithCovarianceStamped) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPoseWithCovarianceStamped) NewMessage() ros.Message {
	m := new(PoseWithCovarianceStamped)
	m.Header = std_msgs.Header{}
	m.Pose = PoseWithCovariance{}
	return m
}

var (
	MsgPoseWithCovarianceStamped = &_MsgPoseWithCovarianceStamped{
		`# This expresses an estimated pose with a reference coordinate frame and timestamp

Header header
PoseWithCovariance pose
`,
		"geometry_msgs/PoseWithCovarianceStamped",
		"953b798c0f514ff060a53a3498ce6246",
	}
)

type PoseWithCovarianceStamped struct {
	Header std_msgs.Header    `rosmsg:"header:Header"`
	Pose   PoseWithCovariance `rosmsg:"pose:PoseWithCovariance"`
}

func (m *PoseWithCovarianceStamped) GetType() ros.MessageType {
	return MsgPoseWithCovarianceStamped
}

func (m *PoseWithCovarianceStamped) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	if err = m.Pose.Serialize(buf); err != nil {
		return err
	}
	return err
}

func (m *PoseWithCovarianceStamped) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = m.Pose.Deserialize(buf); err != nil {
		return err
	}
	return err
}

// Automatically generated from the message definition "geometry_msgs/PoseStamped.msg"
package geometry_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgPoseStamped struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPoseStamped) Text() string {
	return t.text
}

func (t *_MsgPoseStamped) Name() string {
	return t.name
}

func (t *_MsgPoseStamped) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPoseStamped) NewMessage() ros.Message {
	m := new(PoseStamped)
	m.Header = std_msgs.Header{}
	m.Pose = Pose{}
	return m
}

var (
	MsgPoseStamped = &_MsgPoseStamped{
		`# A Pose with reference coordinate frame and timestamp
Header header
Pose pose
`,
		"geometry_msgs/PoseStamped",
		"d3812c3cbc69362b77dc0b19b345f8f5",
	}
)

type PoseStamped struct {
	Header std_msgs.Header `rosmsg:"header:Header"`
	Pose   Pose            `rosmsg:"pose:Pose"`
}

func (m *PoseStamped) GetType() ros.MessageType {
	return MsgPoseStamped
}

func (m *PoseStamped) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	if err = m.Pose.Serialize(buf); err != nil {
		return err
	}
	return err
}

func (m *PoseStamped) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = m.Pose.Deserialize(buf); err != nil {
		return err
	}
	return err
}

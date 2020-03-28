// Automatically generated from the message definition "geometry_msgs/PoseArray.msg"
package geometry_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgPoseArray struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPoseArray) Text() string {
	return t.text
}

func (t *_MsgPoseArray) Name() string {
	return t.name
}

func (t *_MsgPoseArray) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPoseArray) NewMessage() ros.Message {
	m := new(PoseArray)
	m.Header = std_msgs.Header{}
	m.Poses = []Pose{}
	return m
}

var (
	MsgPoseArray = &_MsgPoseArray{
		`# An array of poses with a header for global reference.

Header header

Pose[] poses
`,
		"geometry_msgs/PoseArray",
		"916c28c5764443f268b296bb671b9d97",
	}
)

type PoseArray struct {
	Header std_msgs.Header `rosmsg:"header:Header"`
	Poses  []Pose          `rosmsg:"poses:Pose[]"`
}

func (m *PoseArray) GetType() ros.MessageType {
	return MsgPoseArray
}

func (m *PoseArray) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Poses)))
	for _, e := range m.Poses {
		if err = e.Serialize(buf); err != nil {
			return err
		}
	}
	return err
}

func (m *PoseArray) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Poses = make([]Pose, int(size))
		for i := 0; i < int(size); i++ {
			if err = m.Poses[i].Deserialize(buf); err != nil {
				return err
			}
		}
	}
	return err
}

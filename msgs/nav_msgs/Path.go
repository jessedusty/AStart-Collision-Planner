// Automatically generated from the message definition "nav_msgs/Path.msg"
package nav_msgs

import (
	"bytes"
	"ee631_midterm/msgs/geometry_msgs"
	"ee631_midterm/msgs/std_msgs"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgPath struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPath) Text() string {
	return t.text
}

func (t *_MsgPath) Name() string {
	return t.name
}

func (t *_MsgPath) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPath) NewMessage() ros.Message {
	m := new(Path)
	m.Header = std_msgs.Header{}
	m.Poses = []geometry_msgs.PoseStamped{}
	return m
}

var (
	MsgPath = &_MsgPath{
		`#An array of poses that represents a Path for a robot to follow
Header header
geometry_msgs/PoseStamped[] poses
`,
		"nav_msgs/Path",
		"6227e2b7e9cce15051f669a5e197bbf7",
	}
)

type Path struct {
	Header std_msgs.Header             `rosmsg:"header:Header"`
	Poses  []geometry_msgs.PoseStamped `rosmsg:"poses:PoseStamped[]"`
}

func (m *Path) GetType() ros.MessageType {
	return MsgPath
}

func (m *Path) Serialize(buf *bytes.Buffer) error {
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

func (m *Path) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Poses = make([]geometry_msgs.PoseStamped, int(size))
		for i := 0; i < int(size); i++ {
			if err = m.Poses[i].Deserialize(buf); err != nil {
				return err
			}
		}
	}
	return err
}

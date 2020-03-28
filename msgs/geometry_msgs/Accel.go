// Automatically generated from the message definition "geometry_msgs/Accel.msg"
package geometry_msgs

import (
	"bytes"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgAccel struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgAccel) Text() string {
	return t.text
}

func (t *_MsgAccel) Name() string {
	return t.name
}

func (t *_MsgAccel) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgAccel) NewMessage() ros.Message {
	m := new(Accel)
	m.Linear = Vector3{}
	m.Angular = Vector3{}
	return m
}

var (
	MsgAccel = &_MsgAccel{
		`# This expresses acceleration in free space broken into its linear and angular parts.
Vector3  linear
Vector3  angular
`,
		"geometry_msgs/Accel",
		"9f195f881246fdfa2798d1d3eebca84a",
	}
)

type Accel struct {
	Linear  Vector3 `rosmsg:"linear:Vector3"`
	Angular Vector3 `rosmsg:"angular:Vector3"`
}

func (m *Accel) GetType() ros.MessageType {
	return MsgAccel
}

func (m *Accel) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Linear.Serialize(buf); err != nil {
		return err
	}
	if err = m.Angular.Serialize(buf); err != nil {
		return err
	}
	return err
}

func (m *Accel) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Linear.Deserialize(buf); err != nil {
		return err
	}
	if err = m.Angular.Deserialize(buf); err != nil {
		return err
	}
	return err
}

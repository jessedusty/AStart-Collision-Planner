// Automatically generated from the message definition "geometry_msgs/Wrench.msg"
package geometry_msgs

import (
	"bytes"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgWrench struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgWrench) Text() string {
	return t.text
}

func (t *_MsgWrench) Name() string {
	return t.name
}

func (t *_MsgWrench) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgWrench) NewMessage() ros.Message {
	m := new(Wrench)
	m.Force = Vector3{}
	m.Torque = Vector3{}
	return m
}

var (
	MsgWrench = &_MsgWrench{
		`# This represents force in free space, separated into
# its linear and angular parts.
Vector3  force
Vector3  torque
`,
		"geometry_msgs/Wrench",
		"4f539cf138b23283b520fd271b567936",
	}
)

type Wrench struct {
	Force  Vector3 `rosmsg:"force:Vector3"`
	Torque Vector3 `rosmsg:"torque:Vector3"`
}

func (m *Wrench) GetType() ros.MessageType {
	return MsgWrench
}

func (m *Wrench) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Force.Serialize(buf); err != nil {
		return err
	}
	if err = m.Torque.Serialize(buf); err != nil {
		return err
	}
	return err
}

func (m *Wrench) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Force.Deserialize(buf); err != nil {
		return err
	}
	if err = m.Torque.Deserialize(buf); err != nil {
		return err
	}
	return err
}

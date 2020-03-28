// Automatically generated from the message definition "geometry_msgs/Vector3Stamped.msg"
package geometry_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgVector3Stamped struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgVector3Stamped) Text() string {
	return t.text
}

func (t *_MsgVector3Stamped) Name() string {
	return t.name
}

func (t *_MsgVector3Stamped) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgVector3Stamped) NewMessage() ros.Message {
	m := new(Vector3Stamped)
	m.Header = std_msgs.Header{}
	m.Vector = Vector3{}
	return m
}

var (
	MsgVector3Stamped = &_MsgVector3Stamped{
		`# This represents a Vector3 with reference coordinate frame and timestamp
Header header
Vector3 vector
`,
		"geometry_msgs/Vector3Stamped",
		"7b324c7325e683bf02a9b14b01090ec7",
	}
)

type Vector3Stamped struct {
	Header std_msgs.Header `rosmsg:"header:Header"`
	Vector Vector3         `rosmsg:"vector:Vector3"`
}

func (m *Vector3Stamped) GetType() ros.MessageType {
	return MsgVector3Stamped
}

func (m *Vector3Stamped) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	if err = m.Vector.Serialize(buf); err != nil {
		return err
	}
	return err
}

func (m *Vector3Stamped) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = m.Vector.Deserialize(buf); err != nil {
		return err
	}
	return err
}

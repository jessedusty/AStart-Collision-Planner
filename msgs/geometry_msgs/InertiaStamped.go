// Automatically generated from the message definition "geometry_msgs/InertiaStamped.msg"
package geometry_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgInertiaStamped struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgInertiaStamped) Text() string {
	return t.text
}

func (t *_MsgInertiaStamped) Name() string {
	return t.name
}

func (t *_MsgInertiaStamped) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgInertiaStamped) NewMessage() ros.Message {
	m := new(InertiaStamped)
	m.Header = std_msgs.Header{}
	m.Inertia = Inertia{}
	return m
}

var (
	MsgInertiaStamped = &_MsgInertiaStamped{
		`Header header
Inertia inertia
`,
		"geometry_msgs/InertiaStamped",
		"ddee48caeab5a966c5e8d166654a9ac7",
	}
)

type InertiaStamped struct {
	Header  std_msgs.Header `rosmsg:"header:Header"`
	Inertia Inertia         `rosmsg:"inertia:Inertia"`
}

func (m *InertiaStamped) GetType() ros.MessageType {
	return MsgInertiaStamped
}

func (m *InertiaStamped) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	if err = m.Inertia.Serialize(buf); err != nil {
		return err
	}
	return err
}

func (m *InertiaStamped) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = m.Inertia.Deserialize(buf); err != nil {
		return err
	}
	return err
}

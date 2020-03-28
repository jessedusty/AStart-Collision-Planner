// Automatically generated from the message definition "geometry_msgs/AccelStamped.msg"
package geometry_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgAccelStamped struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgAccelStamped) Text() string {
	return t.text
}

func (t *_MsgAccelStamped) Name() string {
	return t.name
}

func (t *_MsgAccelStamped) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgAccelStamped) NewMessage() ros.Message {
	m := new(AccelStamped)
	m.Header = std_msgs.Header{}
	m.Accel = Accel{}
	return m
}

var (
	MsgAccelStamped = &_MsgAccelStamped{
		`# An accel with reference coordinate frame and timestamp
Header header
Accel accel
`,
		"geometry_msgs/AccelStamped",
		"d8a98a5d81351b6eb0578c78557e7659",
	}
)

type AccelStamped struct {
	Header std_msgs.Header `rosmsg:"header:Header"`
	Accel  Accel           `rosmsg:"accel:Accel"`
}

func (m *AccelStamped) GetType() ros.MessageType {
	return MsgAccelStamped
}

func (m *AccelStamped) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	if err = m.Accel.Serialize(buf); err != nil {
		return err
	}
	return err
}

func (m *AccelStamped) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = m.Accel.Deserialize(buf); err != nil {
		return err
	}
	return err
}

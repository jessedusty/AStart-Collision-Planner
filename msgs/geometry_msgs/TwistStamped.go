// Automatically generated from the message definition "geometry_msgs/TwistStamped.msg"
package geometry_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgTwistStamped struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgTwistStamped) Text() string {
	return t.text
}

func (t *_MsgTwistStamped) Name() string {
	return t.name
}

func (t *_MsgTwistStamped) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgTwistStamped) NewMessage() ros.Message {
	m := new(TwistStamped)
	m.Header = std_msgs.Header{}
	m.Twist = Twist{}
	return m
}

var (
	MsgTwistStamped = &_MsgTwistStamped{
		`# A twist with reference coordinate frame and timestamp
Header header
Twist twist
`,
		"geometry_msgs/TwistStamped",
		"98d34b0043a2093cf9d9345ab6eef12e",
	}
)

type TwistStamped struct {
	Header std_msgs.Header `rosmsg:"header:Header"`
	Twist  Twist           `rosmsg:"twist:Twist"`
}

func (m *TwistStamped) GetType() ros.MessageType {
	return MsgTwistStamped
}

func (m *TwistStamped) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	if err = m.Twist.Serialize(buf); err != nil {
		return err
	}
	return err
}

func (m *TwistStamped) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = m.Twist.Deserialize(buf); err != nil {
		return err
	}
	return err
}

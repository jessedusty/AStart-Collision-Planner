// Automatically generated from the message definition "geometry_msgs/PointStamped.msg"
package geometry_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgPointStamped struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPointStamped) Text() string {
	return t.text
}

func (t *_MsgPointStamped) Name() string {
	return t.name
}

func (t *_MsgPointStamped) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPointStamped) NewMessage() ros.Message {
	m := new(PointStamped)
	m.Header = std_msgs.Header{}
	m.Point = Point{}
	return m
}

var (
	MsgPointStamped = &_MsgPointStamped{
		`# This represents a Point with reference coordinate frame and timestamp
Header header
Point point
`,
		"geometry_msgs/PointStamped",
		"c63aecb41bfdfd6b7e1fac37c7cbe7bf",
	}
)

type PointStamped struct {
	Header std_msgs.Header `rosmsg:"header:Header"`
	Point  Point           `rosmsg:"point:Point"`
}

func (m *PointStamped) GetType() ros.MessageType {
	return MsgPointStamped
}

func (m *PointStamped) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	if err = m.Point.Serialize(buf); err != nil {
		return err
	}
	return err
}

func (m *PointStamped) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = m.Point.Deserialize(buf); err != nil {
		return err
	}
	return err
}

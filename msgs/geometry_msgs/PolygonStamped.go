// Automatically generated from the message definition "geometry_msgs/PolygonStamped.msg"
package geometry_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgPolygonStamped struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPolygonStamped) Text() string {
	return t.text
}

func (t *_MsgPolygonStamped) Name() string {
	return t.name
}

func (t *_MsgPolygonStamped) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPolygonStamped) NewMessage() ros.Message {
	m := new(PolygonStamped)
	m.Header = std_msgs.Header{}
	m.Polygon = Polygon{}
	return m
}

var (
	MsgPolygonStamped = &_MsgPolygonStamped{
		`# This represents a Polygon with reference coordinate frame and timestamp
Header header
Polygon polygon
`,
		"geometry_msgs/PolygonStamped",
		"c6be8f7dc3bee7fe9e8d296070f53340",
	}
)

type PolygonStamped struct {
	Header  std_msgs.Header `rosmsg:"header:Header"`
	Polygon Polygon         `rosmsg:"polygon:Polygon"`
}

func (m *PolygonStamped) GetType() ros.MessageType {
	return MsgPolygonStamped
}

func (m *PolygonStamped) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	if err = m.Polygon.Serialize(buf); err != nil {
		return err
	}
	return err
}

func (m *PolygonStamped) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = m.Polygon.Deserialize(buf); err != nil {
		return err
	}
	return err
}

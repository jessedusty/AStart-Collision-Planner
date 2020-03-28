// Automatically generated from the message definition "geometry_msgs/WrenchStamped.msg"
package geometry_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgWrenchStamped struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgWrenchStamped) Text() string {
	return t.text
}

func (t *_MsgWrenchStamped) Name() string {
	return t.name
}

func (t *_MsgWrenchStamped) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgWrenchStamped) NewMessage() ros.Message {
	m := new(WrenchStamped)
	m.Header = std_msgs.Header{}
	m.Wrench = Wrench{}
	return m
}

var (
	MsgWrenchStamped = &_MsgWrenchStamped{
		`# A wrench with reference coordinate frame and timestamp
Header header
Wrench wrench
`,
		"geometry_msgs/WrenchStamped",
		"d78d3cb249ce23087ade7e7d0c40cfa7",
	}
)

type WrenchStamped struct {
	Header std_msgs.Header `rosmsg:"header:Header"`
	Wrench Wrench          `rosmsg:"wrench:Wrench"`
}

func (m *WrenchStamped) GetType() ros.MessageType {
	return MsgWrenchStamped
}

func (m *WrenchStamped) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	if err = m.Wrench.Serialize(buf); err != nil {
		return err
	}
	return err
}

func (m *WrenchStamped) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = m.Wrench.Deserialize(buf); err != nil {
		return err
	}
	return err
}

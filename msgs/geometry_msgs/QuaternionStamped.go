// Automatically generated from the message definition "geometry_msgs/QuaternionStamped.msg"
package geometry_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgQuaternionStamped struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgQuaternionStamped) Text() string {
	return t.text
}

func (t *_MsgQuaternionStamped) Name() string {
	return t.name
}

func (t *_MsgQuaternionStamped) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgQuaternionStamped) NewMessage() ros.Message {
	m := new(QuaternionStamped)
	m.Header = std_msgs.Header{}
	m.Quaternion = Quaternion{}
	return m
}

var (
	MsgQuaternionStamped = &_MsgQuaternionStamped{
		`# This represents an orientation with reference coordinate frame and timestamp.

Header header
Quaternion quaternion
`,
		"geometry_msgs/QuaternionStamped",
		"e57f1e547e0e1fd13504588ffc8334e2",
	}
)

type QuaternionStamped struct {
	Header     std_msgs.Header `rosmsg:"header:Header"`
	Quaternion Quaternion      `rosmsg:"quaternion:Quaternion"`
}

func (m *QuaternionStamped) GetType() ros.MessageType {
	return MsgQuaternionStamped
}

func (m *QuaternionStamped) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	if err = m.Quaternion.Serialize(buf); err != nil {
		return err
	}
	return err
}

func (m *QuaternionStamped) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = m.Quaternion.Deserialize(buf); err != nil {
		return err
	}
	return err
}

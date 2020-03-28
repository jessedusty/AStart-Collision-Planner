// Automatically generated from the message definition "geometry_msgs/TwistWithCovarianceStamped.msg"
package geometry_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgTwistWithCovarianceStamped struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgTwistWithCovarianceStamped) Text() string {
	return t.text
}

func (t *_MsgTwistWithCovarianceStamped) Name() string {
	return t.name
}

func (t *_MsgTwistWithCovarianceStamped) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgTwistWithCovarianceStamped) NewMessage() ros.Message {
	m := new(TwistWithCovarianceStamped)
	m.Header = std_msgs.Header{}
	m.Twist = TwistWithCovariance{}
	return m
}

var (
	MsgTwistWithCovarianceStamped = &_MsgTwistWithCovarianceStamped{
		`# This represents an estimated twist with reference coordinate frame and timestamp.
Header header
TwistWithCovariance twist
`,
		"geometry_msgs/TwistWithCovarianceStamped",
		"8927a1a12fb2607ceea095b2dc440a96",
	}
)

type TwistWithCovarianceStamped struct {
	Header std_msgs.Header     `rosmsg:"header:Header"`
	Twist  TwistWithCovariance `rosmsg:"twist:TwistWithCovariance"`
}

func (m *TwistWithCovarianceStamped) GetType() ros.MessageType {
	return MsgTwistWithCovarianceStamped
}

func (m *TwistWithCovarianceStamped) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	if err = m.Twist.Serialize(buf); err != nil {
		return err
	}
	return err
}

func (m *TwistWithCovarianceStamped) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = m.Twist.Deserialize(buf); err != nil {
		return err
	}
	return err
}

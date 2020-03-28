// Automatically generated from the message definition "geometry_msgs/AccelWithCovarianceStamped.msg"
package geometry_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgAccelWithCovarianceStamped struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgAccelWithCovarianceStamped) Text() string {
	return t.text
}

func (t *_MsgAccelWithCovarianceStamped) Name() string {
	return t.name
}

func (t *_MsgAccelWithCovarianceStamped) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgAccelWithCovarianceStamped) NewMessage() ros.Message {
	m := new(AccelWithCovarianceStamped)
	m.Header = std_msgs.Header{}
	m.Accel = AccelWithCovariance{}
	return m
}

var (
	MsgAccelWithCovarianceStamped = &_MsgAccelWithCovarianceStamped{
		`# This represents an estimated accel with reference coordinate frame and timestamp.
Header header
AccelWithCovariance accel
`,
		"geometry_msgs/AccelWithCovarianceStamped",
		"96adb295225031ec8d57fb4251b0a886",
	}
)

type AccelWithCovarianceStamped struct {
	Header std_msgs.Header     `rosmsg:"header:Header"`
	Accel  AccelWithCovariance `rosmsg:"accel:AccelWithCovariance"`
}

func (m *AccelWithCovarianceStamped) GetType() ros.MessageType {
	return MsgAccelWithCovarianceStamped
}

func (m *AccelWithCovarianceStamped) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	if err = m.Accel.Serialize(buf); err != nil {
		return err
	}
	return err
}

func (m *AccelWithCovarianceStamped) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = m.Accel.Deserialize(buf); err != nil {
		return err
	}
	return err
}

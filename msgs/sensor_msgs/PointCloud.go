// Automatically generated from the message definition "sensor_msgs/PointCloud.msg"
package sensor_msgs

import (
	"bytes"
	"ee631_midterm/msgs/geometry_msgs"
	"ee631_midterm/msgs/std_msgs"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgPointCloud struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPointCloud) Text() string {
	return t.text
}

func (t *_MsgPointCloud) Name() string {
	return t.name
}

func (t *_MsgPointCloud) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPointCloud) NewMessage() ros.Message {
	m := new(PointCloud)
	m.Header = std_msgs.Header{}
	m.Points = []geometry_msgs.Point32{}
	m.Channels = []ChannelFloat32{}
	return m
}

var (
	MsgPointCloud = &_MsgPointCloud{
		`# This message holds a collection of 3d points, plus optional additional
# information about each point.

# Time of sensor data acquisition, coordinate frame ID.
Header header

# Array of 3d points. Each Point32 should be interpreted as a 3d point
# in the frame given in the header.
geometry_msgs/Point32[] points

# Each channel should have the same number of elements as points array,
# and the data in each channel should correspond 1:1 with each point.
# Channel names in common practice are listed in ChannelFloat32.msg.
ChannelFloat32[] channels
`,
		"sensor_msgs/PointCloud",
		"d8e9c3f5afbdd8a130fd1d2763945fca",
	}
)

type PointCloud struct {
	Header   std_msgs.Header         `rosmsg:"header:Header"`
	Points   []geometry_msgs.Point32 `rosmsg:"points:Point32[]"`
	Channels []ChannelFloat32        `rosmsg:"channels:ChannelFloat32[]"`
}

func (m *PointCloud) GetType() ros.MessageType {
	return MsgPointCloud
}

func (m *PointCloud) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Points)))
	for _, e := range m.Points {
		if err = e.Serialize(buf); err != nil {
			return err
		}
	}
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Channels)))
	for _, e := range m.Channels {
		if err = e.Serialize(buf); err != nil {
			return err
		}
	}
	return err
}

func (m *PointCloud) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Points = make([]geometry_msgs.Point32, int(size))
		for i := 0; i < int(size); i++ {
			if err = m.Points[i].Deserialize(buf); err != nil {
				return err
			}
		}
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Channels = make([]ChannelFloat32, int(size))
		for i := 0; i < int(size); i++ {
			if err = m.Channels[i].Deserialize(buf); err != nil {
				return err
			}
		}
	}
	return err
}

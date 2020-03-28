// Automatically generated from the message definition "map_msgs/PointCloud2Update.msg"
package map_msgs

import (
	"bytes"
	"ee631_midterm/msgs/sensor_msgs"
	"ee631_midterm/msgs/std_msgs"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

const (
	PointCloud2Update_ADD    uint32 = 0
	PointCloud2Update_DELETE uint32 = 1
)

type _MsgPointCloud2Update struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPointCloud2Update) Text() string {
	return t.text
}

func (t *_MsgPointCloud2Update) Name() string {
	return t.name
}

func (t *_MsgPointCloud2Update) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPointCloud2Update) NewMessage() ros.Message {
	m := new(PointCloud2Update)
	m.Header = std_msgs.Header{}
	m.Type = 0
	m.Points = sensor_msgs.PointCloud2{}
	return m
}

var (
	MsgPointCloud2Update = &_MsgPointCloud2Update{
		`uint32 ADD=0
uint32 DELETE=1
Header header
uint32 type          # type of update, one of ADD or DELETE
sensor_msgs/PointCloud2 points
`,
		"map_msgs/PointCloud2Update",
		"6c58e4f249ae9cd2b24fb1ee0f99195e",
	}
)

type PointCloud2Update struct {
	Header std_msgs.Header         `rosmsg:"header:Header"`
	Type   uint32                  `rosmsg:"type:uint32"`
	Points sensor_msgs.PointCloud2 `rosmsg:"points:PointCloud2"`
}

func (m *PointCloud2Update) GetType() ros.MessageType {
	return MsgPointCloud2Update
}

func (m *PointCloud2Update) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, m.Type)
	if err = m.Points.Serialize(buf); err != nil {
		return err
	}
	return err
}

func (m *PointCloud2Update) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Type); err != nil {
		return err
	}
	if err = m.Points.Deserialize(buf); err != nil {
		return err
	}
	return err
}

// Automatically generated from the message definition "sensor_msgs/PointCloud2.msg"
package sensor_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgPointCloud2 struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPointCloud2) Text() string {
	return t.text
}

func (t *_MsgPointCloud2) Name() string {
	return t.name
}

func (t *_MsgPointCloud2) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPointCloud2) NewMessage() ros.Message {
	m := new(PointCloud2)
	m.Header = std_msgs.Header{}
	m.Height = 0
	m.Width = 0
	m.Fields = []PointField{}
	m.IsBigendian = false
	m.PointStep = 0
	m.RowStep = 0
	m.Data = []uint8{}
	m.IsDense = false
	return m
}

var (
	MsgPointCloud2 = &_MsgPointCloud2{
		`# This message holds a collection of N-dimensional points, which may
# contain additional information such as normals, intensity, etc. The
# point data is stored as a binary blob, its layout described by the
# contents of the "fields" array.

# The point cloud data may be organized 2d (image-like) or 1d
# (unordered). Point clouds organized as 2d images may be produced by
# camera depth sensors such as stereo or time-of-flight.

# Time of sensor data acquisition, and the coordinate frame ID (for 3d
# points).
Header header

# 2D structure of the point cloud. If the cloud is unordered, height is
# 1 and width is the length of the point cloud.
uint32 height
uint32 width

# Describes the channels and their layout in the binary data blob.
PointField[] fields

bool    is_bigendian # Is this data bigendian?
uint32  point_step   # Length of a point in bytes
uint32  row_step     # Length of a row in bytes
uint8[] data         # Actual point data, size is (row_step*height)

bool is_dense        # True if there are no invalid points
`,
		"sensor_msgs/PointCloud2",
		"1158d486dd51d683ce2f1be655c3c181",
	}
)

type PointCloud2 struct {
	Header      std_msgs.Header `rosmsg:"header:Header"`
	Height      uint32          `rosmsg:"height:uint32"`
	Width       uint32          `rosmsg:"width:uint32"`
	Fields      []PointField    `rosmsg:"fields:PointField[]"`
	IsBigendian bool            `rosmsg:"is_bigendian:bool"`
	PointStep   uint32          `rosmsg:"point_step:uint32"`
	RowStep     uint32          `rosmsg:"row_step:uint32"`
	Data        []uint8         `rosmsg:"data:uint8[]"`
	IsDense     bool            `rosmsg:"is_dense:bool"`
}

func (m *PointCloud2) GetType() ros.MessageType {
	return MsgPointCloud2
}

func (m *PointCloud2) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, m.Height)
	binary.Write(buf, binary.LittleEndian, m.Width)
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Fields)))
	for _, e := range m.Fields {
		if err = e.Serialize(buf); err != nil {
			return err
		}
	}
	binary.Write(buf, binary.LittleEndian, m.IsBigendian)
	binary.Write(buf, binary.LittleEndian, m.PointStep)
	binary.Write(buf, binary.LittleEndian, m.RowStep)
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Data)))
	for _, e := range m.Data {
		binary.Write(buf, binary.LittleEndian, e)
	}
	binary.Write(buf, binary.LittleEndian, m.IsDense)
	return err
}

func (m *PointCloud2) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Height); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Width); err != nil {
		return err
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Fields = make([]PointField, int(size))
		for i := 0; i < int(size); i++ {
			if err = m.Fields[i].Deserialize(buf); err != nil {
				return err
			}
		}
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.IsBigendian); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.PointStep); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.RowStep); err != nil {
		return err
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Data = make([]uint8, int(size))
		for i := 0; i < int(size); i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.Data[i]); err != nil {
				return err
			}
		}
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.IsDense); err != nil {
		return err
	}
	return err
}

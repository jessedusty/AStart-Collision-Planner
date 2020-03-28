// Automatically generated from the message definition "geometry_msgs/Point32.msg"
package geometry_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgPoint32 struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPoint32) Text() string {
	return t.text
}

func (t *_MsgPoint32) Name() string {
	return t.name
}

func (t *_MsgPoint32) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPoint32) NewMessage() ros.Message {
	m := new(Point32)
	m.X = 0.0
	m.Y = 0.0
	m.Z = 0.0
	return m
}

var (
	MsgPoint32 = &_MsgPoint32{
		`# This contains the position of a point in free space(with 32 bits of precision).
# It is recommeded to use Point wherever possible instead of Point32.  
# 
# This recommendation is to promote interoperability.  
#
# This message is designed to take up less space when sending
# lots of points at once, as in the case of a PointCloud.  

float32 x
float32 y
float32 z`,
		"geometry_msgs/Point32",
		"cc153912f1453b708d221682bc23d9ac",
	}
)

type Point32 struct {
	X float32 `rosmsg:"x:float32"`
	Y float32 `rosmsg:"y:float32"`
	Z float32 `rosmsg:"z:float32"`
}

func (m *Point32) GetType() ros.MessageType {
	return MsgPoint32
}

func (m *Point32) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, m.X)
	binary.Write(buf, binary.LittleEndian, m.Y)
	binary.Write(buf, binary.LittleEndian, m.Z)
	return err
}

func (m *Point32) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = binary.Read(buf, binary.LittleEndian, &m.X); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Y); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Z); err != nil {
		return err
	}
	return err
}

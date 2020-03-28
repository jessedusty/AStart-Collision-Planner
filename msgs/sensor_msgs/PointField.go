// Automatically generated from the message definition "sensor_msgs/PointField.msg"
package sensor_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

const (
	PointField_INT8    uint8 = 1
	PointField_UINT8   uint8 = 2
	PointField_INT16   uint8 = 3
	PointField_UINT16  uint8 = 4
	PointField_INT32   uint8 = 5
	PointField_UINT32  uint8 = 6
	PointField_FLOAT32 uint8 = 7
	PointField_FLOAT64 uint8 = 8
)

type _MsgPointField struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPointField) Text() string {
	return t.text
}

func (t *_MsgPointField) Name() string {
	return t.name
}

func (t *_MsgPointField) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPointField) NewMessage() ros.Message {
	m := new(PointField)
	m.Name = ""
	m.Offset = 0
	m.Datatype = 0
	m.Count = 0
	return m
}

var (
	MsgPointField = &_MsgPointField{
		`# This message holds the description of one point entry in the
# PointCloud2 message format.
uint8 INT8    = 1
uint8 UINT8   = 2
uint8 INT16   = 3
uint8 UINT16  = 4
uint8 INT32   = 5
uint8 UINT32  = 6
uint8 FLOAT32 = 7
uint8 FLOAT64 = 8

string name      # Name of field
uint32 offset    # Offset from start of point struct
uint8  datatype  # Datatype enumeration, see above
uint32 count     # How many elements in the field
`,
		"sensor_msgs/PointField",
		"268eacb2962780ceac86cbd17e328150",
	}
)

type PointField struct {
	Name     string `rosmsg:"name:string"`
	Offset   uint32 `rosmsg:"offset:uint32"`
	Datatype uint8  `rosmsg:"datatype:uint8"`
	Count    uint32 `rosmsg:"count:uint32"`
}

func (m *PointField) GetType() ros.MessageType {
	return MsgPointField
}

func (m *PointField) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, uint32(len([]byte(m.Name))))
	buf.Write([]byte(m.Name))
	binary.Write(buf, binary.LittleEndian, m.Offset)
	binary.Write(buf, binary.LittleEndian, m.Datatype)
	binary.Write(buf, binary.LittleEndian, m.Count)
	return err
}

func (m *PointField) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		data := make([]byte, int(size))
		if err = binary.Read(buf, binary.LittleEndian, data); err != nil {
			return err
		}
		m.Name = string(data)
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Offset); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Datatype); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Count); err != nil {
		return err
	}
	return err
}

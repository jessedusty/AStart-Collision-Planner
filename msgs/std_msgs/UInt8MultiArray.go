// Automatically generated from the message definition "std_msgs/UInt8MultiArray.msg"
package std_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgUInt8MultiArray struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgUInt8MultiArray) Text() string {
	return t.text
}

func (t *_MsgUInt8MultiArray) Name() string {
	return t.name
}

func (t *_MsgUInt8MultiArray) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgUInt8MultiArray) NewMessage() ros.Message {
	m := new(UInt8MultiArray)
	m.Layout = MultiArrayLayout{}
	m.Data = []uint8{}
	return m
}

var (
	MsgUInt8MultiArray = &_MsgUInt8MultiArray{
		`# Please look at the MultiArrayLayout message definition for
# documentation on all multiarrays.

MultiArrayLayout  layout        # specification of data layout
uint8[]           data          # array of data

`,
		"std_msgs/UInt8MultiArray",
		"82373f1612381bb6ee473b5cd6f5d89c",
	}
)

type UInt8MultiArray struct {
	Layout MultiArrayLayout `rosmsg:"layout:MultiArrayLayout"`
	Data   []uint8          `rosmsg:"data:uint8[]"`
}

func (m *UInt8MultiArray) GetType() ros.MessageType {
	return MsgUInt8MultiArray
}

func (m *UInt8MultiArray) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Layout.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Data)))
	for _, e := range m.Data {
		binary.Write(buf, binary.LittleEndian, e)
	}
	return err
}

func (m *UInt8MultiArray) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Layout.Deserialize(buf); err != nil {
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
	return err
}

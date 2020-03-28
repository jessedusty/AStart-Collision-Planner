// Automatically generated from the message definition "std_msgs/UInt32MultiArray.msg"
package std_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgUInt32MultiArray struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgUInt32MultiArray) Text() string {
	return t.text
}

func (t *_MsgUInt32MultiArray) Name() string {
	return t.name
}

func (t *_MsgUInt32MultiArray) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgUInt32MultiArray) NewMessage() ros.Message {
	m := new(UInt32MultiArray)
	m.Layout = MultiArrayLayout{}
	m.Data = []uint32{}
	return m
}

var (
	MsgUInt32MultiArray = &_MsgUInt32MultiArray{
		`# Please look at the MultiArrayLayout message definition for
# documentation on all multiarrays.

MultiArrayLayout  layout        # specification of data layout
uint32[]          data          # array of data

`,
		"std_msgs/UInt32MultiArray",
		"4d6a180abc9be191b96a7eda6c8a233d",
	}
)

type UInt32MultiArray struct {
	Layout MultiArrayLayout `rosmsg:"layout:MultiArrayLayout"`
	Data   []uint32         `rosmsg:"data:uint32[]"`
}

func (m *UInt32MultiArray) GetType() ros.MessageType {
	return MsgUInt32MultiArray
}

func (m *UInt32MultiArray) Serialize(buf *bytes.Buffer) error {
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

func (m *UInt32MultiArray) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Layout.Deserialize(buf); err != nil {
		return err
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Data = make([]uint32, int(size))
		for i := 0; i < int(size); i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.Data[i]); err != nil {
				return err
			}
		}
	}
	return err
}

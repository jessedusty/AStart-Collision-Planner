// Automatically generated from the message definition "std_msgs/UInt64MultiArray.msg"
package std_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgUInt64MultiArray struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgUInt64MultiArray) Text() string {
	return t.text
}

func (t *_MsgUInt64MultiArray) Name() string {
	return t.name
}

func (t *_MsgUInt64MultiArray) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgUInt64MultiArray) NewMessage() ros.Message {
	m := new(UInt64MultiArray)
	m.Layout = MultiArrayLayout{}
	m.Data = []uint64{}
	return m
}

var (
	MsgUInt64MultiArray = &_MsgUInt64MultiArray{
		`# Please look at the MultiArrayLayout message definition for
# documentation on all multiarrays.

MultiArrayLayout  layout        # specification of data layout
uint64[]          data          # array of data

`,
		"std_msgs/UInt64MultiArray",
		"6088f127afb1d6c72927aa1247e945af",
	}
)

type UInt64MultiArray struct {
	Layout MultiArrayLayout `rosmsg:"layout:MultiArrayLayout"`
	Data   []uint64         `rosmsg:"data:uint64[]"`
}

func (m *UInt64MultiArray) GetType() ros.MessageType {
	return MsgUInt64MultiArray
}

func (m *UInt64MultiArray) Serialize(buf *bytes.Buffer) error {
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

func (m *UInt64MultiArray) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Layout.Deserialize(buf); err != nil {
		return err
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Data = make([]uint64, int(size))
		for i := 0; i < int(size); i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.Data[i]); err != nil {
				return err
			}
		}
	}
	return err
}

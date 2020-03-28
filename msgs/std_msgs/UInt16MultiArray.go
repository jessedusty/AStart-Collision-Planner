// Automatically generated from the message definition "std_msgs/UInt16MultiArray.msg"
package std_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgUInt16MultiArray struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgUInt16MultiArray) Text() string {
	return t.text
}

func (t *_MsgUInt16MultiArray) Name() string {
	return t.name
}

func (t *_MsgUInt16MultiArray) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgUInt16MultiArray) NewMessage() ros.Message {
	m := new(UInt16MultiArray)
	m.Layout = MultiArrayLayout{}
	m.Data = []uint16{}
	return m
}

var (
	MsgUInt16MultiArray = &_MsgUInt16MultiArray{
		`# Please look at the MultiArrayLayout message definition for
# documentation on all multiarrays.

MultiArrayLayout  layout        # specification of data layout
uint16[]            data        # array of data

`,
		"std_msgs/UInt16MultiArray",
		"52f264f1c973c4b73790d384c6cb4484",
	}
)

type UInt16MultiArray struct {
	Layout MultiArrayLayout `rosmsg:"layout:MultiArrayLayout"`
	Data   []uint16         `rosmsg:"data:uint16[]"`
}

func (m *UInt16MultiArray) GetType() ros.MessageType {
	return MsgUInt16MultiArray
}

func (m *UInt16MultiArray) Serialize(buf *bytes.Buffer) error {
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

func (m *UInt16MultiArray) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Layout.Deserialize(buf); err != nil {
		return err
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Data = make([]uint16, int(size))
		for i := 0; i < int(size); i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.Data[i]); err != nil {
				return err
			}
		}
	}
	return err
}

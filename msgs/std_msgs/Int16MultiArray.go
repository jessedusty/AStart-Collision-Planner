// Automatically generated from the message definition "std_msgs/Int16MultiArray.msg"
package std_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgInt16MultiArray struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgInt16MultiArray) Text() string {
	return t.text
}

func (t *_MsgInt16MultiArray) Name() string {
	return t.name
}

func (t *_MsgInt16MultiArray) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgInt16MultiArray) NewMessage() ros.Message {
	m := new(Int16MultiArray)
	m.Layout = MultiArrayLayout{}
	m.Data = []int16{}
	return m
}

var (
	MsgInt16MultiArray = &_MsgInt16MultiArray{
		`# Please look at the MultiArrayLayout message definition for
# documentation on all multiarrays.

MultiArrayLayout  layout        # specification of data layout
int16[]           data          # array of data

`,
		"std_msgs/Int16MultiArray",
		"d9338d7f523fcb692fae9d0a0e9f067c",
	}
)

type Int16MultiArray struct {
	Layout MultiArrayLayout `rosmsg:"layout:MultiArrayLayout"`
	Data   []int16          `rosmsg:"data:int16[]"`
}

func (m *Int16MultiArray) GetType() ros.MessageType {
	return MsgInt16MultiArray
}

func (m *Int16MultiArray) Serialize(buf *bytes.Buffer) error {
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

func (m *Int16MultiArray) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Layout.Deserialize(buf); err != nil {
		return err
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Data = make([]int16, int(size))
		for i := 0; i < int(size); i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.Data[i]); err != nil {
				return err
			}
		}
	}
	return err
}

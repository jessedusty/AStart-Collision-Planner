// Automatically generated from the message definition "std_msgs/Int8MultiArray.msg"
package std_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgInt8MultiArray struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgInt8MultiArray) Text() string {
	return t.text
}

func (t *_MsgInt8MultiArray) Name() string {
	return t.name
}

func (t *_MsgInt8MultiArray) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgInt8MultiArray) NewMessage() ros.Message {
	m := new(Int8MultiArray)
	m.Layout = MultiArrayLayout{}
	m.Data = []int8{}
	return m
}

var (
	MsgInt8MultiArray = &_MsgInt8MultiArray{
		`# Please look at the MultiArrayLayout message definition for
# documentation on all multiarrays.

MultiArrayLayout  layout        # specification of data layout
int8[]            data          # array of data

`,
		"std_msgs/Int8MultiArray",
		"d7c1af35a1b4781bbe79e03dd94b7c13",
	}
)

type Int8MultiArray struct {
	Layout MultiArrayLayout `rosmsg:"layout:MultiArrayLayout"`
	Data   []int8           `rosmsg:"data:int8[]"`
}

func (m *Int8MultiArray) GetType() ros.MessageType {
	return MsgInt8MultiArray
}

func (m *Int8MultiArray) Serialize(buf *bytes.Buffer) error {
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

func (m *Int8MultiArray) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Layout.Deserialize(buf); err != nil {
		return err
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Data = make([]int8, int(size))
		for i := 0; i < int(size); i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.Data[i]); err != nil {
				return err
			}
		}
	}
	return err
}

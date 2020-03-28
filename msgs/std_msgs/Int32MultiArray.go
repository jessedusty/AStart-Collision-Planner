// Automatically generated from the message definition "std_msgs/Int32MultiArray.msg"
package std_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgInt32MultiArray struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgInt32MultiArray) Text() string {
	return t.text
}

func (t *_MsgInt32MultiArray) Name() string {
	return t.name
}

func (t *_MsgInt32MultiArray) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgInt32MultiArray) NewMessage() ros.Message {
	m := new(Int32MultiArray)
	m.Layout = MultiArrayLayout{}
	m.Data = []int32{}
	return m
}

var (
	MsgInt32MultiArray = &_MsgInt32MultiArray{
		`# Please look at the MultiArrayLayout message definition for
# documentation on all multiarrays.

MultiArrayLayout  layout        # specification of data layout
int32[]           data          # array of data

`,
		"std_msgs/Int32MultiArray",
		"1d99f79f8b325b44fee908053e9c945b",
	}
)

type Int32MultiArray struct {
	Layout MultiArrayLayout `rosmsg:"layout:MultiArrayLayout"`
	Data   []int32          `rosmsg:"data:int32[]"`
}

func (m *Int32MultiArray) GetType() ros.MessageType {
	return MsgInt32MultiArray
}

func (m *Int32MultiArray) Serialize(buf *bytes.Buffer) error {
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

func (m *Int32MultiArray) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Layout.Deserialize(buf); err != nil {
		return err
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Data = make([]int32, int(size))
		for i := 0; i < int(size); i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.Data[i]); err != nil {
				return err
			}
		}
	}
	return err
}

// Automatically generated from the message definition "std_msgs/Int64MultiArray.msg"
package std_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgInt64MultiArray struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgInt64MultiArray) Text() string {
	return t.text
}

func (t *_MsgInt64MultiArray) Name() string {
	return t.name
}

func (t *_MsgInt64MultiArray) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgInt64MultiArray) NewMessage() ros.Message {
	m := new(Int64MultiArray)
	m.Layout = MultiArrayLayout{}
	m.Data = []int64{}
	return m
}

var (
	MsgInt64MultiArray = &_MsgInt64MultiArray{
		`# Please look at the MultiArrayLayout message definition for
# documentation on all multiarrays.

MultiArrayLayout  layout        # specification of data layout
int64[]           data          # array of data

`,
		"std_msgs/Int64MultiArray",
		"54865aa6c65be0448113a2afc6a49270",
	}
)

type Int64MultiArray struct {
	Layout MultiArrayLayout `rosmsg:"layout:MultiArrayLayout"`
	Data   []int64          `rosmsg:"data:int64[]"`
}

func (m *Int64MultiArray) GetType() ros.MessageType {
	return MsgInt64MultiArray
}

func (m *Int64MultiArray) Serialize(buf *bytes.Buffer) error {
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

func (m *Int64MultiArray) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Layout.Deserialize(buf); err != nil {
		return err
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Data = make([]int64, int(size))
		for i := 0; i < int(size); i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.Data[i]); err != nil {
				return err
			}
		}
	}
	return err
}

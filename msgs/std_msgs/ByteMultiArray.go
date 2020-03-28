// Automatically generated from the message definition "std_msgs/ByteMultiArray.msg"
package std_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgByteMultiArray struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgByteMultiArray) Text() string {
	return t.text
}

func (t *_MsgByteMultiArray) Name() string {
	return t.name
}

func (t *_MsgByteMultiArray) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgByteMultiArray) NewMessage() ros.Message {
	m := new(ByteMultiArray)
	m.Layout = MultiArrayLayout{}
	m.Data = []uint8{}
	return m
}

var (
	MsgByteMultiArray = &_MsgByteMultiArray{
		`# Please look at the MultiArrayLayout message definition for
# documentation on all multiarrays.

MultiArrayLayout  layout        # specification of data layout
byte[]            data          # array of data

`,
		"std_msgs/ByteMultiArray",
		"70ea476cbcfd65ac2f68f3cda1e891fe",
	}
)

type ByteMultiArray struct {
	Layout MultiArrayLayout `rosmsg:"layout:MultiArrayLayout"`
	Data   []uint8          `rosmsg:"data:byte[]"`
}

func (m *ByteMultiArray) GetType() ros.MessageType {
	return MsgByteMultiArray
}

func (m *ByteMultiArray) Serialize(buf *bytes.Buffer) error {
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

func (m *ByteMultiArray) Deserialize(buf *bytes.Reader) error {
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

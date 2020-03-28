// Automatically generated from the message definition "std_msgs/MultiArrayDimension.msg"
package std_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgMultiArrayDimension struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgMultiArrayDimension) Text() string {
	return t.text
}

func (t *_MsgMultiArrayDimension) Name() string {
	return t.name
}

func (t *_MsgMultiArrayDimension) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgMultiArrayDimension) NewMessage() ros.Message {
	m := new(MultiArrayDimension)
	m.Label = ""
	m.Size = 0
	m.Stride = 0
	return m
}

var (
	MsgMultiArrayDimension = &_MsgMultiArrayDimension{
		`string label   # label of given dimension
uint32 size    # size of given dimension (in type units)
uint32 stride  # stride of given dimension`,
		"std_msgs/MultiArrayDimension",
		"4cd0c83a8683deae40ecdac60e53bfa8",
	}
)

type MultiArrayDimension struct {
	Label  string `rosmsg:"label:string"`
	Size   uint32 `rosmsg:"size:uint32"`
	Stride uint32 `rosmsg:"stride:uint32"`
}

func (m *MultiArrayDimension) GetType() ros.MessageType {
	return MsgMultiArrayDimension
}

func (m *MultiArrayDimension) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, uint32(len([]byte(m.Label))))
	buf.Write([]byte(m.Label))
	binary.Write(buf, binary.LittleEndian, m.Size)
	binary.Write(buf, binary.LittleEndian, m.Stride)
	return err
}

func (m *MultiArrayDimension) Deserialize(buf *bytes.Reader) error {
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
		m.Label = string(data)
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Size); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Stride); err != nil {
		return err
	}
	return err
}

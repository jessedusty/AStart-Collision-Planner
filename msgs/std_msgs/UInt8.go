// Automatically generated from the message definition "std_msgs/UInt8.msg"
package std_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgUInt8 struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgUInt8) Text() string {
	return t.text
}

func (t *_MsgUInt8) Name() string {
	return t.name
}

func (t *_MsgUInt8) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgUInt8) NewMessage() ros.Message {
	m := new(UInt8)
	m.Data = 0
	return m
}

var (
	MsgUInt8 = &_MsgUInt8{
		`uint8 data
`,
		"std_msgs/UInt8",
		"7c8164229e7d2c17eb95e9231617fdee",
	}
)

type UInt8 struct {
	Data uint8 `rosmsg:"data:uint8"`
}

func (m *UInt8) GetType() ros.MessageType {
	return MsgUInt8
}

func (m *UInt8) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, m.Data)
	return err
}

func (m *UInt8) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = binary.Read(buf, binary.LittleEndian, &m.Data); err != nil {
		return err
	}
	return err
}

// Automatically generated from the message definition "std_msgs/UInt32.msg"
package std_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgUInt32 struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgUInt32) Text() string {
	return t.text
}

func (t *_MsgUInt32) Name() string {
	return t.name
}

func (t *_MsgUInt32) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgUInt32) NewMessage() ros.Message {
	m := new(UInt32)
	m.Data = 0
	return m
}

var (
	MsgUInt32 = &_MsgUInt32{
		`uint32 data`,
		"std_msgs/UInt32",
		"304a39449588c7f8ce2df6e8001c5fce",
	}
)

type UInt32 struct {
	Data uint32 `rosmsg:"data:uint32"`
}

func (m *UInt32) GetType() ros.MessageType {
	return MsgUInt32
}

func (m *UInt32) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, m.Data)
	return err
}

func (m *UInt32) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = binary.Read(buf, binary.LittleEndian, &m.Data); err != nil {
		return err
	}
	return err
}

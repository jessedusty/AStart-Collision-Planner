// Automatically generated from the message definition "std_msgs/UInt16.msg"
package std_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgUInt16 struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgUInt16) Text() string {
	return t.text
}

func (t *_MsgUInt16) Name() string {
	return t.name
}

func (t *_MsgUInt16) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgUInt16) NewMessage() ros.Message {
	m := new(UInt16)
	m.Data = 0
	return m
}

var (
	MsgUInt16 = &_MsgUInt16{
		`uint16 data
`,
		"std_msgs/UInt16",
		"1df79edf208b629fe6b81923a544552d",
	}
)

type UInt16 struct {
	Data uint16 `rosmsg:"data:uint16"`
}

func (m *UInt16) GetType() ros.MessageType {
	return MsgUInt16
}

func (m *UInt16) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, m.Data)
	return err
}

func (m *UInt16) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = binary.Read(buf, binary.LittleEndian, &m.Data); err != nil {
		return err
	}
	return err
}

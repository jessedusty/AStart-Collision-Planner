// Automatically generated from the message definition "std_msgs/UInt64.msg"
package std_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgUInt64 struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgUInt64) Text() string {
	return t.text
}

func (t *_MsgUInt64) Name() string {
	return t.name
}

func (t *_MsgUInt64) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgUInt64) NewMessage() ros.Message {
	m := new(UInt64)
	m.Data = 0
	return m
}

var (
	MsgUInt64 = &_MsgUInt64{
		`uint64 data`,
		"std_msgs/UInt64",
		"1b2a79973e8bf53d7b53acb71299cb57",
	}
)

type UInt64 struct {
	Data uint64 `rosmsg:"data:uint64"`
}

func (m *UInt64) GetType() ros.MessageType {
	return MsgUInt64
}

func (m *UInt64) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, m.Data)
	return err
}

func (m *UInt64) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = binary.Read(buf, binary.LittleEndian, &m.Data); err != nil {
		return err
	}
	return err
}

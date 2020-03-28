// Automatically generated from the message definition "std_msgs/Char.msg"
package std_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgChar struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgChar) Text() string {
	return t.text
}

func (t *_MsgChar) Name() string {
	return t.name
}

func (t *_MsgChar) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgChar) NewMessage() ros.Message {
	m := new(Char)
	m.Data = 0
	return m
}

var (
	MsgChar = &_MsgChar{
		`char data`,
		"std_msgs/Char",
		"1bf77f25acecdedba0e224b162199717",
	}
)

type Char struct {
	Data uint8 `rosmsg:"data:char"`
}

func (m *Char) GetType() ros.MessageType {
	return MsgChar
}

func (m *Char) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, m.Data)
	return err
}

func (m *Char) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = binary.Read(buf, binary.LittleEndian, &m.Data); err != nil {
		return err
	}
	return err
}

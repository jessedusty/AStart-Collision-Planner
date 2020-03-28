// Automatically generated from the message definition "std_msgs/Byte.msg"
package std_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgByte struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgByte) Text() string {
	return t.text
}

func (t *_MsgByte) Name() string {
	return t.name
}

func (t *_MsgByte) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgByte) NewMessage() ros.Message {
	m := new(Byte)
	m.Data = 0
	return m
}

var (
	MsgByte = &_MsgByte{
		`byte data
`,
		"std_msgs/Byte",
		"ad736a2e8818154c487bb80fe42ce43b",
	}
)

type Byte struct {
	Data uint8 `rosmsg:"data:byte"`
}

func (m *Byte) GetType() ros.MessageType {
	return MsgByte
}

func (m *Byte) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, m.Data)
	return err
}

func (m *Byte) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = binary.Read(buf, binary.LittleEndian, &m.Data); err != nil {
		return err
	}
	return err
}

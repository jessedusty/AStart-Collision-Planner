// Automatically generated from the message definition "std_msgs/Int8.msg"
package std_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgInt8 struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgInt8) Text() string {
	return t.text
}

func (t *_MsgInt8) Name() string {
	return t.name
}

func (t *_MsgInt8) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgInt8) NewMessage() ros.Message {
	m := new(Int8)
	m.Data = 0
	return m
}

var (
	MsgInt8 = &_MsgInt8{
		`int8 data
`,
		"std_msgs/Int8",
		"27ffa0c9c4b8fb8492252bcad9e5c57b",
	}
)

type Int8 struct {
	Data int8 `rosmsg:"data:int8"`
}

func (m *Int8) GetType() ros.MessageType {
	return MsgInt8
}

func (m *Int8) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, m.Data)
	return err
}

func (m *Int8) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = binary.Read(buf, binary.LittleEndian, &m.Data); err != nil {
		return err
	}
	return err
}

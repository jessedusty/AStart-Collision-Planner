// Automatically generated from the message definition "std_msgs/Int64.msg"
package std_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgInt64 struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgInt64) Text() string {
	return t.text
}

func (t *_MsgInt64) Name() string {
	return t.name
}

func (t *_MsgInt64) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgInt64) NewMessage() ros.Message {
	m := new(Int64)
	m.Data = 0
	return m
}

var (
	MsgInt64 = &_MsgInt64{
		`int64 data`,
		"std_msgs/Int64",
		"34add168574510e6e17f5d23ecc077ef",
	}
)

type Int64 struct {
	Data int64 `rosmsg:"data:int64"`
}

func (m *Int64) GetType() ros.MessageType {
	return MsgInt64
}

func (m *Int64) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, m.Data)
	return err
}

func (m *Int64) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = binary.Read(buf, binary.LittleEndian, &m.Data); err != nil {
		return err
	}
	return err
}

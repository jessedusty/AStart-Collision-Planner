// Automatically generated from the message definition "std_msgs/Float32.msg"
package std_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgFloat32 struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgFloat32) Text() string {
	return t.text
}

func (t *_MsgFloat32) Name() string {
	return t.name
}

func (t *_MsgFloat32) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgFloat32) NewMessage() ros.Message {
	m := new(Float32)
	m.Data = 0.0
	return m
}

var (
	MsgFloat32 = &_MsgFloat32{
		`float32 data`,
		"std_msgs/Float32",
		"73fcbf46b49191e672908e50842a83d4",
	}
)

type Float32 struct {
	Data float32 `rosmsg:"data:float32"`
}

func (m *Float32) GetType() ros.MessageType {
	return MsgFloat32
}

func (m *Float32) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, m.Data)
	return err
}

func (m *Float32) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = binary.Read(buf, binary.LittleEndian, &m.Data); err != nil {
		return err
	}
	return err
}

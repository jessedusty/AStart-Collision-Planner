// Automatically generated from the message definition "std_msgs/Duration.msg"
package std_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgDuration struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgDuration) Text() string {
	return t.text
}

func (t *_MsgDuration) Name() string {
	return t.name
}

func (t *_MsgDuration) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgDuration) NewMessage() ros.Message {
	m := new(Duration)
	m.Data = ros.Duration{}
	return m
}

var (
	MsgDuration = &_MsgDuration{
		`duration data
`,
		"std_msgs/Duration",
		"3e286caf4241d664e55f3ad380e2ae46",
	}
)

type Duration struct {
	Data ros.Duration `rosmsg:"data:duration"`
}

func (m *Duration) GetType() ros.MessageType {
	return MsgDuration
}

func (m *Duration) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, m.Data.Sec)
	binary.Write(buf, binary.LittleEndian, m.Data.NSec)
	return err
}

func (m *Duration) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	{
		if err = binary.Read(buf, binary.LittleEndian, &m.Data.Sec); err != nil {
			return err
		}
		if err = binary.Read(buf, binary.LittleEndian, &m.Data.NSec); err != nil {
			return err
		}
	}
	return err
}

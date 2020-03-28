// Automatically generated from the message definition "std_msgs/Empty.msg"
package std_msgs

import (
	"bytes"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgEmpty struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgEmpty) Text() string {
	return t.text
}

func (t *_MsgEmpty) Name() string {
	return t.name
}

func (t *_MsgEmpty) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgEmpty) NewMessage() ros.Message {
	m := new(Empty)
	return m
}

var (
	MsgEmpty = &_MsgEmpty{
		``,
		"std_msgs/Empty",
		"d41d8cd98f00b204e9800998ecf8427e",
	}
)

type Empty struct {
}

func (m *Empty) GetType() ros.MessageType {
	return MsgEmpty
}

func (m *Empty) Serialize(buf *bytes.Buffer) error {
	var err error
	return err
}

func (m *Empty) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	return err
}

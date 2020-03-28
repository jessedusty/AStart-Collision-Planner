// Automatically generated from the message definition "tf2_msgs/LookupTransformFeedback.msg"
package tf2_msgs

import (
	"bytes"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgLookupTransformFeedback struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgLookupTransformFeedback) Text() string {
	return t.text
}

func (t *_MsgLookupTransformFeedback) Name() string {
	return t.name
}

func (t *_MsgLookupTransformFeedback) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgLookupTransformFeedback) NewMessage() ros.Message {
	m := new(LookupTransformFeedback)
	return m
}

var (
	MsgLookupTransformFeedback = &_MsgLookupTransformFeedback{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======

`,
		"tf2_msgs/LookupTransformFeedback",
		"d41d8cd98f00b204e9800998ecf8427e",
	}
)

type LookupTransformFeedback struct {
}

func (m *LookupTransformFeedback) GetType() ros.MessageType {
	return MsgLookupTransformFeedback
}

func (m *LookupTransformFeedback) Serialize(buf *bytes.Buffer) error {
	var err error
	return err
}

func (m *LookupTransformFeedback) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	return err
}

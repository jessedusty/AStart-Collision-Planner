// Automatically generated from the message definition "tf2_msgs/LookupTransformResult.msg"
package tf2_msgs

import (
	"bytes"
	"ee631_midterm/msgs/geometry_msgs"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgLookupTransformResult struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgLookupTransformResult) Text() string {
	return t.text
}

func (t *_MsgLookupTransformResult) Name() string {
	return t.name
}

func (t *_MsgLookupTransformResult) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgLookupTransformResult) NewMessage() ros.Message {
	m := new(LookupTransformResult)
	m.Transform = geometry_msgs.TransformStamped{}
	m.Error = TF2Error{}
	return m
}

var (
	MsgLookupTransformResult = &_MsgLookupTransformResult{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======
geometry_msgs/TransformStamped transform
tf2_msgs/TF2Error error
`,
		"tf2_msgs/LookupTransformResult",
		"3fe5db6a19ca9cfb675418c5ad875c36",
	}
)

type LookupTransformResult struct {
	Transform geometry_msgs.TransformStamped `rosmsg:"transform:TransformStamped"`
	Error     TF2Error                       `rosmsg:"error:TF2Error"`
}

func (m *LookupTransformResult) GetType() ros.MessageType {
	return MsgLookupTransformResult
}

func (m *LookupTransformResult) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Transform.Serialize(buf); err != nil {
		return err
	}
	if err = m.Error.Serialize(buf); err != nil {
		return err
	}
	return err
}

func (m *LookupTransformResult) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Transform.Deserialize(buf); err != nil {
		return err
	}
	if err = m.Error.Deserialize(buf); err != nil {
		return err
	}
	return err
}
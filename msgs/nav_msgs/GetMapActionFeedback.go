// Automatically generated from the message definition "nav_msgs/GetMapActionFeedback.msg"
package nav_msgs

import (
	"bytes"
	"ee631_midterm/msgs/actionlib_msgs"
	"ee631_midterm/msgs/std_msgs"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgGetMapActionFeedback struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgGetMapActionFeedback) Text() string {
	return t.text
}

func (t *_MsgGetMapActionFeedback) Name() string {
	return t.name
}

func (t *_MsgGetMapActionFeedback) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgGetMapActionFeedback) NewMessage() ros.Message {
	m := new(GetMapActionFeedback)
	m.Header = std_msgs.Header{}
	m.Status = actionlib_msgs.GoalStatus{}
	m.Feedback = GetMapFeedback{}
	return m
}

var (
	MsgGetMapActionFeedback = &_MsgGetMapActionFeedback{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======

Header header
actionlib_msgs/GoalStatus status
GetMapFeedback feedback
`,
		"nav_msgs/GetMapActionFeedback",
		"aae20e09065c3809e8a8e87c4c8953fd",
	}
)

type GetMapActionFeedback struct {
	Header   std_msgs.Header           `rosmsg:"header:Header"`
	Status   actionlib_msgs.GoalStatus `rosmsg:"status:GoalStatus"`
	Feedback GetMapFeedback            `rosmsg:"feedback:GetMapFeedback"`
}

func (m *GetMapActionFeedback) GetType() ros.MessageType {
	return MsgGetMapActionFeedback
}

func (m *GetMapActionFeedback) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	if err = m.Status.Serialize(buf); err != nil {
		return err
	}
	if err = m.Feedback.Serialize(buf); err != nil {
		return err
	}
	return err
}

func (m *GetMapActionFeedback) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = m.Status.Deserialize(buf); err != nil {
		return err
	}
	if err = m.Feedback.Deserialize(buf); err != nil {
		return err
	}
	return err
}
// Automatically generated from the message definition "sensor_msgs/JoyFeedbackArray.msg"
package sensor_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgJoyFeedbackArray struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgJoyFeedbackArray) Text() string {
	return t.text
}

func (t *_MsgJoyFeedbackArray) Name() string {
	return t.name
}

func (t *_MsgJoyFeedbackArray) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgJoyFeedbackArray) NewMessage() ros.Message {
	m := new(JoyFeedbackArray)
	m.Array = []JoyFeedback{}
	return m
}

var (
	MsgJoyFeedbackArray = &_MsgJoyFeedbackArray{
		`# This message publishes values for multiple feedback at once. 
JoyFeedback[] array`,
		"sensor_msgs/JoyFeedbackArray",
		"cde5730a895b1fc4dee6f91b754b213d",
	}
)

type JoyFeedbackArray struct {
	Array []JoyFeedback `rosmsg:"array:JoyFeedback[]"`
}

func (m *JoyFeedbackArray) GetType() ros.MessageType {
	return MsgJoyFeedbackArray
}

func (m *JoyFeedbackArray) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Array)))
	for _, e := range m.Array {
		if err = e.Serialize(buf); err != nil {
			return err
		}
	}
	return err
}

func (m *JoyFeedbackArray) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Array = make([]JoyFeedback, int(size))
		for i := 0; i < int(size); i++ {
			if err = m.Array[i].Deserialize(buf); err != nil {
				return err
			}
		}
	}
	return err
}

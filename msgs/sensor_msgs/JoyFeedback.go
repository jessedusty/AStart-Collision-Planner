// Automatically generated from the message definition "sensor_msgs/JoyFeedback.msg"
package sensor_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

const (
	JoyFeedback_TYPE_LED    uint8 = 0
	JoyFeedback_TYPE_RUMBLE uint8 = 1
	JoyFeedback_TYPE_BUZZER uint8 = 2
)

type _MsgJoyFeedback struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgJoyFeedback) Text() string {
	return t.text
}

func (t *_MsgJoyFeedback) Name() string {
	return t.name
}

func (t *_MsgJoyFeedback) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgJoyFeedback) NewMessage() ros.Message {
	m := new(JoyFeedback)
	m.Type = 0
	m.Id = 0
	m.Intensity = 0.0
	return m
}

var (
	MsgJoyFeedback = &_MsgJoyFeedback{
		`# Declare of the type of feedback
uint8 TYPE_LED    = 0
uint8 TYPE_RUMBLE = 1
uint8 TYPE_BUZZER = 2

uint8 type

# This will hold an id number for each type of each feedback.
# Example, the first led would be id=0, the second would be id=1
uint8 id

# Intensity of the feedback, from 0.0 to 1.0, inclusive.  If device is
# actually binary, driver should treat 0<=x<0.5 as off, 0.5<=x<=1 as on.
float32 intensity

`,
		"sensor_msgs/JoyFeedback",
		"f4dcd73460360d98f36e55ee7f2e46f1",
	}
)

type JoyFeedback struct {
	Type      uint8   `rosmsg:"type:uint8"`
	Id        uint8   `rosmsg:"id:uint8"`
	Intensity float32 `rosmsg:"intensity:float32"`
}

func (m *JoyFeedback) GetType() ros.MessageType {
	return MsgJoyFeedback
}

func (m *JoyFeedback) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, m.Type)
	binary.Write(buf, binary.LittleEndian, m.Id)
	binary.Write(buf, binary.LittleEndian, m.Intensity)
	return err
}

func (m *JoyFeedback) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = binary.Read(buf, binary.LittleEndian, &m.Type); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Id); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Intensity); err != nil {
		return err
	}
	return err
}

// Automatically generated from the message definition "sensor_msgs/Joy.msg"
package sensor_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgJoy struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgJoy) Text() string {
	return t.text
}

func (t *_MsgJoy) Name() string {
	return t.name
}

func (t *_MsgJoy) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgJoy) NewMessage() ros.Message {
	m := new(Joy)
	m.Header = std_msgs.Header{}
	m.Axes = []float32{}
	m.Buttons = []int32{}
	return m
}

var (
	MsgJoy = &_MsgJoy{
		`# Reports the state of a joysticks axes and buttons.
Header header           # timestamp in the header is the time the data is received from the joystick
float32[] axes          # the axes measurements from a joystick
int32[] buttons         # the buttons measurements from a joystick 
`,
		"sensor_msgs/Joy",
		"5a9ea5f83505693b71e785041e67a8bb",
	}
)

type Joy struct {
	Header  std_msgs.Header `rosmsg:"header:Header"`
	Axes    []float32       `rosmsg:"axes:float32[]"`
	Buttons []int32         `rosmsg:"buttons:int32[]"`
}

func (m *Joy) GetType() ros.MessageType {
	return MsgJoy
}

func (m *Joy) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Axes)))
	for _, e := range m.Axes {
		binary.Write(buf, binary.LittleEndian, e)
	}
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Buttons)))
	for _, e := range m.Buttons {
		binary.Write(buf, binary.LittleEndian, e)
	}
	return err
}

func (m *Joy) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Axes = make([]float32, int(size))
		for i := 0; i < int(size); i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.Axes[i]); err != nil {
				return err
			}
		}
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Buttons = make([]int32, int(size))
		for i := 0; i < int(size); i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.Buttons[i]); err != nil {
				return err
			}
		}
	}
	return err
}

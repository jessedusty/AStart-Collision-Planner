// Automatically generated from the message definition "sensor_msgs/JointState.msg"
package sensor_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgJointState struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgJointState) Text() string {
	return t.text
}

func (t *_MsgJointState) Name() string {
	return t.name
}

func (t *_MsgJointState) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgJointState) NewMessage() ros.Message {
	m := new(JointState)
	m.Header = std_msgs.Header{}
	m.Name = []string{}
	m.Position = []float64{}
	m.Velocity = []float64{}
	m.Effort = []float64{}
	return m
}

var (
	MsgJointState = &_MsgJointState{
		`# This is a message that holds data to describe the state of a set of torque controlled joints. 
#
# The state of each joint (revolute or prismatic) is defined by:
#  * the position of the joint (rad or m),
#  * the velocity of the joint (rad/s or m/s) and 
#  * the effort that is applied in the joint (Nm or N).
#
# Each joint is uniquely identified by its name
# The header specifies the time at which the joint states were recorded. All the joint states
# in one message have to be recorded at the same time.
#
# This message consists of a multiple arrays, one for each part of the joint state. 
# The goal is to make each of the fields optional. When e.g. your joints have no
# effort associated with them, you can leave the effort array empty. 
#
# All arrays in this message should have the same size, or be empty.
# This is the only way to uniquely associate the joint name with the correct
# states.


Header header

string[] name
float64[] position
float64[] velocity
float64[] effort
`,
		"sensor_msgs/JointState",
		"3066dcd76a6cfaef579bd0f34173e9fd",
	}
)

type JointState struct {
	Header   std_msgs.Header `rosmsg:"header:Header"`
	Name     []string        `rosmsg:"name:string[]"`
	Position []float64       `rosmsg:"position:float64[]"`
	Velocity []float64       `rosmsg:"velocity:float64[]"`
	Effort   []float64       `rosmsg:"effort:float64[]"`
}

func (m *JointState) GetType() ros.MessageType {
	return MsgJointState
}

func (m *JointState) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Name)))
	for _, e := range m.Name {
		binary.Write(buf, binary.LittleEndian, uint32(len([]byte(e))))
		buf.Write([]byte(e))
	}
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Position)))
	for _, e := range m.Position {
		binary.Write(buf, binary.LittleEndian, e)
	}
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Velocity)))
	for _, e := range m.Velocity {
		binary.Write(buf, binary.LittleEndian, e)
	}
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Effort)))
	for _, e := range m.Effort {
		binary.Write(buf, binary.LittleEndian, e)
	}
	return err
}

func (m *JointState) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Name = make([]string, int(size))
		for i := 0; i < int(size); i++ {
			{
				var size uint32
				if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
					return err
				}
				data := make([]byte, int(size))
				if err = binary.Read(buf, binary.LittleEndian, data); err != nil {
					return err
				}
				m.Name[i] = string(data)
			}
		}
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Position = make([]float64, int(size))
		for i := 0; i < int(size); i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.Position[i]); err != nil {
				return err
			}
		}
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Velocity = make([]float64, int(size))
		for i := 0; i < int(size); i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.Velocity[i]); err != nil {
				return err
			}
		}
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Effort = make([]float64, int(size))
		for i := 0; i < int(size); i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.Effort[i]); err != nil {
				return err
			}
		}
	}
	return err
}

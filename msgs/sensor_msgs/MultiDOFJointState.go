// Automatically generated from the message definition "sensor_msgs/MultiDOFJointState.msg"
package sensor_msgs

import (
	"bytes"
	"ee631_midterm/msgs/geometry_msgs"
	"ee631_midterm/msgs/std_msgs"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgMultiDOFJointState struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgMultiDOFJointState) Text() string {
	return t.text
}

func (t *_MsgMultiDOFJointState) Name() string {
	return t.name
}

func (t *_MsgMultiDOFJointState) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgMultiDOFJointState) NewMessage() ros.Message {
	m := new(MultiDOFJointState)
	m.Header = std_msgs.Header{}
	m.JointNames = []string{}
	m.Transforms = []geometry_msgs.Transform{}
	m.Twist = []geometry_msgs.Twist{}
	m.Wrench = []geometry_msgs.Wrench{}
	return m
}

var (
	MsgMultiDOFJointState = &_MsgMultiDOFJointState{
		`# Representation of state for joints with multiple degrees of freedom, 
# following the structure of JointState.
#
# It is assumed that a joint in a system corresponds to a transform that gets applied 
# along the kinematic chain. For example, a planar joint (as in URDF) is 3DOF (x, y, yaw)
# and those 3DOF can be expressed as a transformation matrix, and that transformation
# matrix can be converted back to (x, y, yaw)
#
# Each joint is uniquely identified by its name
# The header specifies the time at which the joint states were recorded. All the joint states
# in one message have to be recorded at the same time.
#
# This message consists of a multiple arrays, one for each part of the joint state. 
# The goal is to make each of the fields optional. When e.g. your joints have no
# wrench associated with them, you can leave the wrench array empty. 
#
# All arrays in this message should have the same size, or be empty.
# This is the only way to uniquely associate the joint name with the correct
# states.

Header header

string[] joint_names
geometry_msgs/Transform[] transforms
geometry_msgs/Twist[] twist
geometry_msgs/Wrench[] wrench
`,
		"sensor_msgs/MultiDOFJointState",
		"690f272f0640d2631c305eeb8301e59d",
	}
)

type MultiDOFJointState struct {
	Header     std_msgs.Header           `rosmsg:"header:Header"`
	JointNames []string                  `rosmsg:"joint_names:string[]"`
	Transforms []geometry_msgs.Transform `rosmsg:"transforms:Transform[]"`
	Twist      []geometry_msgs.Twist     `rosmsg:"twist:Twist[]"`
	Wrench     []geometry_msgs.Wrench    `rosmsg:"wrench:Wrench[]"`
}

func (m *MultiDOFJointState) GetType() ros.MessageType {
	return MsgMultiDOFJointState
}

func (m *MultiDOFJointState) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, uint32(len(m.JointNames)))
	for _, e := range m.JointNames {
		binary.Write(buf, binary.LittleEndian, uint32(len([]byte(e))))
		buf.Write([]byte(e))
	}
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Transforms)))
	for _, e := range m.Transforms {
		if err = e.Serialize(buf); err != nil {
			return err
		}
	}
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Twist)))
	for _, e := range m.Twist {
		if err = e.Serialize(buf); err != nil {
			return err
		}
	}
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Wrench)))
	for _, e := range m.Wrench {
		if err = e.Serialize(buf); err != nil {
			return err
		}
	}
	return err
}

func (m *MultiDOFJointState) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.JointNames = make([]string, int(size))
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
				m.JointNames[i] = string(data)
			}
		}
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Transforms = make([]geometry_msgs.Transform, int(size))
		for i := 0; i < int(size); i++ {
			if err = m.Transforms[i].Deserialize(buf); err != nil {
				return err
			}
		}
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Twist = make([]geometry_msgs.Twist, int(size))
		for i := 0; i < int(size); i++ {
			if err = m.Twist[i].Deserialize(buf); err != nil {
				return err
			}
		}
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Wrench = make([]geometry_msgs.Wrench, int(size))
		for i := 0; i < int(size); i++ {
			if err = m.Wrench[i].Deserialize(buf); err != nil {
				return err
			}
		}
	}
	return err
}

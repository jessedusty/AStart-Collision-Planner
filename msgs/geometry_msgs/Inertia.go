// Automatically generated from the message definition "geometry_msgs/Inertia.msg"
package geometry_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgInertia struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgInertia) Text() string {
	return t.text
}

func (t *_MsgInertia) Name() string {
	return t.name
}

func (t *_MsgInertia) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgInertia) NewMessage() ros.Message {
	m := new(Inertia)
	m.M = 0.0
	m.Com = Vector3{}
	m.Ixx = 0.0
	m.Ixy = 0.0
	m.Ixz = 0.0
	m.Iyy = 0.0
	m.Iyz = 0.0
	m.Izz = 0.0
	return m
}

var (
	MsgInertia = &_MsgInertia{
		`# Mass [kg]
float64 m

# Center of mass [m]
geometry_msgs/Vector3 com

# Inertia Tensor [kg-m^2]
#     | ixx ixy ixz |
# I = | ixy iyy iyz |
#     | ixz iyz izz |
float64 ixx
float64 ixy
float64 ixz
float64 iyy
float64 iyz
float64 izz
`,
		"geometry_msgs/Inertia",
		"1d26e4bb6c83ff141c5cf0d883c2b0fe",
	}
)

type Inertia struct {
	M   float64 `rosmsg:"m:float64"`
	Com Vector3 `rosmsg:"com:Vector3"`
	Ixx float64 `rosmsg:"ixx:float64"`
	Ixy float64 `rosmsg:"ixy:float64"`
	Ixz float64 `rosmsg:"ixz:float64"`
	Iyy float64 `rosmsg:"iyy:float64"`
	Iyz float64 `rosmsg:"iyz:float64"`
	Izz float64 `rosmsg:"izz:float64"`
}

func (m *Inertia) GetType() ros.MessageType {
	return MsgInertia
}

func (m *Inertia) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, m.M)
	if err = m.Com.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, m.Ixx)
	binary.Write(buf, binary.LittleEndian, m.Ixy)
	binary.Write(buf, binary.LittleEndian, m.Ixz)
	binary.Write(buf, binary.LittleEndian, m.Iyy)
	binary.Write(buf, binary.LittleEndian, m.Iyz)
	binary.Write(buf, binary.LittleEndian, m.Izz)
	return err
}

func (m *Inertia) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = binary.Read(buf, binary.LittleEndian, &m.M); err != nil {
		return err
	}
	if err = m.Com.Deserialize(buf); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Ixx); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Ixy); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Ixz); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Iyy); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Iyz); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Izz); err != nil {
		return err
	}
	return err
}

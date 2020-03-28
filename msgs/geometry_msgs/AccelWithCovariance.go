// Automatically generated from the message definition "geometry_msgs/AccelWithCovariance.msg"
package geometry_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgAccelWithCovariance struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgAccelWithCovariance) Text() string {
	return t.text
}

func (t *_MsgAccelWithCovariance) Name() string {
	return t.name
}

func (t *_MsgAccelWithCovariance) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgAccelWithCovariance) NewMessage() ros.Message {
	m := new(AccelWithCovariance)
	m.Accel = Accel{}
	for i := 0; i < 36; i++ {
		m.Covariance[i] = 0.0
	}
	return m
}

var (
	MsgAccelWithCovariance = &_MsgAccelWithCovariance{
		`# This expresses acceleration in free space with uncertainty.

Accel accel

# Row-major representation of the 6x6 covariance matrix
# The orientation parameters use a fixed-axis representation.
# In order, the parameters are:
# (x, y, z, rotation about X axis, rotation about Y axis, rotation about Z axis)
float64[36] covariance
`,
		"geometry_msgs/AccelWithCovariance",
		"ad5a718d699c6be72a02b8d6a139f334",
	}
)

type AccelWithCovariance struct {
	Accel      Accel       `rosmsg:"accel:Accel"`
	Covariance [36]float64 `rosmsg:"covariance:float64[36]"`
}

func (m *AccelWithCovariance) GetType() ros.MessageType {
	return MsgAccelWithCovariance
}

func (m *AccelWithCovariance) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Accel.Serialize(buf); err != nil {
		return err
	}
	for _, e := range m.Covariance {
		binary.Write(buf, binary.LittleEndian, e)
	}
	return err
}

func (m *AccelWithCovariance) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Accel.Deserialize(buf); err != nil {
		return err
	}
	{
		for i := 0; i < 36; i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.Covariance[i]); err != nil {
				return err
			}
		}
	}
	return err
}

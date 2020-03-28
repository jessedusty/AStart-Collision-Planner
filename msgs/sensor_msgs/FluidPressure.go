// Automatically generated from the message definition "sensor_msgs/FluidPressure.msg"
package sensor_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgFluidPressure struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgFluidPressure) Text() string {
	return t.text
}

func (t *_MsgFluidPressure) Name() string {
	return t.name
}

func (t *_MsgFluidPressure) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgFluidPressure) NewMessage() ros.Message {
	m := new(FluidPressure)
	m.Header = std_msgs.Header{}
	m.FluidPressure = 0.0
	m.Variance = 0.0
	return m
}

var (
	MsgFluidPressure = &_MsgFluidPressure{
		` # Single pressure reading.  This message is appropriate for measuring the
 # pressure inside of a fluid (air, water, etc).  This also includes
 # atmospheric or barometric pressure.

 # This message is not appropriate for force/pressure contact sensors.

 Header header           # timestamp of the measurement
                         # frame_id is the location of the pressure sensor

 float64 fluid_pressure  # Absolute pressure reading in Pascals.

 float64 variance        # 0 is interpreted as variance unknown`,
		"sensor_msgs/FluidPressure",
		"804dc5cea1c5306d6a2eb80b9833befe",
	}
)

type FluidPressure struct {
	Header        std_msgs.Header `rosmsg:"header:Header"`
	FluidPressure float64         `rosmsg:"fluid_pressure:float64"`
	Variance      float64         `rosmsg:"variance:float64"`
}

func (m *FluidPressure) GetType() ros.MessageType {
	return MsgFluidPressure
}

func (m *FluidPressure) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, m.FluidPressure)
	binary.Write(buf, binary.LittleEndian, m.Variance)
	return err
}

func (m *FluidPressure) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.FluidPressure); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Variance); err != nil {
		return err
	}
	return err
}

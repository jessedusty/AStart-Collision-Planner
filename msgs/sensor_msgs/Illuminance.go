// Automatically generated from the message definition "sensor_msgs/Illuminance.msg"
package sensor_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgIlluminance struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgIlluminance) Text() string {
	return t.text
}

func (t *_MsgIlluminance) Name() string {
	return t.name
}

func (t *_MsgIlluminance) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgIlluminance) NewMessage() ros.Message {
	m := new(Illuminance)
	m.Header = std_msgs.Header{}
	m.Illuminance = 0.0
	m.Variance = 0.0
	return m
}

var (
	MsgIlluminance = &_MsgIlluminance{
		` # Single photometric illuminance measurement.  Light should be assumed to be
 # measured along the sensor's x-axis (the area of detection is the y-z plane).
 # The illuminance should have a 0 or positive value and be received with
 # the sensor's +X axis pointing toward the light source.

 # Photometric illuminance is the measure of the human eye's sensitivity of the
 # intensity of light encountering or passing through a surface.

 # All other Photometric and Radiometric measurements should
 # not use this message.
 # This message cannot represent:
 # Luminous intensity (candela/light source output)
 # Luminance (nits/light output per area)
 # Irradiance (watt/area), etc.

 Header header           # timestamp is the time the illuminance was measured
                         # frame_id is the location and direction of the reading

 float64 illuminance     # Measurement of the Photometric Illuminance in Lux.

 float64 variance        # 0 is interpreted as variance unknown`,
		"sensor_msgs/Illuminance",
		"8cf5febb0952fca9d650c3d11a81a188",
	}
)

type Illuminance struct {
	Header      std_msgs.Header `rosmsg:"header:Header"`
	Illuminance float64         `rosmsg:"illuminance:float64"`
	Variance    float64         `rosmsg:"variance:float64"`
}

func (m *Illuminance) GetType() ros.MessageType {
	return MsgIlluminance
}

func (m *Illuminance) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, m.Illuminance)
	binary.Write(buf, binary.LittleEndian, m.Variance)
	return err
}

func (m *Illuminance) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Illuminance); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Variance); err != nil {
		return err
	}
	return err
}

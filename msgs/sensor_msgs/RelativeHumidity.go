// Automatically generated from the message definition "sensor_msgs/RelativeHumidity.msg"
package sensor_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgRelativeHumidity struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgRelativeHumidity) Text() string {
	return t.text
}

func (t *_MsgRelativeHumidity) Name() string {
	return t.name
}

func (t *_MsgRelativeHumidity) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgRelativeHumidity) NewMessage() ros.Message {
	m := new(RelativeHumidity)
	m.Header = std_msgs.Header{}
	m.RelativeHumidity = 0.0
	m.Variance = 0.0
	return m
}

var (
	MsgRelativeHumidity = &_MsgRelativeHumidity{
		` # Single reading from a relative humidity sensor.  Defines the ratio of partial
 # pressure of water vapor to the saturated vapor pressure at a temperature.

 Header header             # timestamp of the measurement
                           # frame_id is the location of the humidity sensor

 float64 relative_humidity # Expression of the relative humidity
                           # from 0.0 to 1.0.
                           # 0.0 is no partial pressure of water vapor
                           # 1.0 represents partial pressure of saturation

 float64 variance          # 0 is interpreted as variance unknown`,
		"sensor_msgs/RelativeHumidity",
		"8730015b05955b7e992ce29a2678d90f",
	}
)

type RelativeHumidity struct {
	Header           std_msgs.Header `rosmsg:"header:Header"`
	RelativeHumidity float64         `rosmsg:"relative_humidity:float64"`
	Variance         float64         `rosmsg:"variance:float64"`
}

func (m *RelativeHumidity) GetType() ros.MessageType {
	return MsgRelativeHumidity
}

func (m *RelativeHumidity) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, m.RelativeHumidity)
	binary.Write(buf, binary.LittleEndian, m.Variance)
	return err
}

func (m *RelativeHumidity) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.RelativeHumidity); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Variance); err != nil {
		return err
	}
	return err
}

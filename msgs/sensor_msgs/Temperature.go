// Automatically generated from the message definition "sensor_msgs/Temperature.msg"
package sensor_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgTemperature struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgTemperature) Text() string {
	return t.text
}

func (t *_MsgTemperature) Name() string {
	return t.name
}

func (t *_MsgTemperature) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgTemperature) NewMessage() ros.Message {
	m := new(Temperature)
	m.Header = std_msgs.Header{}
	m.Temperature = 0.0
	m.Variance = 0.0
	return m
}

var (
	MsgTemperature = &_MsgTemperature{
		` # Single temperature reading.

 Header header           # timestamp is the time the temperature was measured
                         # frame_id is the location of the temperature reading

 float64 temperature     # Measurement of the Temperature in Degrees Celsius

 float64 variance        # 0 is interpreted as variance unknown`,
		"sensor_msgs/Temperature",
		"ff71b307acdbe7c871a5a6d7ed359100",
	}
)

type Temperature struct {
	Header      std_msgs.Header `rosmsg:"header:Header"`
	Temperature float64         `rosmsg:"temperature:float64"`
	Variance    float64         `rosmsg:"variance:float64"`
}

func (m *Temperature) GetType() ros.MessageType {
	return MsgTemperature
}

func (m *Temperature) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, m.Temperature)
	binary.Write(buf, binary.LittleEndian, m.Variance)
	return err
}

func (m *Temperature) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Temperature); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Variance); err != nil {
		return err
	}
	return err
}

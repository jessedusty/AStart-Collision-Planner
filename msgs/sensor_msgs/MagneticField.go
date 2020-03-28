// Automatically generated from the message definition "sensor_msgs/MagneticField.msg"
package sensor_msgs

import (
	"bytes"
	"ee631_midterm/msgs/geometry_msgs"
	"ee631_midterm/msgs/std_msgs"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgMagneticField struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgMagneticField) Text() string {
	return t.text
}

func (t *_MsgMagneticField) Name() string {
	return t.name
}

func (t *_MsgMagneticField) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgMagneticField) NewMessage() ros.Message {
	m := new(MagneticField)
	m.Header = std_msgs.Header{}
	m.MagneticField = geometry_msgs.Vector3{}
	for i := 0; i < 9; i++ {
		m.MagneticFieldCovariance[i] = 0.0
	}
	return m
}

var (
	MsgMagneticField = &_MsgMagneticField{
		` # Measurement of the Magnetic Field vector at a specific location.

 # If the covariance of the measurement is known, it should be filled in
 # (if all you know is the variance of each measurement, e.g. from the datasheet,
 #just put those along the diagonal)
 # A covariance matrix of all zeros will be interpreted as "covariance unknown",
 # and to use the data a covariance will have to be assumed or gotten from some
 # other source


 Header header                        # timestamp is the time the
                                      # field was measured
                                      # frame_id is the location and orientation
                                      # of the field measurement

 geometry_msgs/Vector3 magnetic_field # x, y, and z components of the
                                      # field vector in Tesla
                                      # If your sensor does not output 3 axes,
                                      # put NaNs in the components not reported.

 float64[9] magnetic_field_covariance # Row major about x, y, z axes
                                      # 0 is interpreted as variance unknown`,
		"sensor_msgs/MagneticField",
		"2f3b0b43eed0c9501de0fa3ff89a45aa",
	}
)

type MagneticField struct {
	Header                  std_msgs.Header       `rosmsg:"header:Header"`
	MagneticField           geometry_msgs.Vector3 `rosmsg:"magnetic_field:Vector3"`
	MagneticFieldCovariance [9]float64            `rosmsg:"magnetic_field_covariance:float64[9]"`
}

func (m *MagneticField) GetType() ros.MessageType {
	return MsgMagneticField
}

func (m *MagneticField) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	if err = m.MagneticField.Serialize(buf); err != nil {
		return err
	}
	for _, e := range m.MagneticFieldCovariance {
		binary.Write(buf, binary.LittleEndian, e)
	}
	return err
}

func (m *MagneticField) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = m.MagneticField.Deserialize(buf); err != nil {
		return err
	}
	{
		for i := 0; i < 9; i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.MagneticFieldCovariance[i]); err != nil {
				return err
			}
		}
	}
	return err
}

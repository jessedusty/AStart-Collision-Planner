// Automatically generated from the message definition "sensor_msgs/Imu.msg"
package sensor_msgs

import (
	"bytes"
	"ee631_midterm/msgs/geometry_msgs"
	"ee631_midterm/msgs/std_msgs"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgImu struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgImu) Text() string {
	return t.text
}

func (t *_MsgImu) Name() string {
	return t.name
}

func (t *_MsgImu) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgImu) NewMessage() ros.Message {
	m := new(Imu)
	m.Header = std_msgs.Header{}
	m.Orientation = geometry_msgs.Quaternion{}
	for i := 0; i < 9; i++ {
		m.OrientationCovariance[i] = 0.0
	}
	m.AngularVelocity = geometry_msgs.Vector3{}
	for i := 0; i < 9; i++ {
		m.AngularVelocityCovariance[i] = 0.0
	}
	m.LinearAcceleration = geometry_msgs.Vector3{}
	for i := 0; i < 9; i++ {
		m.LinearAccelerationCovariance[i] = 0.0
	}
	return m
}

var (
	MsgImu = &_MsgImu{
		`# This is a message to hold data from an IMU (Inertial Measurement Unit)
#
# Accelerations should be in m/s^2 (not in g's), and rotational velocity should be in rad/sec
#
# If the covariance of the measurement is known, it should be filled in (if all you know is the 
# variance of each measurement, e.g. from the datasheet, just put those along the diagonal)
# A covariance matrix of all zeros will be interpreted as "covariance unknown", and to use the
# data a covariance will have to be assumed or gotten from some other source
#
# If you have no estimate for one of the data elements (e.g. your IMU doesn't produce an orientation 
# estimate), please set element 0 of the associated covariance matrix to -1
# If you are interpreting this message, please check for a value of -1 in the first element of each 
# covariance matrix, and disregard the associated estimate.

Header header

geometry_msgs/Quaternion orientation
float64[9] orientation_covariance # Row major about x, y, z axes

geometry_msgs/Vector3 angular_velocity
float64[9] angular_velocity_covariance # Row major about x, y, z axes

geometry_msgs/Vector3 linear_acceleration
float64[9] linear_acceleration_covariance # Row major x, y z 
`,
		"sensor_msgs/Imu",
		"6a62c6daae103f4ff57a132d6f95cec2",
	}
)

type Imu struct {
	Header                       std_msgs.Header          `rosmsg:"header:Header"`
	Orientation                  geometry_msgs.Quaternion `rosmsg:"orientation:Quaternion"`
	OrientationCovariance        [9]float64               `rosmsg:"orientation_covariance:float64[9]"`
	AngularVelocity              geometry_msgs.Vector3    `rosmsg:"angular_velocity:Vector3"`
	AngularVelocityCovariance    [9]float64               `rosmsg:"angular_velocity_covariance:float64[9]"`
	LinearAcceleration           geometry_msgs.Vector3    `rosmsg:"linear_acceleration:Vector3"`
	LinearAccelerationCovariance [9]float64               `rosmsg:"linear_acceleration_covariance:float64[9]"`
}

func (m *Imu) GetType() ros.MessageType {
	return MsgImu
}

func (m *Imu) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	if err = m.Orientation.Serialize(buf); err != nil {
		return err
	}
	for _, e := range m.OrientationCovariance {
		binary.Write(buf, binary.LittleEndian, e)
	}
	if err = m.AngularVelocity.Serialize(buf); err != nil {
		return err
	}
	for _, e := range m.AngularVelocityCovariance {
		binary.Write(buf, binary.LittleEndian, e)
	}
	if err = m.LinearAcceleration.Serialize(buf); err != nil {
		return err
	}
	for _, e := range m.LinearAccelerationCovariance {
		binary.Write(buf, binary.LittleEndian, e)
	}
	return err
}

func (m *Imu) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = m.Orientation.Deserialize(buf); err != nil {
		return err
	}
	{
		for i := 0; i < 9; i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.OrientationCovariance[i]); err != nil {
				return err
			}
		}
	}
	if err = m.AngularVelocity.Deserialize(buf); err != nil {
		return err
	}
	{
		for i := 0; i < 9; i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.AngularVelocityCovariance[i]); err != nil {
				return err
			}
		}
	}
	if err = m.LinearAcceleration.Deserialize(buf); err != nil {
		return err
	}
	{
		for i := 0; i < 9; i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.LinearAccelerationCovariance[i]); err != nil {
				return err
			}
		}
	}
	return err
}

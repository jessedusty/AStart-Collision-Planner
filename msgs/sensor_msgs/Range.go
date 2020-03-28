// Automatically generated from the message definition "sensor_msgs/Range.msg"
package sensor_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

const (
	Range_ULTRASOUND uint8 = 0
	Range_INFRARED   uint8 = 1
)

type _MsgRange struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgRange) Text() string {
	return t.text
}

func (t *_MsgRange) Name() string {
	return t.name
}

func (t *_MsgRange) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgRange) NewMessage() ros.Message {
	m := new(Range)
	m.Header = std_msgs.Header{}
	m.RadiationType = 0
	m.FieldOfView = 0.0
	m.MinRange = 0.0
	m.MaxRange = 0.0
	m.Range = 0.0
	return m
}

var (
	MsgRange = &_MsgRange{
		`# Single range reading from an active ranger that emits energy and reports
# one range reading that is valid along an arc at the distance measured. 
# This message is  not appropriate for laser scanners. See the LaserScan
# message if you are working with a laser scanner.

# This message also can represent a fixed-distance (binary) ranger.  This
# sensor will have min_range===max_range===distance of detection.
# These sensors follow REP 117 and will output -Inf if the object is detected
# and +Inf if the object is outside of the detection range.

Header header           # timestamp in the header is the time the ranger
                        # returned the distance reading

# Radiation type enums
# If you want a value added to this list, send an email to the ros-users list
uint8 ULTRASOUND=0
uint8 INFRARED=1

uint8 radiation_type    # the type of radiation used by the sensor
                        # (sound, IR, etc) [enum]

float32 field_of_view   # the size of the arc that the distance reading is
                        # valid for [rad]
                        # the object causing the range reading may have
                        # been anywhere within -field_of_view/2 and
                        # field_of_view/2 at the measured range. 
                        # 0 angle corresponds to the x-axis of the sensor.

float32 min_range       # minimum range value [m]
float32 max_range       # maximum range value [m]
                        # Fixed distance rangers require min_range==max_range

float32 range           # range data [m]
                        # (Note: values < range_min or > range_max
                        # should be discarded)
                        # Fixed distance rangers only output -Inf or +Inf.
                        # -Inf represents a detection within fixed distance.
                        # (Detection too close to the sensor to quantify)
                        # +Inf represents no detection within the fixed distance.
                        # (Object out of range)`,
		"sensor_msgs/Range",
		"c005c34273dc426c67a020a87bc24148",
	}
)

type Range struct {
	Header        std_msgs.Header `rosmsg:"header:Header"`
	RadiationType uint8           `rosmsg:"radiation_type:uint8"`
	FieldOfView   float32         `rosmsg:"field_of_view:float32"`
	MinRange      float32         `rosmsg:"min_range:float32"`
	MaxRange      float32         `rosmsg:"max_range:float32"`
	Range         float32         `rosmsg:"range:float32"`
}

func (m *Range) GetType() ros.MessageType {
	return MsgRange
}

func (m *Range) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, m.RadiationType)
	binary.Write(buf, binary.LittleEndian, m.FieldOfView)
	binary.Write(buf, binary.LittleEndian, m.MinRange)
	binary.Write(buf, binary.LittleEndian, m.MaxRange)
	binary.Write(buf, binary.LittleEndian, m.Range)
	return err
}

func (m *Range) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.RadiationType); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.FieldOfView); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.MinRange); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.MaxRange); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Range); err != nil {
		return err
	}
	return err
}

// Automatically generated from the message definition "sensor_msgs/MultiEchoLaserScan.msg"
package sensor_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgMultiEchoLaserScan struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgMultiEchoLaserScan) Text() string {
	return t.text
}

func (t *_MsgMultiEchoLaserScan) Name() string {
	return t.name
}

func (t *_MsgMultiEchoLaserScan) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgMultiEchoLaserScan) NewMessage() ros.Message {
	m := new(MultiEchoLaserScan)
	m.Header = std_msgs.Header{}
	m.AngleMin = 0.0
	m.AngleMax = 0.0
	m.AngleIncrement = 0.0
	m.TimeIncrement = 0.0
	m.ScanTime = 0.0
	m.RangeMin = 0.0
	m.RangeMax = 0.0
	m.Ranges = []LaserEcho{}
	m.Intensities = []LaserEcho{}
	return m
}

var (
	MsgMultiEchoLaserScan = &_MsgMultiEchoLaserScan{
		`# Single scan from a multi-echo planar laser range-finder
#
# If you have another ranging device with different behavior (e.g. a sonar
# array), please find or create a different message, since applications
# will make fairly laser-specific assumptions about this data

Header header            # timestamp in the header is the acquisition time of 
                         # the first ray in the scan.
                         #
                         # in frame frame_id, angles are measured around 
                         # the positive Z axis (counterclockwise, if Z is up)
                         # with zero angle being forward along the x axis
                         
float32 angle_min        # start angle of the scan [rad]
float32 angle_max        # end angle of the scan [rad]
float32 angle_increment  # angular distance between measurements [rad]

float32 time_increment   # time between measurements [seconds] - if your scanner
                         # is moving, this will be used in interpolating position
                         # of 3d points
float32 scan_time        # time between scans [seconds]

float32 range_min        # minimum range value [m]
float32 range_max        # maximum range value [m]

LaserEcho[] ranges       # range data [m] (Note: NaNs, values < range_min or > range_max should be discarded)
                         # +Inf measurements are out of range
                         # -Inf measurements are too close to determine exact distance.
LaserEcho[] intensities  # intensity data [device-specific units].  If your
                         # device does not provide intensities, please leave
                         # the array empty.`,
		"sensor_msgs/MultiEchoLaserScan",
		"6fefb0c6da89d7c8abe4b339f5c2f8fb",
	}
)

type MultiEchoLaserScan struct {
	Header         std_msgs.Header `rosmsg:"header:Header"`
	AngleMin       float32         `rosmsg:"angle_min:float32"`
	AngleMax       float32         `rosmsg:"angle_max:float32"`
	AngleIncrement float32         `rosmsg:"angle_increment:float32"`
	TimeIncrement  float32         `rosmsg:"time_increment:float32"`
	ScanTime       float32         `rosmsg:"scan_time:float32"`
	RangeMin       float32         `rosmsg:"range_min:float32"`
	RangeMax       float32         `rosmsg:"range_max:float32"`
	Ranges         []LaserEcho     `rosmsg:"ranges:LaserEcho[]"`
	Intensities    []LaserEcho     `rosmsg:"intensities:LaserEcho[]"`
}

func (m *MultiEchoLaserScan) GetType() ros.MessageType {
	return MsgMultiEchoLaserScan
}

func (m *MultiEchoLaserScan) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, m.AngleMin)
	binary.Write(buf, binary.LittleEndian, m.AngleMax)
	binary.Write(buf, binary.LittleEndian, m.AngleIncrement)
	binary.Write(buf, binary.LittleEndian, m.TimeIncrement)
	binary.Write(buf, binary.LittleEndian, m.ScanTime)
	binary.Write(buf, binary.LittleEndian, m.RangeMin)
	binary.Write(buf, binary.LittleEndian, m.RangeMax)
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Ranges)))
	for _, e := range m.Ranges {
		if err = e.Serialize(buf); err != nil {
			return err
		}
	}
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Intensities)))
	for _, e := range m.Intensities {
		if err = e.Serialize(buf); err != nil {
			return err
		}
	}
	return err
}

func (m *MultiEchoLaserScan) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.AngleMin); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.AngleMax); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.AngleIncrement); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.TimeIncrement); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.ScanTime); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.RangeMin); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.RangeMax); err != nil {
		return err
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Ranges = make([]LaserEcho, int(size))
		for i := 0; i < int(size); i++ {
			if err = m.Ranges[i].Deserialize(buf); err != nil {
				return err
			}
		}
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Intensities = make([]LaserEcho, int(size))
		for i := 0; i < int(size); i++ {
			if err = m.Intensities[i].Deserialize(buf); err != nil {
				return err
			}
		}
	}
	return err
}

// Automatically generated from the message definition "sensor_msgs/Image.msg"
package sensor_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgImage struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgImage) Text() string {
	return t.text
}

func (t *_MsgImage) Name() string {
	return t.name
}

func (t *_MsgImage) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgImage) NewMessage() ros.Message {
	m := new(Image)
	m.Header = std_msgs.Header{}
	m.Height = 0
	m.Width = 0
	m.Encoding = ""
	m.IsBigendian = 0
	m.Step = 0
	m.Data = []uint8{}
	return m
}

var (
	MsgImage = &_MsgImage{
		`# This message contains an uncompressed image
# (0, 0) is at top-left corner of image
#

Header header        # Header timestamp should be acquisition time of image
                     # Header frame_id should be optical frame of camera
                     # origin of frame should be optical center of camera
                     # +x should point to the right in the image
                     # +y should point down in the image
                     # +z should point into to plane of the image
                     # If the frame_id here and the frame_id of the CameraInfo
                     # message associated with the image conflict
                     # the behavior is undefined

uint32 height         # image height, that is, number of rows
uint32 width          # image width, that is, number of columns

# The legal values for encoding are in file src/image_encodings.cpp
# If you want to standardize a new string format, join
# ros-users@lists.sourceforge.net and send an email proposing a new encoding.

string encoding       # Encoding of pixels -- channel meaning, ordering, size
                      # taken from the list of strings in include/sensor_msgs/image_encodings.h

uint8 is_bigendian    # is this data bigendian?
uint32 step           # Full row length in bytes
uint8[] data          # actual matrix data, size is (step * rows)
`,
		"sensor_msgs/Image",
		"060021388200f6f0f447d0fcd9c64743",
	}
)

type Image struct {
	Header      std_msgs.Header `rosmsg:"header:Header"`
	Height      uint32          `rosmsg:"height:uint32"`
	Width       uint32          `rosmsg:"width:uint32"`
	Encoding    string          `rosmsg:"encoding:string"`
	IsBigendian uint8           `rosmsg:"is_bigendian:uint8"`
	Step        uint32          `rosmsg:"step:uint32"`
	Data        []uint8         `rosmsg:"data:uint8[]"`
}

func (m *Image) GetType() ros.MessageType {
	return MsgImage
}

func (m *Image) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, m.Height)
	binary.Write(buf, binary.LittleEndian, m.Width)
	binary.Write(buf, binary.LittleEndian, uint32(len([]byte(m.Encoding))))
	buf.Write([]byte(m.Encoding))
	binary.Write(buf, binary.LittleEndian, m.IsBigendian)
	binary.Write(buf, binary.LittleEndian, m.Step)
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Data)))
	for _, e := range m.Data {
		binary.Write(buf, binary.LittleEndian, e)
	}
	return err
}

func (m *Image) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Height); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Width); err != nil {
		return err
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		data := make([]byte, int(size))
		if err = binary.Read(buf, binary.LittleEndian, data); err != nil {
			return err
		}
		m.Encoding = string(data)
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.IsBigendian); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Step); err != nil {
		return err
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Data = make([]uint8, int(size))
		for i := 0; i < int(size); i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.Data[i]); err != nil {
				return err
			}
		}
	}
	return err
}

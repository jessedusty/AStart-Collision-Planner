// Automatically generated from the message definition "sensor_msgs/CompressedImage.msg"
package sensor_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgCompressedImage struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgCompressedImage) Text() string {
	return t.text
}

func (t *_MsgCompressedImage) Name() string {
	return t.name
}

func (t *_MsgCompressedImage) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgCompressedImage) NewMessage() ros.Message {
	m := new(CompressedImage)
	m.Header = std_msgs.Header{}
	m.Format = ""
	m.Data = []uint8{}
	return m
}

var (
	MsgCompressedImage = &_MsgCompressedImage{
		`# This message contains a compressed image

Header header        # Header timestamp should be acquisition time of image
                     # Header frame_id should be optical frame of camera
                     # origin of frame should be optical center of camera
                     # +x should point to the right in the image
                     # +y should point down in the image
                     # +z should point into to plane of the image

string format        # Specifies the format of the data
                     #   Acceptable values:
                     #     jpeg, png
uint8[] data         # Compressed image buffer
`,
		"sensor_msgs/CompressedImage",
		"8f7a12909da2c9d3332d540a0977563f",
	}
)

type CompressedImage struct {
	Header std_msgs.Header `rosmsg:"header:Header"`
	Format string          `rosmsg:"format:string"`
	Data   []uint8         `rosmsg:"data:uint8[]"`
}

func (m *CompressedImage) GetType() ros.MessageType {
	return MsgCompressedImage
}

func (m *CompressedImage) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, uint32(len([]byte(m.Format))))
	buf.Write([]byte(m.Format))
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Data)))
	for _, e := range m.Data {
		binary.Write(buf, binary.LittleEndian, e)
	}
	return err
}

func (m *CompressedImage) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
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
		m.Format = string(data)
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

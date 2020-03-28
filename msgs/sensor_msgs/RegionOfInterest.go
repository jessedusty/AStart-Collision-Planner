// Automatically generated from the message definition "sensor_msgs/RegionOfInterest.msg"
package sensor_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgRegionOfInterest struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgRegionOfInterest) Text() string {
	return t.text
}

func (t *_MsgRegionOfInterest) Name() string {
	return t.name
}

func (t *_MsgRegionOfInterest) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgRegionOfInterest) NewMessage() ros.Message {
	m := new(RegionOfInterest)
	m.XOffset = 0
	m.YOffset = 0
	m.Height = 0
	m.Width = 0
	m.DoRectify = false
	return m
}

var (
	MsgRegionOfInterest = &_MsgRegionOfInterest{
		`# This message is used to specify a region of interest within an image.
#
# When used to specify the ROI setting of the camera when the image was
# taken, the height and width fields should either match the height and
# width fields for the associated image; or height = width = 0
# indicates that the full resolution image was captured.

uint32 x_offset  # Leftmost pixel of the ROI
                 # (0 if the ROI includes the left edge of the image)
uint32 y_offset  # Topmost pixel of the ROI
                 # (0 if the ROI includes the top edge of the image)
uint32 height    # Height of ROI
uint32 width     # Width of ROI

# True if a distinct rectified ROI should be calculated from the "raw"
# ROI in this message. Typically this should be False if the full image
# is captured (ROI not used), and True if a subwindow is captured (ROI
# used).
bool do_rectify
`,
		"sensor_msgs/RegionOfInterest",
		"bdb633039d588fcccb441a4d43ccfe09",
	}
)

type RegionOfInterest struct {
	XOffset   uint32 `rosmsg:"x_offset:uint32"`
	YOffset   uint32 `rosmsg:"y_offset:uint32"`
	Height    uint32 `rosmsg:"height:uint32"`
	Width     uint32 `rosmsg:"width:uint32"`
	DoRectify bool   `rosmsg:"do_rectify:bool"`
}

func (m *RegionOfInterest) GetType() ros.MessageType {
	return MsgRegionOfInterest
}

func (m *RegionOfInterest) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, m.XOffset)
	binary.Write(buf, binary.LittleEndian, m.YOffset)
	binary.Write(buf, binary.LittleEndian, m.Height)
	binary.Write(buf, binary.LittleEndian, m.Width)
	binary.Write(buf, binary.LittleEndian, m.DoRectify)
	return err
}

func (m *RegionOfInterest) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = binary.Read(buf, binary.LittleEndian, &m.XOffset); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.YOffset); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Height); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Width); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.DoRectify); err != nil {
		return err
	}
	return err
}

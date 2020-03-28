// Automatically generated from the message definition "map_msgs/OccupancyGridUpdate.msg"
package map_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgOccupancyGridUpdate struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgOccupancyGridUpdate) Text() string {
	return t.text
}

func (t *_MsgOccupancyGridUpdate) Name() string {
	return t.name
}

func (t *_MsgOccupancyGridUpdate) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgOccupancyGridUpdate) NewMessage() ros.Message {
	m := new(OccupancyGridUpdate)
	m.Header = std_msgs.Header{}
	m.X = 0
	m.Y = 0
	m.Width = 0
	m.Height = 0
	m.Data = []int8{}
	return m
}

var (
	MsgOccupancyGridUpdate = &_MsgOccupancyGridUpdate{
		`Header header
int32 x
int32 y
uint32 width
uint32 height
int8[] data
`,
		"map_msgs/OccupancyGridUpdate",
		"b295be292b335c34718bd939deebe1c9",
	}
)

type OccupancyGridUpdate struct {
	Header std_msgs.Header `rosmsg:"header:Header"`
	X      int32           `rosmsg:"x:int32"`
	Y      int32           `rosmsg:"y:int32"`
	Width  uint32          `rosmsg:"width:uint32"`
	Height uint32          `rosmsg:"height:uint32"`
	Data   []int8          `rosmsg:"data:int8[]"`
}

func (m *OccupancyGridUpdate) GetType() ros.MessageType {
	return MsgOccupancyGridUpdate
}

func (m *OccupancyGridUpdate) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, m.X)
	binary.Write(buf, binary.LittleEndian, m.Y)
	binary.Write(buf, binary.LittleEndian, m.Width)
	binary.Write(buf, binary.LittleEndian, m.Height)
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Data)))
	for _, e := range m.Data {
		binary.Write(buf, binary.LittleEndian, e)
	}
	return err
}

func (m *OccupancyGridUpdate) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.X); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Y); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Width); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Height); err != nil {
		return err
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Data = make([]int8, int(size))
		for i := 0; i < int(size); i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.Data[i]); err != nil {
				return err
			}
		}
	}
	return err
}

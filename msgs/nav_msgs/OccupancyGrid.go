// Automatically generated from the message definition "nav_msgs/OccupancyGrid.msg"
package nav_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgOccupancyGrid struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgOccupancyGrid) Text() string {
	return t.text
}

func (t *_MsgOccupancyGrid) Name() string {
	return t.name
}

func (t *_MsgOccupancyGrid) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgOccupancyGrid) NewMessage() ros.Message {
	m := new(OccupancyGrid)
	m.Header = std_msgs.Header{}
	m.Info = MapMetaData{}
	m.Data = []int8{}
	return m
}

var (
	MsgOccupancyGrid = &_MsgOccupancyGrid{
		`# This represents a 2-D grid map, in which each cell represents the probability of
# occupancy.

Header header 

#MetaData for the map
MapMetaData info

# The map data, in row-major order, starting with (0,0).  Occupancy
# probabilities are in the range [0,100].  Unknown is -1.
int8[] data
`,
		"nav_msgs/OccupancyGrid",
		"3381f2d731d4076ec5c71b0759edbe4e",
	}
)

type OccupancyGrid struct {
	Header std_msgs.Header `rosmsg:"header:Header"`
	Info   MapMetaData     `rosmsg:"info:MapMetaData"`
	Data   []int8          `rosmsg:"data:int8[]"`
}

func (m *OccupancyGrid) GetType() ros.MessageType {
	return MsgOccupancyGrid
}

func (m *OccupancyGrid) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	if err = m.Info.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Data)))
	for _, e := range m.Data {
		binary.Write(buf, binary.LittleEndian, e)
	}
	return err
}

func (m *OccupancyGrid) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = m.Info.Deserialize(buf); err != nil {
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

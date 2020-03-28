// Automatically generated from the message definition "nav_msgs/MapMetaData.msg"
package nav_msgs

import (
	"bytes"
	"ee631_midterm/msgs/geometry_msgs"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgMapMetaData struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgMapMetaData) Text() string {
	return t.text
}

func (t *_MsgMapMetaData) Name() string {
	return t.name
}

func (t *_MsgMapMetaData) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgMapMetaData) NewMessage() ros.Message {
	m := new(MapMetaData)
	m.MapLoadTime = ros.Time{}
	m.Resolution = 0.0
	m.Width = 0
	m.Height = 0
	m.Origin = geometry_msgs.Pose{}
	return m
}

var (
	MsgMapMetaData = &_MsgMapMetaData{
		`# This hold basic information about the characterists of the OccupancyGrid

# The time at which the map was loaded
time map_load_time
# The map resolution [m/cell]
float32 resolution
# Map width [cells]
uint32 width
# Map height [cells]
uint32 height
# The origin of the map [m, m, rad].  This is the real-world pose of the
# cell (0,0) in the map.
geometry_msgs/Pose origin`,
		"nav_msgs/MapMetaData",
		"10cfc8a2818024d3248802c00c95f11b",
	}
)

type MapMetaData struct {
	MapLoadTime ros.Time           `rosmsg:"map_load_time:time"`
	Resolution  float32            `rosmsg:"resolution:float32"`
	Width       uint32             `rosmsg:"width:uint32"`
	Height      uint32             `rosmsg:"height:uint32"`
	Origin      geometry_msgs.Pose `rosmsg:"origin:Pose"`
}

func (m *MapMetaData) GetType() ros.MessageType {
	return MsgMapMetaData
}

func (m *MapMetaData) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, m.MapLoadTime.Sec)
	binary.Write(buf, binary.LittleEndian, m.MapLoadTime.NSec)
	binary.Write(buf, binary.LittleEndian, m.Resolution)
	binary.Write(buf, binary.LittleEndian, m.Width)
	binary.Write(buf, binary.LittleEndian, m.Height)
	if err = m.Origin.Serialize(buf); err != nil {
		return err
	}
	return err
}

func (m *MapMetaData) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	{
		if err = binary.Read(buf, binary.LittleEndian, &m.MapLoadTime.Sec); err != nil {
			return err
		}
		if err = binary.Read(buf, binary.LittleEndian, &m.MapLoadTime.NSec); err != nil {
			return err
		}
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Resolution); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Width); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Height); err != nil {
		return err
	}
	if err = m.Origin.Deserialize(buf); err != nil {
		return err
	}
	return err
}

// Automatically generated from the message definition "map_msgs/ProjectedMap.msg"
package map_msgs

import (
	"bytes"
	"ee631_midterm/msgs/nav_msgs"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgProjectedMap struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgProjectedMap) Text() string {
	return t.text
}

func (t *_MsgProjectedMap) Name() string {
	return t.name
}

func (t *_MsgProjectedMap) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgProjectedMap) NewMessage() ros.Message {
	m := new(ProjectedMap)
	m.Map = nav_msgs.OccupancyGrid{}
	m.MinZ = 0.0
	m.MaxZ = 0.0
	return m
}

var (
	MsgProjectedMap = &_MsgProjectedMap{
		`nav_msgs/OccupancyGrid map
float64 min_z
float64 max_z`,
		"map_msgs/ProjectedMap",
		"7bbe8f96e45089681dc1ea7d023cbfca",
	}
)

type ProjectedMap struct {
	Map  nav_msgs.OccupancyGrid `rosmsg:"map:OccupancyGrid"`
	MinZ float64                `rosmsg:"min_z:float64"`
	MaxZ float64                `rosmsg:"max_z:float64"`
}

func (m *ProjectedMap) GetType() ros.MessageType {
	return MsgProjectedMap
}

func (m *ProjectedMap) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Map.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, m.MinZ)
	binary.Write(buf, binary.LittleEndian, m.MaxZ)
	return err
}

func (m *ProjectedMap) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Map.Deserialize(buf); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.MinZ); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.MaxZ); err != nil {
		return err
	}
	return err
}

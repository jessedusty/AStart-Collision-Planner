// Automatically generated from the message definition "map_msgs/ProjectedMapInfo.msg"
package map_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgProjectedMapInfo struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgProjectedMapInfo) Text() string {
	return t.text
}

func (t *_MsgProjectedMapInfo) Name() string {
	return t.name
}

func (t *_MsgProjectedMapInfo) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgProjectedMapInfo) NewMessage() ros.Message {
	m := new(ProjectedMapInfo)
	m.FrameId = ""
	m.X = 0.0
	m.Y = 0.0
	m.Width = 0.0
	m.Height = 0.0
	m.MinZ = 0.0
	m.MaxZ = 0.0
	return m
}

var (
	MsgProjectedMapInfo = &_MsgProjectedMapInfo{
		`string frame_id
float64 x
float64 y
float64 width
float64 height
float64 min_z
float64 max_z`,
		"map_msgs/ProjectedMapInfo",
		"2dc10595ae94de23f22f8a6d2a0eef7a",
	}
)

type ProjectedMapInfo struct {
	FrameId string  `rosmsg:"frame_id:string"`
	X       float64 `rosmsg:"x:float64"`
	Y       float64 `rosmsg:"y:float64"`
	Width   float64 `rosmsg:"width:float64"`
	Height  float64 `rosmsg:"height:float64"`
	MinZ    float64 `rosmsg:"min_z:float64"`
	MaxZ    float64 `rosmsg:"max_z:float64"`
}

func (m *ProjectedMapInfo) GetType() ros.MessageType {
	return MsgProjectedMapInfo
}

func (m *ProjectedMapInfo) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, uint32(len([]byte(m.FrameId))))
	buf.Write([]byte(m.FrameId))
	binary.Write(buf, binary.LittleEndian, m.X)
	binary.Write(buf, binary.LittleEndian, m.Y)
	binary.Write(buf, binary.LittleEndian, m.Width)
	binary.Write(buf, binary.LittleEndian, m.Height)
	binary.Write(buf, binary.LittleEndian, m.MinZ)
	binary.Write(buf, binary.LittleEndian, m.MaxZ)
	return err
}

func (m *ProjectedMapInfo) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		data := make([]byte, int(size))
		if err = binary.Read(buf, binary.LittleEndian, data); err != nil {
			return err
		}
		m.FrameId = string(data)
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
	if err = binary.Read(buf, binary.LittleEndian, &m.MinZ); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.MaxZ); err != nil {
		return err
	}
	return err
}

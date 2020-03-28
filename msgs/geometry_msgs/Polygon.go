// Automatically generated from the message definition "geometry_msgs/Polygon.msg"
package geometry_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgPolygon struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPolygon) Text() string {
	return t.text
}

func (t *_MsgPolygon) Name() string {
	return t.name
}

func (t *_MsgPolygon) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPolygon) NewMessage() ros.Message {
	m := new(Polygon)
	m.Points = []Point32{}
	return m
}

var (
	MsgPolygon = &_MsgPolygon{
		`#A specification of a polygon where the first and last points are assumed to be connected
Point32[] points
`,
		"geometry_msgs/Polygon",
		"cd60a26494a087f577976f0329fa120e",
	}
)

type Polygon struct {
	Points []Point32 `rosmsg:"points:Point32[]"`
}

func (m *Polygon) GetType() ros.MessageType {
	return MsgPolygon
}

func (m *Polygon) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Points)))
	for _, e := range m.Points {
		if err = e.Serialize(buf); err != nil {
			return err
		}
	}
	return err
}

func (m *Polygon) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Points = make([]Point32, int(size))
		for i := 0; i < int(size); i++ {
			if err = m.Points[i].Deserialize(buf); err != nil {
				return err
			}
		}
	}
	return err
}

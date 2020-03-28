// Automatically generated from the message definition "nav_msgs/GridCells.msg"
package nav_msgs

import (
	"bytes"
	"ee631_midterm/msgs/geometry_msgs"
	"ee631_midterm/msgs/std_msgs"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgGridCells struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgGridCells) Text() string {
	return t.text
}

func (t *_MsgGridCells) Name() string {
	return t.name
}

func (t *_MsgGridCells) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgGridCells) NewMessage() ros.Message {
	m := new(GridCells)
	m.Header = std_msgs.Header{}
	m.CellWidth = 0.0
	m.CellHeight = 0.0
	m.Cells = []geometry_msgs.Point{}
	return m
}

var (
	MsgGridCells = &_MsgGridCells{
		`#an array of cells in a 2D grid
Header header
float32 cell_width
float32 cell_height
geometry_msgs/Point[] cells
`,
		"nav_msgs/GridCells",
		"b9e4f5df6d28e272ebde00a3994830f5",
	}
)

type GridCells struct {
	Header     std_msgs.Header       `rosmsg:"header:Header"`
	CellWidth  float32               `rosmsg:"cell_width:float32"`
	CellHeight float32               `rosmsg:"cell_height:float32"`
	Cells      []geometry_msgs.Point `rosmsg:"cells:Point[]"`
}

func (m *GridCells) GetType() ros.MessageType {
	return MsgGridCells
}

func (m *GridCells) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, m.CellWidth)
	binary.Write(buf, binary.LittleEndian, m.CellHeight)
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Cells)))
	for _, e := range m.Cells {
		if err = e.Serialize(buf); err != nil {
			return err
		}
	}
	return err
}

func (m *GridCells) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.CellWidth); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.CellHeight); err != nil {
		return err
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Cells = make([]geometry_msgs.Point, int(size))
		for i := 0; i < int(size); i++ {
			if err = m.Cells[i].Deserialize(buf); err != nil {
				return err
			}
		}
	}
	return err
}

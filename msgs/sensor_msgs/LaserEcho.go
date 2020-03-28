// Automatically generated from the message definition "sensor_msgs/LaserEcho.msg"
package sensor_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgLaserEcho struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgLaserEcho) Text() string {
	return t.text
}

func (t *_MsgLaserEcho) Name() string {
	return t.name
}

func (t *_MsgLaserEcho) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgLaserEcho) NewMessage() ros.Message {
	m := new(LaserEcho)
	m.Echoes = []float32{}
	return m
}

var (
	MsgLaserEcho = &_MsgLaserEcho{
		`# This message is a submessage of MultiEchoLaserScan and is not intended
# to be used separately.

float32[] echoes  # Multiple values of ranges or intensities.
                  # Each array represents data from the same angle increment.`,
		"sensor_msgs/LaserEcho",
		"8bc5ae449b200fba4d552b4225586696",
	}
)

type LaserEcho struct {
	Echoes []float32 `rosmsg:"echoes:float32[]"`
}

func (m *LaserEcho) GetType() ros.MessageType {
	return MsgLaserEcho
}

func (m *LaserEcho) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Echoes)))
	for _, e := range m.Echoes {
		binary.Write(buf, binary.LittleEndian, e)
	}
	return err
}

func (m *LaserEcho) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Echoes = make([]float32, int(size))
		for i := 0; i < int(size); i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.Echoes[i]); err != nil {
				return err
			}
		}
	}
	return err
}

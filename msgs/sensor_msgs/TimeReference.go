// Automatically generated from the message definition "sensor_msgs/TimeReference.msg"
package sensor_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgTimeReference struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgTimeReference) Text() string {
	return t.text
}

func (t *_MsgTimeReference) Name() string {
	return t.name
}

func (t *_MsgTimeReference) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgTimeReference) NewMessage() ros.Message {
	m := new(TimeReference)
	m.Header = std_msgs.Header{}
	m.TimeRef = ros.Time{}
	m.Source = ""
	return m
}

var (
	MsgTimeReference = &_MsgTimeReference{
		`# Measurement from an external time source not actively synchronized with the system clock.

Header header    # stamp is system time for which measurement was valid
                 # frame_id is not used 

time   time_ref  # corresponding time from this external source
string source    # (optional) name of time source
`,
		"sensor_msgs/TimeReference",
		"fded64a0265108ba86c3d38fb11c0c16",
	}
)

type TimeReference struct {
	Header  std_msgs.Header `rosmsg:"header:Header"`
	TimeRef ros.Time        `rosmsg:"time_ref:time"`
	Source  string          `rosmsg:"source:string"`
}

func (m *TimeReference) GetType() ros.MessageType {
	return MsgTimeReference
}

func (m *TimeReference) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, m.TimeRef.Sec)
	binary.Write(buf, binary.LittleEndian, m.TimeRef.NSec)
	binary.Write(buf, binary.LittleEndian, uint32(len([]byte(m.Source))))
	buf.Write([]byte(m.Source))
	return err
}

func (m *TimeReference) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	{
		if err = binary.Read(buf, binary.LittleEndian, &m.TimeRef.Sec); err != nil {
			return err
		}
		if err = binary.Read(buf, binary.LittleEndian, &m.TimeRef.NSec); err != nil {
			return err
		}
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
		m.Source = string(data)
	}
	return err
}

// Automatically generated from the message definition "std_msgs/Float64MultiArray.msg"
package std_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgFloat64MultiArray struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgFloat64MultiArray) Text() string {
	return t.text
}

func (t *_MsgFloat64MultiArray) Name() string {
	return t.name
}

func (t *_MsgFloat64MultiArray) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgFloat64MultiArray) NewMessage() ros.Message {
	m := new(Float64MultiArray)
	m.Layout = MultiArrayLayout{}
	m.Data = []float64{}
	return m
}

var (
	MsgFloat64MultiArray = &_MsgFloat64MultiArray{
		`# Please look at the MultiArrayLayout message definition for
# documentation on all multiarrays.

MultiArrayLayout  layout        # specification of data layout
float64[]         data          # array of data

`,
		"std_msgs/Float64MultiArray",
		"4b7d974086d4060e7db4613a7e6c3ba4",
	}
)

type Float64MultiArray struct {
	Layout MultiArrayLayout `rosmsg:"layout:MultiArrayLayout"`
	Data   []float64        `rosmsg:"data:float64[]"`
}

func (m *Float64MultiArray) GetType() ros.MessageType {
	return MsgFloat64MultiArray
}

func (m *Float64MultiArray) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Layout.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Data)))
	for _, e := range m.Data {
		binary.Write(buf, binary.LittleEndian, e)
	}
	return err
}

func (m *Float64MultiArray) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Layout.Deserialize(buf); err != nil {
		return err
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Data = make([]float64, int(size))
		for i := 0; i < int(size); i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.Data[i]); err != nil {
				return err
			}
		}
	}
	return err
}

// Automatically generated from the message definition "std_msgs/Float32MultiArray.msg"
package std_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgFloat32MultiArray struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgFloat32MultiArray) Text() string {
	return t.text
}

func (t *_MsgFloat32MultiArray) Name() string {
	return t.name
}

func (t *_MsgFloat32MultiArray) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgFloat32MultiArray) NewMessage() ros.Message {
	m := new(Float32MultiArray)
	m.Layout = MultiArrayLayout{}
	m.Data = []float32{}
	return m
}

var (
	MsgFloat32MultiArray = &_MsgFloat32MultiArray{
		`# Please look at the MultiArrayLayout message definition for
# documentation on all multiarrays.

MultiArrayLayout  layout        # specification of data layout
float32[]         data          # array of data

`,
		"std_msgs/Float32MultiArray",
		"6a40e0ffa6a17a503ac3f8616991b1f6",
	}
)

type Float32MultiArray struct {
	Layout MultiArrayLayout `rosmsg:"layout:MultiArrayLayout"`
	Data   []float32        `rosmsg:"data:float32[]"`
}

func (m *Float32MultiArray) GetType() ros.MessageType {
	return MsgFloat32MultiArray
}

func (m *Float32MultiArray) Serialize(buf *bytes.Buffer) error {
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

func (m *Float32MultiArray) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Layout.Deserialize(buf); err != nil {
		return err
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Data = make([]float32, int(size))
		for i := 0; i < int(size); i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.Data[i]); err != nil {
				return err
			}
		}
	}
	return err
}

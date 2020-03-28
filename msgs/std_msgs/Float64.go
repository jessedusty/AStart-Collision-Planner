// Automatically generated from the message definition "std_msgs/Float64.msg"
package std_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgFloat64 struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgFloat64) Text() string {
	return t.text
}

func (t *_MsgFloat64) Name() string {
	return t.name
}

func (t *_MsgFloat64) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgFloat64) NewMessage() ros.Message {
	m := new(Float64)
	m.Data = 0.0
	return m
}

var (
	MsgFloat64 = &_MsgFloat64{
		`float64 data`,
		"std_msgs/Float64",
		"fdb28210bfa9d7c91146260178d9a584",
	}
)

type Float64 struct {
	Data float64 `rosmsg:"data:float64"`
}

func (m *Float64) GetType() ros.MessageType {
	return MsgFloat64
}

func (m *Float64) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, m.Data)
	return err
}

func (m *Float64) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = binary.Read(buf, binary.LittleEndian, &m.Data); err != nil {
		return err
	}
	return err
}

// Automatically generated from the message definition "tf2_msgs/TF2Error.msg"
package tf2_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

const (
	TF2Error_NO_ERROR               uint8 = 0
	TF2Error_LOOKUP_ERROR           uint8 = 1
	TF2Error_CONNECTIVITY_ERROR     uint8 = 2
	TF2Error_EXTRAPOLATION_ERROR    uint8 = 3
	TF2Error_INVALID_ARGUMENT_ERROR uint8 = 4
	TF2Error_TIMEOUT_ERROR          uint8 = 5
	TF2Error_TRANSFORM_ERROR        uint8 = 6
)

type _MsgTF2Error struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgTF2Error) Text() string {
	return t.text
}

func (t *_MsgTF2Error) Name() string {
	return t.name
}

func (t *_MsgTF2Error) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgTF2Error) NewMessage() ros.Message {
	m := new(TF2Error)
	m.Error = 0
	m.ErrorString = ""
	return m
}

var (
	MsgTF2Error = &_MsgTF2Error{
		`uint8 NO_ERROR = 0
uint8 LOOKUP_ERROR = 1
uint8 CONNECTIVITY_ERROR = 2
uint8 EXTRAPOLATION_ERROR = 3
uint8 INVALID_ARGUMENT_ERROR = 4
uint8 TIMEOUT_ERROR = 5
uint8 TRANSFORM_ERROR = 6

uint8 error
string error_string
`,
		"tf2_msgs/TF2Error",
		"bc6848fd6fd750c92e38575618a4917d",
	}
)

type TF2Error struct {
	Error       uint8  `rosmsg:"error:uint8"`
	ErrorString string `rosmsg:"error_string:string"`
}

func (m *TF2Error) GetType() ros.MessageType {
	return MsgTF2Error
}

func (m *TF2Error) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, m.Error)
	binary.Write(buf, binary.LittleEndian, uint32(len([]byte(m.ErrorString))))
	buf.Write([]byte(m.ErrorString))
	return err
}

func (m *TF2Error) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = binary.Read(buf, binary.LittleEndian, &m.Error); err != nil {
		return err
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
		m.ErrorString = string(data)
	}
	return err
}

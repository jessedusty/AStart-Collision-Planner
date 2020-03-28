// Automatically generated from the message definition "nav_msgs/GetMapResult.msg"
package nav_msgs

import (
	"bytes"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgGetMapResult struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgGetMapResult) Text() string {
	return t.text
}

func (t *_MsgGetMapResult) Name() string {
	return t.name
}

func (t *_MsgGetMapResult) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgGetMapResult) NewMessage() ros.Message {
	m := new(GetMapResult)
	m.Map = OccupancyGrid{}
	return m
}

var (
	MsgGetMapResult = &_MsgGetMapResult{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======
nav_msgs/OccupancyGrid map
`,
		"nav_msgs/GetMapResult",
		"6cdd0a18e0aff5b0a3ca2326a89b54ff",
	}
)

type GetMapResult struct {
	Map OccupancyGrid `rosmsg:"map:OccupancyGrid"`
}

func (m *GetMapResult) GetType() ros.MessageType {
	return MsgGetMapResult
}

func (m *GetMapResult) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Map.Serialize(buf); err != nil {
		return err
	}
	return err
}

func (m *GetMapResult) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Map.Deserialize(buf); err != nil {
		return err
	}
	return err
}

// Automatically generated from the message definition "sensor_msgs/ChannelFloat32.msg"
package sensor_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgChannelFloat32 struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgChannelFloat32) Text() string {
	return t.text
}

func (t *_MsgChannelFloat32) Name() string {
	return t.name
}

func (t *_MsgChannelFloat32) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgChannelFloat32) NewMessage() ros.Message {
	m := new(ChannelFloat32)
	m.Name = ""
	m.Values = []float32{}
	return m
}

var (
	MsgChannelFloat32 = &_MsgChannelFloat32{
		`# This message is used by the PointCloud message to hold optional data
# associated with each point in the cloud. The length of the values
# array should be the same as the length of the points array in the
# PointCloud, and each value should be associated with the corresponding
# point.

# Channel names in existing practice include:
#   "u", "v" - row and column (respectively) in the left stereo image.
#              This is opposite to usual conventions but remains for
#              historical reasons. The newer PointCloud2 message has no
#              such problem.
#   "rgb" - For point clouds produced by color stereo cameras. uint8
#           (R,G,B) values packed into the least significant 24 bits,
#           in order.
#   "intensity" - laser or pixel intensity.
#   "distance"

# The channel name should give semantics of the channel (e.g.
# "intensity" instead of "value").
string name

# The values array should be 1-1 with the elements of the associated
# PointCloud.
float32[] values
`,
		"sensor_msgs/ChannelFloat32",
		"3d40139cdd33dfedcb71ffeeeb42ae7f",
	}
)

type ChannelFloat32 struct {
	Name   string    `rosmsg:"name:string"`
	Values []float32 `rosmsg:"values:float32[]"`
}

func (m *ChannelFloat32) GetType() ros.MessageType {
	return MsgChannelFloat32
}

func (m *ChannelFloat32) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, uint32(len([]byte(m.Name))))
	buf.Write([]byte(m.Name))
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Values)))
	for _, e := range m.Values {
		binary.Write(buf, binary.LittleEndian, e)
	}
	return err
}

func (m *ChannelFloat32) Deserialize(buf *bytes.Reader) error {
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
		m.Name = string(data)
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Values = make([]float32, int(size))
		for i := 0; i < int(size); i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.Values[i]); err != nil {
				return err
			}
		}
	}
	return err
}

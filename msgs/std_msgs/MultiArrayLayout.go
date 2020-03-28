// Automatically generated from the message definition "std_msgs/MultiArrayLayout.msg"
package std_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgMultiArrayLayout struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgMultiArrayLayout) Text() string {
	return t.text
}

func (t *_MsgMultiArrayLayout) Name() string {
	return t.name
}

func (t *_MsgMultiArrayLayout) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgMultiArrayLayout) NewMessage() ros.Message {
	m := new(MultiArrayLayout)
	m.Dim = []MultiArrayDimension{}
	m.DataOffset = 0
	return m
}

var (
	MsgMultiArrayLayout = &_MsgMultiArrayLayout{
		`# The multiarray declares a generic multi-dimensional array of a
# particular data type.  Dimensions are ordered from outer most
# to inner most.

MultiArrayDimension[] dim # Array of dimension properties
uint32 data_offset        # padding elements at front of data

# Accessors should ALWAYS be written in terms of dimension stride
# and specified outer-most dimension first.
# 
# multiarray(i,j,k) = data[data_offset + dim_stride[1]*i + dim_stride[2]*j + k]
#
# A standard, 3-channel 640x480 image with interleaved color channels
# would be specified as:
#
# dim[0].label  = "height"
# dim[0].size   = 480
# dim[0].stride = 3*640*480 = 921600  (note dim[0] stride is just size of image)
# dim[1].label  = "width"
# dim[1].size   = 640
# dim[1].stride = 3*640 = 1920
# dim[2].label  = "channel"
# dim[2].size   = 3
# dim[2].stride = 3
#
# multiarray(i,j,k) refers to the ith row, jth column, and kth channel.
`,
		"std_msgs/MultiArrayLayout",
		"0fed2a11c13e11c5571b4e2a995a91a3",
	}
)

type MultiArrayLayout struct {
	Dim        []MultiArrayDimension `rosmsg:"dim:MultiArrayDimension[]"`
	DataOffset uint32                `rosmsg:"data_offset:uint32"`
}

func (m *MultiArrayLayout) GetType() ros.MessageType {
	return MsgMultiArrayLayout
}

func (m *MultiArrayLayout) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Dim)))
	for _, e := range m.Dim {
		if err = e.Serialize(buf); err != nil {
			return err
		}
	}
	binary.Write(buf, binary.LittleEndian, m.DataOffset)
	return err
}

func (m *MultiArrayLayout) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Dim = make([]MultiArrayDimension, int(size))
		for i := 0; i < int(size); i++ {
			if err = m.Dim[i].Deserialize(buf); err != nil {
				return err
			}
		}
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.DataOffset); err != nil {
		return err
	}
	return err
}

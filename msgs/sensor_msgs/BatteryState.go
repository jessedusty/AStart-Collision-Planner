// Automatically generated from the message definition "sensor_msgs/BatteryState.msg"
package sensor_msgs

import (
	"bytes"
	"ee631_midterm/msgs/std_msgs"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

const (
	BatteryState_POWER_SUPPLY_STATUS_UNKNOWN               uint8 = 0
	BatteryState_POWER_SUPPLY_STATUS_CHARGING              uint8 = 1
	BatteryState_POWER_SUPPLY_STATUS_DISCHARGING           uint8 = 2
	BatteryState_POWER_SUPPLY_STATUS_NOT_CHARGING          uint8 = 3
	BatteryState_POWER_SUPPLY_STATUS_FULL                  uint8 = 4
	BatteryState_POWER_SUPPLY_HEALTH_UNKNOWN               uint8 = 0
	BatteryState_POWER_SUPPLY_HEALTH_GOOD                  uint8 = 1
	BatteryState_POWER_SUPPLY_HEALTH_OVERHEAT              uint8 = 2
	BatteryState_POWER_SUPPLY_HEALTH_DEAD                  uint8 = 3
	BatteryState_POWER_SUPPLY_HEALTH_OVERVOLTAGE           uint8 = 4
	BatteryState_POWER_SUPPLY_HEALTH_UNSPEC_FAILURE        uint8 = 5
	BatteryState_POWER_SUPPLY_HEALTH_COLD                  uint8 = 6
	BatteryState_POWER_SUPPLY_HEALTH_WATCHDOG_TIMER_EXPIRE uint8 = 7
	BatteryState_POWER_SUPPLY_HEALTH_SAFETY_TIMER_EXPIRE   uint8 = 8
	BatteryState_POWER_SUPPLY_TECHNOLOGY_UNKNOWN           uint8 = 0
	BatteryState_POWER_SUPPLY_TECHNOLOGY_NIMH              uint8 = 1
	BatteryState_POWER_SUPPLY_TECHNOLOGY_LION              uint8 = 2
	BatteryState_POWER_SUPPLY_TECHNOLOGY_LIPO              uint8 = 3
	BatteryState_POWER_SUPPLY_TECHNOLOGY_LIFE              uint8 = 4
	BatteryState_POWER_SUPPLY_TECHNOLOGY_NICD              uint8 = 5
	BatteryState_POWER_SUPPLY_TECHNOLOGY_LIMN              uint8 = 6
)

type _MsgBatteryState struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgBatteryState) Text() string {
	return t.text
}

func (t *_MsgBatteryState) Name() string {
	return t.name
}

func (t *_MsgBatteryState) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgBatteryState) NewMessage() ros.Message {
	m := new(BatteryState)
	m.Header = std_msgs.Header{}
	m.Voltage = 0.0
	m.Current = 0.0
	m.Charge = 0.0
	m.Capacity = 0.0
	m.DesignCapacity = 0.0
	m.Percentage = 0.0
	m.PowerSupplyStatus = 0
	m.PowerSupplyHealth = 0
	m.PowerSupplyTechnology = 0
	m.Present = false
	m.CellVoltage = []float32{}
	m.Location = ""
	m.SerialNumber = ""
	return m
}

var (
	MsgBatteryState = &_MsgBatteryState{
		`
# Constants are chosen to match the enums in the linux kernel
# defined in include/linux/power_supply.h as of version 3.7
# The one difference is for style reasons the constants are
# all uppercase not mixed case.

# Power supply status constants
uint8 POWER_SUPPLY_STATUS_UNKNOWN = 0
uint8 POWER_SUPPLY_STATUS_CHARGING = 1
uint8 POWER_SUPPLY_STATUS_DISCHARGING = 2
uint8 POWER_SUPPLY_STATUS_NOT_CHARGING = 3
uint8 POWER_SUPPLY_STATUS_FULL = 4

# Power supply health constants
uint8 POWER_SUPPLY_HEALTH_UNKNOWN = 0
uint8 POWER_SUPPLY_HEALTH_GOOD = 1
uint8 POWER_SUPPLY_HEALTH_OVERHEAT = 2
uint8 POWER_SUPPLY_HEALTH_DEAD = 3
uint8 POWER_SUPPLY_HEALTH_OVERVOLTAGE = 4
uint8 POWER_SUPPLY_HEALTH_UNSPEC_FAILURE = 5
uint8 POWER_SUPPLY_HEALTH_COLD = 6
uint8 POWER_SUPPLY_HEALTH_WATCHDOG_TIMER_EXPIRE = 7
uint8 POWER_SUPPLY_HEALTH_SAFETY_TIMER_EXPIRE = 8

# Power supply technology (chemistry) constants
uint8 POWER_SUPPLY_TECHNOLOGY_UNKNOWN = 0
uint8 POWER_SUPPLY_TECHNOLOGY_NIMH = 1
uint8 POWER_SUPPLY_TECHNOLOGY_LION = 2
uint8 POWER_SUPPLY_TECHNOLOGY_LIPO = 3
uint8 POWER_SUPPLY_TECHNOLOGY_LIFE = 4
uint8 POWER_SUPPLY_TECHNOLOGY_NICD = 5
uint8 POWER_SUPPLY_TECHNOLOGY_LIMN = 6

Header  header
float32 voltage          # Voltage in Volts (Mandatory)
float32 current          # Negative when discharging (A)  (If unmeasured NaN)
float32 charge           # Current charge in Ah  (If unmeasured NaN)
float32 capacity         # Capacity in Ah (last full capacity)  (If unmeasured NaN)
float32 design_capacity  # Capacity in Ah (design capacity)  (If unmeasured NaN)
float32 percentage       # Charge percentage on 0 to 1 range  (If unmeasured NaN)
uint8   power_supply_status     # The charging status as reported. Values defined above
uint8   power_supply_health     # The battery health metric. Values defined above
uint8   power_supply_technology # The battery chemistry. Values defined above
bool    present          # True if the battery is present

float32[] cell_voltage   # An array of individual cell voltages for each cell in the pack
                         # If individual voltages unknown but number of cells known set each to NaN
string location          # The location into which the battery is inserted. (slot number or plug)
string serial_number     # The best approximation of the battery serial number
`,
		"sensor_msgs/BatteryState",
		"476f837fa6771f6e16e3bf4ef96f8770",
	}
)

type BatteryState struct {
	Header                std_msgs.Header `rosmsg:"header:Header"`
	Voltage               float32         `rosmsg:"voltage:float32"`
	Current               float32         `rosmsg:"current:float32"`
	Charge                float32         `rosmsg:"charge:float32"`
	Capacity              float32         `rosmsg:"capacity:float32"`
	DesignCapacity        float32         `rosmsg:"design_capacity:float32"`
	Percentage            float32         `rosmsg:"percentage:float32"`
	PowerSupplyStatus     uint8           `rosmsg:"power_supply_status:uint8"`
	PowerSupplyHealth     uint8           `rosmsg:"power_supply_health:uint8"`
	PowerSupplyTechnology uint8           `rosmsg:"power_supply_technology:uint8"`
	Present               bool            `rosmsg:"present:bool"`
	CellVoltage           []float32       `rosmsg:"cell_voltage:float32[]"`
	Location              string          `rosmsg:"location:string"`
	SerialNumber          string          `rosmsg:"serial_number:string"`
}

func (m *BatteryState) GetType() ros.MessageType {
	return MsgBatteryState
}

func (m *BatteryState) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, m.Voltage)
	binary.Write(buf, binary.LittleEndian, m.Current)
	binary.Write(buf, binary.LittleEndian, m.Charge)
	binary.Write(buf, binary.LittleEndian, m.Capacity)
	binary.Write(buf, binary.LittleEndian, m.DesignCapacity)
	binary.Write(buf, binary.LittleEndian, m.Percentage)
	binary.Write(buf, binary.LittleEndian, m.PowerSupplyStatus)
	binary.Write(buf, binary.LittleEndian, m.PowerSupplyHealth)
	binary.Write(buf, binary.LittleEndian, m.PowerSupplyTechnology)
	binary.Write(buf, binary.LittleEndian, m.Present)
	binary.Write(buf, binary.LittleEndian, uint32(len(m.CellVoltage)))
	for _, e := range m.CellVoltage {
		binary.Write(buf, binary.LittleEndian, e)
	}
	binary.Write(buf, binary.LittleEndian, uint32(len([]byte(m.Location))))
	buf.Write([]byte(m.Location))
	binary.Write(buf, binary.LittleEndian, uint32(len([]byte(m.SerialNumber))))
	buf.Write([]byte(m.SerialNumber))
	return err
}

func (m *BatteryState) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Voltage); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Current); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Charge); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Capacity); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.DesignCapacity); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Percentage); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.PowerSupplyStatus); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.PowerSupplyHealth); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.PowerSupplyTechnology); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Present); err != nil {
		return err
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.CellVoltage = make([]float32, int(size))
		for i := 0; i < int(size); i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.CellVoltage[i]); err != nil {
				return err
			}
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
		m.Location = string(data)
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
		m.SerialNumber = string(data)
	}
	return err
}

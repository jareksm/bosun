package snmpDev

type SnmpDevice interface {
	New(h, c string) error
	Check() error
	BuildIface() error
}

type PhysClass int

const (
	other PhysClass = iota + 1
	unknown
	chassis
	backplane
	container
	powerSupply
	fan
	sensor
	module
	port
	stack
	cpu
)

type Pkts struct {
	in, out uint64
}

type Revision struct {
	Firmware string
	Hdw      string
	Software string
}

type Iface struct {
	adminStatus string
	alias       string
	brd         Pkts
	desc        string
	discards    Pkts
	errors      Pkts
	hiSpeed     string
	ifIndex     int
	ifType      string
	lastChange  uint64 // 100 Timeticks = 1sec
	mac         string
	mcst        Pkts
	mtu         string
	name        string
	oct         Pkts
	operStatus  string
	pauseFrames Pkts
	ucst        Pkts
}

type MemPool struct {
	Free     uint
	PoolType string
	Used     uint
}

type PhysHdw struct {
	Alias  string
	Asset  string
	Class  PhysClass
	Desc   string
	FRU    bool
	Mfg    string
	Model  string
	Name   string
	Rev    Revision
	Serial string
	Vendor string
}

type GenericDevice struct {
	Community string
	Cpu       int
	Desc      string
	Hardware  map[int]PhysHdw
	Hostname  string
	Mem       []MemPool
	Ports     map[int]Iface
}

type CiscoSwitch struct {
	GenericDevice
}

// Implements SnmpDevice interface, takes host and community as arguments.
func (s *CiscoSwitch) New(h, c string) {
}
func (s *CiscoSwitch) Check() {
}
func (s *CiscoSwitch) BuildIface() {
}

type AristaSwitch struct {
	GenericDevice
}

// Implements SnmpDevice interface, takes host and community as arguments.
func (s *AristaSwitch) New(h, c string) {
}
func (s *AristaSwitch) Check() {
}
func (s *AristaSwitch) BuildIface() {
}

type PaloAltoFW struct {
	GenericDevice
}

// Implements SnmpDevice interface, takes host and community as arguments.
func (s *PaloAltoFW) New(h, c string) {
}
func (s *PaloAltoFW) Check() {
}
func (s *PaloAltoFW) BuildIface() {
}

type MellanoxSwitch struct {
	GenericDevice
}

// Implements SnmpDevice interface, takes host and community as arguments.
func (s *MellanoxSwitch) New(h, c string) {
}
func (s *MellanoxSwitch) Check() {
}
func (s *MellanoxSwitch) BuildIface() {
}

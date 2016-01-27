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
	Hdw      string
	Firmware string
	Software string
}

type Iface struct {
	ifIndex     int
	name        string
	alias       string
	desc        string
	ifType      string
	mtu         string
	mac         string
	hiSpeed     string
	adminStatus string
	operStatus  string
	lastChange  uint64 // 100 Timeticks = 1sec
	brd         Pkts
	mcst        Pkts
	ucst        Pkts
	oct         Pkts
	discards    Pkts
	errors      Pkts
	pauseFrames Pkts
}

type MemPool struct {
	PoolType string
	Used     uint
	Free     uint
}

type PhysHdw struct {
	Desc   string
	Vendor string
	Class  PhysClass
	Name   string
	Rev    Revision
	Serial string
	Mfg    string
	Model  string
	Alias  string
	Asset  string
	FRU    bool
}

type GenericDevice struct {
	Ports    map[int]Iface
	Hardware map[int]PhysHdw
	Mem      []MemPool
	Cpu      int
	Desc     string
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

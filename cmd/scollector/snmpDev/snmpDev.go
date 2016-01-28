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
	in  uint64
	out uint64
}

type Status struct {
	admin string `oid:"ifAdminStatus"`
	oper  string `oid:"ifOperStatus"`
}

// 100 Timeticks = 1sec
type Iface struct {
	Status
	brd          Pkts `oid:"ifInBroadcastPkts,ifOutBroadcastPkts"`
	bytes        Pkts `oid:"ifInOctets,ifOutOctets"`
	connector    bool `oid:"ifConnectorPresent"`
	discards     Pkts `oid:"ifInDiscards,ifOutDiscards"`
	errors       Pkts `oid:"ifInErrors,ifOutErrors"`
	ifAlias      string
	ifDescr      string
	ifHighSpeed  string
	ifIndex      int
	ifLastChange uint64
	ifName       string
	ifType       string
	ips          []string
	mac          string `oid:"ifPhysAddress"`
	mcst         Pkts   `oid:"ifInMulticastPkts,ifOutMulticastPkts"`
	mtu          string
	name         string `oid:"ifName"`
	pauseFrames  Pkts
	promisc      bool `oid:"ifPromiscuousMode"`
	traps        bool `oid:"ifLinkUpDownTrapEnable"`
	ucst         Pkts `oid:"ifInUcastPkts"`
}

type MemPool struct {
	Free uint
	Type string
	Used uint
}

type Revision struct {
	Firmware string `oid:"entPhysicalFirmwareRev"`
	Hdw      string `oid:"entPhysicalHardwareRev"`
	Software string `oid:"entPhysicalSoftwareRev"`
}

type PhysHdw struct {
	Revision
	Alias  string    `oid:"entPhysicalAlias"`
	Asset  string    `oid:"entPhysicalAssetID"`
	Class  PhysClass `oid:"entPhysicalClass"`
	Desc   string    `oid:"entPhysicalDescr"`
	FRU    bool      `oid:"entPhysicalIsFRU"`
	Mfg    string    `oid:"entPhysicalMfgName"` //arista
	Model  string    `oid:"entPhysicalModelName"`
	Name   string    `oid:"entPhysicalName"` //cisco only
	Serial string    `oid:"entPhysicalSerialNum"`
	Vendor string    `oid:"entPhysicalVendorType"` //cisco only
}

type GenericDevice struct {
	Community string
	Cpu       int
	Desc      string `oid:"sysDescr"`
	Hardware  map[int]PhysHdw
	Hostname  string
	Mem       []MemPool
	Ports     map[int]Iface
	Uptime    int `oid:"sysUpTime"` // Timeticks = 1sec / 100
	fruIndex  []int
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

type AristaCpu struct {
	Desc   string `oid:"hrDeviceDescr"`
	Status int    `oid:"hrDeviceStatus"`
	Load   int    `oid:"hrProcessorLoad"`
}

type AristaMemPool struct {
	Type string `oid:"hrStorageType"`
	Desc string `oid:"hrStorageDescr"`
	Unit string `oid:"hrStorageAllocationUnits"`
	Size uint   `oid:"hrStorageSize"`
	Used uint   `oid:"hrStorageUsed"`
}

type AristaSwitch struct {
	GenericDevice
	MemSize int `oid:"hrMemorySize"`
	Mem     []AristaMemPool
	Cpu     []AristaCpu
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

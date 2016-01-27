package collectors

type SnmpDevice interface {
	New(h, c string) error
	Check() error
	BuildIface() error
}

type Pkts struct {
	in, out uint64
}

type SnmpIface struct {
	name        string
	alias       string
	desc        string
	ifType      string
	mtu         string
	mac         string
	hiSpeed     string
	adminStatus string
	operStatus  string
	Brd         Pkts
	Mcst        Pkts
	Ucst        Pkts
	Oct         Pkts
	Discards    Pkts
	Errors      Pkts
	PauseFrames Pkts
}

type GenericDevice struct {
	Iface []SnmpIface
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

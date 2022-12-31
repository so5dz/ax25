package ax25

type Packet struct {
	Header  Header
	Control byte
	PID     byte
	Data    []byte
}

type Header struct {
	Source      Callsign
	Destination Callsign
	Path        []Callsign
}

type Callsign struct {
	Base       string
	Substation int
	HBR        bool
}

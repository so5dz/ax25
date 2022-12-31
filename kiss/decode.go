package kiss

import (
	"strings"

	"github.com/so5dz/ax25"
)

func DecodePacket(rawPacket []byte) (ax25.Packet, error) {
	var packet ax25.Packet

	headerEnd := 0
	for i, b := range rawPacket {
		if isHeaderEndBitSet(b) {
			headerEnd = i
			break
		}
	}
	rawHeader := rawPacket[0 : headerEnd+1]
	header, err := DecodeHeader(rawHeader)
	if err != nil {
		return packet, err
	}
	packet.Header = header
	packet.Control = rawPacket[headerEnd+1]
	packet.PID = rawPacket[headerEnd+2]
	packet.Data = rawPacket[headerEnd+3:]

	return packet, nil
}

func DecodeHeader(rawHeader []byte) (ax25.Header, error) {
	var header ax25.Header

	callsignCount := len(rawHeader) / 7
	callsigns := make([]ax25.Callsign, 0, callsignCount)
	for i := 0; i < callsignCount; i++ {
		callsign, err := DecodeCallsign(rawHeader[7*i : 7*i+7])
		if err != nil {
			return header, err
		}
		callsigns = append(callsigns, callsign)
	}

	header.Destination = callsigns[0]
	header.Source = callsigns[1]
	header.Path = callsigns[2:]

	return header, nil
}

func DecodeCallsign(rawCallsign []byte) (ax25.Callsign, error) {
	var callsign ax25.Callsign

	rawBase := rawCallsign[0:6]
	for i, b := range rawBase {
		rawBase[i] = b >> 1
	}
	callsign.Base = strings.TrimSpace(string(rawBase))
	rawLastByte := rawCallsign[6]
	callsign.Substation = (int(rawLastByte) & 0b00011110) >> 1
	callsign.HBR = (int(rawLastByte) >> 7) > 0

	return callsign, nil
}

func isHeaderEndBitSet(b byte) bool {
	return b&1 == 1
}

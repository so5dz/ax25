package tnc2

import (
	"fmt"
	"strings"

	"github.com/so5dz/ax25"
)

func EncodePacket(packet ax25.Packet) string {
	return fmt.Sprintf("%s:%s", EncodeHeader(packet.Header), packet.Data)
}

func EncodeHeader(header ax25.Header) string {
	var sb strings.Builder
	sb.WriteString(EncodeCallsign(header.Source))
	sb.WriteByte('>')
	sb.WriteString(EncodeCallsign(header.Destination))
	for _, via := range header.Path {
		sb.WriteByte(',')
		sb.WriteString(EncodeCallsign(via))
	}
	return sb.String()
}

func EncodeCallsign(callsign ax25.Callsign) string {
	var sb strings.Builder
	sb.WriteString(callsign.Base)
	if callsign.Substation > 0 {
		sb.WriteString(fmt.Sprintf("-%d", callsign.Substation))
	}
	if callsign.HBR {
		sb.WriteByte('*')
	}
	return sb.String()
}

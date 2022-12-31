package tnc2

import (
	"errors"
	"strconv"
	"strings"

	"github.com/so5dz/ax25"
)

func DecodePacket(tnc2Packet string) (ax25.Packet, error) {
	var packet ax25.Packet
	var err error

	headerEnd := strings.IndexRune(tnc2Packet, ':')
	if headerEnd < 0 {
		return packet, errors.New("packet does not appear to contain header")
	}

	packet.Header, err = DecodeHeader(tnc2Packet[0:headerEnd])
	if err != nil {
		return packet, err
	}

	packet.PID = 0
	packet.Control = 0
	packet.Data = []byte(tnc2Packet[headerEnd+1:])

	return packet, nil
}

func DecodeHeader(tnc2Header string) (ax25.Header, error) {
	var header ax25.Header
	var err error

	sides := strings.SplitN(tnc2Header, ">", 2)
	if len(sides) != 2 {
		return header, errors.New("no > in tnc2 header")
	}

	header.Source, err = DecodeCallsign(sides[0])
	if err != nil {
		return header, err
	}

	others := strings.Split(sides[1], ",")
	if len(others) < 1 {
		return header, errors.New("empty destination and path")
	}

	header.Destination, err = DecodeCallsign(others[0])
	if err != nil {
		return header, err
	}

	if len(others) > 1 {
		header.Path = make([]ax25.Callsign, 0, len(others)-1)
		for _, pathCallsign := range others[1:] {
			callsign, err := DecodeCallsign(pathCallsign)
			if err != nil {
				return header, err
			}
			header.Path = append(header.Path, callsign)
		}
	}

	return header, nil
}

func DecodeCallsign(tnc2Callsign string) (ax25.Callsign, error) {
	var callsign ax25.Callsign

	containsStar := strings.ContainsRune(tnc2Callsign, '*')
	if containsStar {
		callsign.HBR = true
		tnc2Callsign = strings.Replace(tnc2Callsign, "*", "", -1)
	}

	parts := strings.SplitN(tnc2Callsign, "-", 2)
	callsign.Base = parts[0]

	if len(parts) == 2 {
		ssid, err := strconv.Atoi(parts[1])
		if err != nil {
			return callsign, err
		}
		callsign.Substation = ssid
	}

	return callsign, nil
}

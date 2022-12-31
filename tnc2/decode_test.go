package tnc2

import (
	"reflect"
	"testing"

	"github.com/so5dz/ax25"
)

func TestDecodePacket(t *testing.T) {
	type args struct {
		tnc2Packet string
	}
	tests := []struct {
		name    string
		args    args
		want    ax25.Packet
		wantErr bool
	}{
		{
			name: "basic",
			args: args{tnc2Packet: "SO5DZ-1>APRS:basic"},
			want: ax25.Packet{
				Header: ax25.Header{
					Source: ax25.Callsign{
						Base:       "SO5DZ",
						Substation: 1,
						HBR:        false,
					},
					Destination: ax25.Callsign{
						Base:       "APRS",
						Substation: 0,
						HBR:        false,
					},
				},
				Data: []byte("basic"),
			},
			wantErr: false,
		},
		{
			name: "path",
			args: args{tnc2Packet: "SOURCE-123>DESTINATION-15,VIA1-1,VIA2*:"},
			want: ax25.Packet{
				Header: ax25.Header{
					Source: ax25.Callsign{
						Base:       "SOURCE",
						Substation: 123,
						HBR:        false,
					},
					Destination: ax25.Callsign{
						Base:       "DESTINATION",
						Substation: 15,
						HBR:        false,
					},
					Path: []ax25.Callsign{
						{
							Base:       "VIA1",
							Substation: 1,
							HBR:        false,
						},
						{
							Base:       "VIA2",
							Substation: 0,
							HBR:        true,
						},
					},
				},
				Data: []byte{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodePacket(tt.args.tnc2Packet)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodePacket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodePacket() = %v, want %v", got, tt.want)
			}
		})
	}
}

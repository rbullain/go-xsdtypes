package xsdtypes_test

import (
	"encoding/xml"
	"fmt"
	"testing"
	"time"

	"github.com/rbullain/go-xsdtypes"
)

func TestDateTime_UnmarshalXML(t *testing.T) {
	t.Parallel()

	tzCest := time.FixedZone("CEST", 2*60*60)
	tzCst := time.FixedZone("CST", -5*60*60)

	tests := []struct {
		name    string
		value   string
		want    time.Time
		wantErr bool
	}{
		{
			name:  "ok simple",
			value: "2024-01-01T02:20:00",
			want:  time.Date(2024, 1, 1, 2, 20, 0, 0, time.UTC),
		},
		{
			name:  "ok utc",
			value: "2024-01-01T02:20:00Z",
			want:  time.Date(2024, 1, 1, 2, 20, 0, 0, time.UTC),
		},
		{
			name:  "ok simple with offset",
			value: "2024-01-01T02:20:00+02:00",
			want:  time.Date(2024, 1, 1, 2, 20, 0, 0, tzCest),
		},
		{
			name:  "ok simple with offset negative",
			value: "2024-01-01T02:20:00-05:00",
			want:  time.Date(2024, 1, 1, 2, 20, 0, 0, tzCst),
		},
		{
			name:  "ok simple with fraction of second",
			value: "2024-01-01T02:20:00.000000078",
			want:  time.Date(2024, 1, 1, 2, 20, 0, 78, time.UTC),
		},
		{
			name:  "ok simple with offset and fraction of second",
			value: "2024-01-01T02:20:00.000000078+02:00",
			want:  time.Date(2024, 1, 1, 2, 20, 0, 78, tzCest),
		},
		{
			name:    "ko wrong date-time separator",
			value:   "2024-01-01 02:20:00",
			wantErr: true,
		},
		{
			name:    "ko 0000 year",
			value:   "0000-04-24T02:20:00",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			xmlInput := fmt.Sprintf("<dt>%s</dt>", tt.value)

			var dt xsdtypes.DateTime

			err := xml.Unmarshal([]byte(xmlInput), &dt)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && !time.Time.Equal(time.Time(dt), tt.want) {
				t.Errorf("UnmarshalXML() got = %v, want %v", dt, tt.want)
			}
		})
	}
}

func TestDateTime_MarshalXML(t *testing.T) {
	t.Parallel()

	tzCest := time.FixedZone("CEST", 2*60*60)
	tzCst := time.FixedZone("CST", -5*60*60)

	tests := []struct {
		name  string
		value time.Time
		want  string
	}{
		{
			name:  "ok utc",
			value: time.Date(2024, 1, 1, 2, 20, 0, 0, time.UTC),
			want:  "<DateTime>2024-01-01T02:20:00Z</DateTime>",
		},
		{
			name:  "ok with offset",
			value: time.Date(2024, 1, 1, 2, 20, 0, 0, tzCest),
			want:  "<DateTime>2024-01-01T02:20:00+02:00</DateTime>",
		},
		{
			name:  "ok with offset negative",
			value: time.Date(2024, 1, 1, 2, 20, 0, 0, tzCst),
			want:  "<DateTime>2024-01-01T02:20:00-05:00</DateTime>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := xml.Marshal(xsdtypes.DateTime(tt.value))
			if err != nil {
				t.Fatalf("MarshalXML() error = %v", err)
			}

			if string(got) != tt.want {
				t.Errorf("MarshalXML() got = %s, want %s", got, tt.want)
			}

			var roundTripped xsdtypes.DateTime

			err = xml.Unmarshal(got, &roundTripped)
			if err != nil {
				t.Fatalf("UnmarshalXML() error = %v", err)
			}

			if !time.Time.Equal(time.Time(roundTripped), tt.value) {
				t.Errorf("round trip got = %v, want %v", time.Time(roundTripped), tt.value)
			}
		})
	}
}

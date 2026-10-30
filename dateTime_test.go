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

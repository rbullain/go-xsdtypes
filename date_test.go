package xsdtypes_test

import (
	"encoding/xml"
	"fmt"
	"testing"
	"time"

	"github.com/rbullain/go-xsdtypes"
)

func TestDate_UnmarshalXML(t *testing.T) {
	t.Parallel()

	tzCest := time.FixedZone("CEST", 2*60*60)

	tests := []struct {
		name    string
		value   string
		want    time.Time
		wantErr bool
	}{
		{
			name:  "ok simple",
			value: "1970-04-24",
			want:  time.Date(1970, 4, 24, 0, 0, 0, 0, time.UTC),
		},
		{
			name:  "ok utc",
			value: "1970-04-24Z",
			want:  time.Date(1970, 4, 24, 0, 0, 0, 0, time.UTC),
		},
		{
			name:  "ok negative year",
			value: "-1970-04-24",
			want:  time.Date(-1970, 4, 24, 0, 0, 0, 0, time.UTC),
		},
		{
			name:  "ok with offset",
			value: "1970-04-24+02:00",
			want:  time.Date(1970, 4, 24, 0, 0, 0, 0, tzCest),
		},
		{
			name:    "ko with fraction of second",
			value:   "1970-04-24.000000078",
			wantErr: true,
		},
		{
			name:    "ko wrong format",
			value:   "12:34:56",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			xmlInput := fmt.Sprintf("<dt>%s</dt>", tt.value)

			var dt xsdtypes.Date
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

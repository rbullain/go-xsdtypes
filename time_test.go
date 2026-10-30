package xsdtypes_test

import (
	"encoding/xml"
	"fmt"
	"testing"
	"time"

	"github.com/rbullain/go-xsdtypes"
)

func TestTime_UnmarshalXML(t *testing.T) {
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
			value: "12:34:56",
			want:  time.Date(0, 1, 1, 12, 34, 56, 0, time.UTC),
		},
		{
			name:  "ok utc",
			value: "12:34:56Z",
			want:  time.Date(0, 1, 1, 12, 34, 56, 0, time.UTC),
		},
		{
			name:  "ok with offset",
			value: "12:34:56+02:00",
			want:  time.Date(0, 1, 1, 12, 34, 56, 0, tzCest),
		},
		{
			name:  "ok with fraction of second",
			value: "12:34:56.000000078",
			want:  time.Date(0, 1, 1, 12, 34, 56, 78, time.UTC),
		},
		{
			name:  "ok with fraction of second and offset",
			value: "12:34:56.000000078+02:00",
			want:  time.Date(0, 1, 1, 12, 34, 56, 78, tzCest),
		},
		{
			name:    "ko wrong time",
			value:   "34:34:56",
			wantErr: true,
		},
		{
			name:    "ko wrong format",
			value:   "2024-01-01T12:34:56",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			xmlInput := fmt.Sprintf("<tm>%s</tm>", tt.value)

			var tm xsdtypes.Time

			err := xml.Unmarshal([]byte(xmlInput), &tm)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && !time.Time.Equal(time.Time(tm), tt.want) {
				t.Errorf("UnmarshalXML() got = %v, want %v", tm, tt.want)
			}
		})
	}
}

func Test_DecodeError(t *testing.T) {
	t.Parallel()

	var tm xsdtypes.Time

	err := xml.Unmarshal([]byte("<dt>2024-01-01</tm>"), &tm)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}

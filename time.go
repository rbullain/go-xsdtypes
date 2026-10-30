package xsdtypes

import (
	"encoding/xml"
	"errors"
	"time"
)

// Time as defined in https://www.w3.org/TR/xmlschema-2/#time
type Time time.Time

var (
	_ xml.Marshaler   = (*Time)(nil)
	_ xml.Unmarshaler = (*Time)(nil)
)

// timeMarshalLayout defines the layout for marshaling Time to XML.
const timeMarshalLayout = "15:04:05"

func (t Time) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	formattedTime := time.Time(t).Format(timeMarshalLayout)
	return encoder.EncodeElement(formattedTime, start)
}

func (t *Time) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	// contains the possible layouts used for unmarshaling Time from XML.
	timeUnmarshalLayouts := [...]string{
		"15:04:05",
		"15:04:05Z07:00",
	}

	parsedTime, err := decodeXMLTime(decoder, start, timeUnmarshalLayouts[:])
	if err != nil {
		return err
	}

	*t = Time(parsedTime)

	return nil
}

func decodeXMLTime(decoder *xml.Decoder, start xml.StartElement, layouts []string) (time.Time, error) {
	var val string
	if err := decoder.DecodeElement(&val, &start); err != nil {
		return time.Time{}, err
	}

	for _, layout := range layouts {
		parsed, err := time.Parse(layout, val)
		if err == nil {
			return parsed, nil
		}
	}

	return time.Time{}, errors.New("invalid time format: " + val)
}

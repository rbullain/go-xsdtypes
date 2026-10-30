package xsdtypes

import (
	"encoding/xml"
	"time"
)

// Date as defined in https://www.w3.org/TR/xmlschema-2/#date
type Date time.Time

var (
	_ xml.Marshaler   = (*Date)(nil)
	_ xml.Unmarshaler = (*Date)(nil)
)

// dateMarshalLayout defines the layout for marshaling Date to XML.
const dateMarshalLayout = "2006-01-02"

func (d Date) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	formattedDate := time.Time(d).Format(dateMarshalLayout)
	return encoder.EncodeElement(formattedDate, start)
}

func (d *Date) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	// contains the possible layouts used for unmarshaling Date from XML.
	dateUnmarshalLayouts := [...]string{
		"2006-01-02",
		"2006-01-02Z07:00",
	}

	parsedTime, err := decodeXMLTime(decoder, start, dateUnmarshalLayouts[:])
	if err != nil {
		return err
	}

	*d = Date(parsedTime)

	return nil
}

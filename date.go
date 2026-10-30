package xsdtypes

import (
	"encoding/xml"
	"errors"
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

var dateUnmarshalLayouts = []string{
	"2006-01-02",
	"2006-01-02Z07:00",
}

func (d Date) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	formattedDate := time.Time(d).Format(dateMarshalLayout)
	return encoder.EncodeElement(formattedDate, start)
}

func (d *Date) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var val string
	if err := decoder.DecodeElement(&val, &start); err != nil {
		return err
	}

	var parsedDate time.Time
	var err error

	for _, layout := range dateUnmarshalLayouts {
		parsedDate, err = time.Parse(layout, val)
		if err == nil {
			*d = Date(parsedDate)
			return nil
		}
	}

	return errors.New("invalid date format: " + val)
}

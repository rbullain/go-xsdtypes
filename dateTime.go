package xsdtypes

import (
	"encoding/xml"
	"time"
)

// DateTime as defined in https://www.w3.org/TR/xmlschema-2/#dateTime
type DateTime time.Time

// dateTimeMarshalLayout defines the layout for marshaling DateTime to XML.
const dateTimeMarshalLayout = "15:04:05"

func (dt DateTime) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	formattedTime := time.Time(dt).Format(dateTimeMarshalLayout)
	return encoder.EncodeElement(formattedTime, start)
}

func (dt *DateTime) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	// contains the possible layouts used for unmarshaling DateTime from XML.
	dateTimeUnmarshalLayouts := [...]string{
		"2006-01-02T15:04:05",
		"2006-01-02T15:04:05Z07:00",
	}

	parsedTime, err := decodeXMLTime(decoder, start, dateTimeUnmarshalLayouts[:])
	if err != nil {
		return err
	}

	*dt = DateTime(parsedTime)

	return nil
}

package xmlfile

import (
	"encoding/xml"
	"fmt"
)

//go:generate xsdgen -ns "" -pkg xmlfile ../compound.xsd

var (
	_ xml.UnmarshalerAttr = (*DoxBool)(nil)
)

type DoxBool bool

func (v *DoxBool) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "no":
		*v = false
	case "yes":
		*v = true
	default:
		return fmt.Errorf("unexpected bool value: %q", attr.Value)
	}
	return nil
}

func (v *DoxBool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	switch s {
	case "no":
		*v = false
	case "yes":
		*v = true
	default:
		return fmt.Errorf("unexpected bool value: %q", s)
	}
	return nil
}

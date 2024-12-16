package order2cash

import (
	"encoding/xml"
	"time"
)

type Date struct {
	time.Time
}

func (d Date) MarshalSchema() string {
	return d.Time.Format("2006-01-02")
}

func (t *Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	s := ""
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}

	if s == "" {
		return nil
	}

	layout := "2006-01-02"
	nt, err := time.Parse(layout, s)
	if err == nil {
		*t = Date{Time: nt}
	}

	layout = "2006-01-02T15:04:05"
	nt, err = time.Parse(layout, s)
	if err == nil {
		*t = Date{Time: nt}
	}

	return err
}

func (t Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	layout := "2006-01-02"
	s := t.Format(layout)
	return e.EncodeElement(s, start)
}

type DateTime struct {
	time.Time
}

func (t *DateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	s := ""
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}

	if s == "" {
		return nil
	}

	layout := "2006-01-02T15:04:05"
	nt, err := time.Parse(layout, s)
	*t = DateTime{Time: nt}
	return err
}

func (t DateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	layout := "2006-01-02T15:04:05"
	s := t.Format(layout)
	return e.EncodeElement(s, start)
}

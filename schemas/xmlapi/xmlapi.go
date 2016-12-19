package xmlapi

import (
	"encoding/xml"
	"time"
)

const TimeFormat string = "2006-01-02 15:04:05"

type (
	EveApi struct {
		Version     int   `xml:"version,attr"`
		CurrentTime Time  `xml:"currentTime"`
		Error       Error `xml:"error,omitempty"`
		CachedUntil Time  `xml:"cachedUntil"`
	}

	Time struct {
		time.Time
	}

	DataTime struct {
		DataTime Time `xml:"result>dataTime"`
	}

	Error struct {
		Code int    `xml:"code,attr"`
		Text string `xml:",chardata"`
	}
)

func (e *EveApi) GetCachedFor() (d time.Duration) {
	var t = time.Now()
	if e.CachedUntil.After(t) {
		d = e.CachedUntil.Sub(t)
	}
	return
}

func (e *Time) UnmarshalXMLAttr(attr xml.Attr) (err error) {
	var t time.Time
	if attr.Value == "" {
		return
	}
	if t, err = time.Parse(TimeFormat, attr.Value); err != nil {
		return
	}
	*e = Time{t}
	return
}

func (e *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var t time.Time
	var v string
	if err = d.DecodeElement(&v, &start); err != nil {
		return
	}
	if t, err = time.Parse(TimeFormat, v); err != nil {
		return
	}
	*e = Time{t}
	return
}

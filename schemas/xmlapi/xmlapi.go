package xmlapi

import (
	"encoding/xml"
	"time"
)

const timeFormat string = "2006-01-02 15:04:05"

type (
	// EveAPI represents the root document
	EveAPI struct {
		Version     int   `xml:"version,attr"`
		CurrentTime Time  `xml:"currentTime"`
		Error       Error `xml:"error,omitempty"`
		CachedUntil Time  `xml:"cachedUntil"`
	}

	// Time represents API datetime
	Time struct {
		time.Time
	}

	// DataTime represents specific datetime for some API data
	DataTime struct {
		DataTime Time `xml:"result>dataTime"`
	}

	// Error represents an error message
	Error struct {
		Code int    `xml:"code,attr"`
		Text string `xml:",chardata"`
	}
)

// GetCachedFor returns the duration for which the request is cached on
// CCP's servers
func (e *EveAPI) GetCachedFor() (d time.Duration) {
	var t = time.Now()
	if e.CachedUntil.After(t) {
		d = e.CachedUntil.Sub(t)
	}
	return
}

// UnmarshalXMLAttr TODO
func (e *Time) UnmarshalXMLAttr(attr xml.Attr) (err error) {
	var t time.Time
	if attr.Value == "" {
		return
	}
	if t, err = time.Parse(timeFormat, attr.Value); err != nil {
		return
	}
	*e = Time{t}
	return
}

// UnmarshalXML TODO
func (e *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var t time.Time
	var v string
	if err = d.DecodeElement(&v, &start); err != nil {
		return
	}
	if t, err = time.Parse(timeFormat, v); err != nil {
		return
	}
	*e = Time{t}
	return
}

package order2cash

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	"github.com/cydev/zero"
	"github.com/omniboost/go-order2cash/omitempty"
)

type RequestEnvelope struct {
	XMLName xml.Name
	// Namespaces []xml.Attr `xml:"-"`

	NS     []xml.Attr `xml:"-"`
	Header SOAPHeader `xml:"soap:Header"`
	Body   any        `xml:"soap:Body"`
}

func NewRequestEnvelope(username, password string) RequestEnvelope {
	return RequestEnvelope{
		NS: []xml.Attr{
			{Name: xml.Name{Space: "", Local: "xmlns:soap"}, Value: "http://www.w3.org/2003/05/soap-envelope"},
			{Name: xml.Name{Space: "", Local: "xmlns:ns"}, Value: "http://schemas.invoiceportal.com/invoiceportal/ws/2.0"},
			{Name: xml.Name{Space: "", Local: "xmlns:ns5"}, Value: "http://schemas.invoiceportal.com/invoiceportal/ws/2.0"},
			{Name: xml.Name{Space: "", Local: "xmlns:ns3"}, Value: "http://schemas.anachron.com/ingis/ws/invoice/1.0"},
			{Name: xml.Name{Space: "", Local: "xmlns:ns4"}, Value: "http://schemas.anachron.com/ingis/ws/invoice/1.1"},
		},
		Header: NewSOAPHeader(username, password),
	}
}

type ResponseEnvelope struct {
	XMLName xml.Name

	Header SOAPHeader `xml:"Header"`
	Body   any        `xml:"Body"`
}

func (env RequestEnvelope) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "soap:Envelope"}

	for _, ns := range env.NS {
		start.Attr = append(start.Attr, ns)
	}

	type alias RequestEnvelope
	a := alias(env)
	return e.EncodeElement(a, start)
}

type SOAPHeader struct {
	Security Security `xml:"wsse:Security"`
}

func (h SOAPHeader) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(h, e, start)
}

func (h SOAPHeader) IsEmpty() bool {
	return zero.IsZero(h)
}

func NewSOAPHeader(username, password string) SOAPHeader {
	timestamp := strings.Replace(time.Now().Format("20060102150405.000"), ".", "", 1)
	usernametoken := fmt.Sprintf("UsernameToken-%s", timestamp)

	return SOAPHeader{
		Security: Security{
			MustUnderstand: true,
			WSSENamespace:  "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd",
			UsernameToken: UsernameToken{
				ID:           usernametoken,
				WSUNamespace: "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd",
				Username:     username,
				Password: Password{
					Type:    "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText",
					Content: password,
				},
			},
		},
	}
}

type Security struct {
	XMLName        xml.Name      `xml:"wsse:Security"`
	MustUnderstand bool          `xml:"soap:mustUnderstand,attr"`
	UsernameToken  UsernameToken `xml:"wsse:UsernameToken"`
	WSSENamespace  string        `xml:"xmlns:wsse,attr"`
}

type UsernameToken struct {
	ID           string   `xml:"wsu:Id,attr"`
	Username     string   `xml:"wsse:Username"`
	Password     Password `xml:"wsse:Password"`
	WSUNamespace string   `xml:"xmlns:wsu,attr"`
}

type Password struct {
	Type    string `xml:"Type,attr"`
	Content string `xml:",chardata"`
}

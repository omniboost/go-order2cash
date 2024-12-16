package order2cash

import (
	"encoding/base64"
	"encoding/xml"
)

type JCHeader struct {
	BucketType    string `xml:"BucketType"`
	APIType       string `xml:"APIType"`
	APIVersion    string `xml:"APIVersion"`
	SecurityToken string `xml:"SecurityToken"`
}

type JCParameters struct {
	// A unique code for the partner system
	// Provided by Jonas Chorum
	PartnerCode string `xml:"PartnerCode"`

	// A unique code for the hotel
	// Used by webservice layer to validate customer is active and determine where message needs to be sent
	// Provided by Jonas Chorum
	HotelCode string `xml:"HotelCode"`

	// Provides access privileges to connected PMS systems
	// Provided by Jonas Chorum
	PartnerToken string `xml:"PartnerToken"`
	EchoToken    string `xml:"EchoToken,omitempty"`

	Error string `xml:"Error"`
}

type OCHeader struct {
	UsernameToken string `xml:"UsernameToken"`
	Username      string `xml:"Username"`
	Password      string `xml:"Password"`
}

// type Security struct {
// 	XMLName        xml.Name      `xml:"wsse:Security"`
// 	MustUnderstand bool          `xml:"soap:mustUnderstand,attr"`
// 	UsernameToken  UsernameToken `xml:"wsse:UsernameToken"`
// 	WSSENamespace  string        `xml:"xmlns:wsse,attr"`
// }

// type UsernameToken struct {
// 	ID           string   `xml:"wsu:Id,attr"`
// 	Username     string   `xml:"wsse:Username"`
// 	Password     Password `xml:"wsse:Password"`
// 	WSUNamespace string   `xml:"xmlns:wsu,attr"`
// }

// type Password struct {
// 	Type    string `xml:"Type,attr"`
// 	Content string `xml:",chardata"`
// }

// <!DOCTYPE MsiXmlBucket []>
// <Content>
//   <Header>
//     <BucketType>ProcessRequestBucketError</BucketType>
//     <APIType>NA</APIType>
//     <APIVersion>NA</APIVersion>
//     <SecurityToken>NA</SecurityToken>
//     <Internal>NA</Internal>
//     <CustomDataA></CustomDataA>
//     <CustomDataB></CustomDataB>
//     <CustomDataC></CustomDataC>
//     <CustomDataD></CustomDataD>
//   </Header>
//   <Parameters>
//     <Error>Value cannot be null.
// Parameter name: s</Error>
//   </Parameters>
//   <Body />
// </Content>

type Base64Binary []byte

func (b Base64Binary) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s := base64.StdEncoding.EncodeToString([]byte(b))
	return e.EncodeElement(s, start)
}

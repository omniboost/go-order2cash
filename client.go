package order2cash

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strings"
	"text/template"

	"github.com/elliotchance/pie/v2"
	"github.com/pkg/errors"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-order2cash/" + libraryVersion
	mediaType      = "application/soap+xml"
	charset        = "utf-8"
)

var (
	BaseURL = url.URL{
		Scheme: "https",
		Host:   "stage.anachron.com",
		Path:   "/i2d-gw/services",
	}
)

// NewClient returns a new Exact Globe Client client
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{}

	client.SetHTTPClient(httpClient)
	client.SetBaseURL(BaseURL)
	client.SetDebug(false)
	client.SetUserAgent(userAgent)
	client.SetMediaType(mediaType)
	client.SetCharset(charset)

	return client
}

// Client manages communication with Exact Globe Client
type Client struct {
	// HTTP client used to communicate with the Client.
	http *http.Client

	debug   bool
	baseURL url.URL

	// credentials
	username string
	password string
	// partnerCode  string

	// User agent for client
	userAgent string

	mediaType string
	charset   string

	// Optional function called after every successful request made to the DO Clients
	onRequestCompleted RequestCompletionCallback
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func (c *Client) SetHTTPClient(client *http.Client) {
	c.http = client
}

func (c Client) Debug() bool {
	return c.debug
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c Client) Username() string {
	return c.username
}

func (c *Client) SetUsername(username string) {
	c.username = username
}

func (c Client) Password() string {
	return c.password
}

func (c *Client) SetPassword(password string) {
	c.password = password
}

// func (c Client) PartnerCode() string {
// 	return c.partnerCode
// }

// func (c *Client) SetPartnerCode(partnerCode string) {
// 	c.partnerCode = partnerCode
// }

func (c Client) BaseURL() url.URL {
	return c.baseURL
}

func (c *Client) SetBaseURL(baseURL url.URL) {
	c.baseURL = baseURL
}

func (c *Client) SetMediaType(mediaType string) {
	c.mediaType = mediaType
}

func (c Client) MediaType() string {
	return mediaType
}

func (c *Client) SetCharset(charset string) {
	c.charset = charset
}

func (c Client) Charset() string {
	return charset
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c Client) UserAgent() string {
	return userAgent
}

func (c *Client) GetEndpointURL(p string, pathParams PathParams) url.URL {
	clientURL := c.BaseURL()

	parsed, err := url.Parse(p)
	if err != nil {
		log.Fatal(err)
	}
	q := clientURL.Query()
	for k, vv := range parsed.Query() {
		for _, v := range vv {
			q.Add(k, v)
		}
	}
	clientURL.RawQuery = q.Encode()

	clientURL.Path = path.Join(clientURL.Path, parsed.Path)

	tmpl, err := template.New("path").Parse(clientURL.Path)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	params := pathParams.Params()
	// params["administration_id"] = c.Administration()
	err = tmpl.Execute(buf, params)
	if err != nil {
		log.Fatal(err)
	}

	clientURL.Path = buf.String()
	return clientURL
}

func (c *Client) NewRequest(ctx context.Context, req Request) (*http.Request, error) {
	// convert body struct to xml
	buf := new(bytes.Buffer)
	if req.RequestBodyInterface() != nil {
		// soapRequest := RequestEnvelope{
		// 	Namespaces: []xml.Attr{
		// 		{Name: xml.Name{Space: "", Local: "xmlns:soap"}, Value: "http://schemas.xmlsoap.org/soap/envelope/"},
		// 	},
		// 	Header: req.RequestHeaderInterface(),
		// 	Body: Body{
		// 		ActionBody: req.RequestBodyInterface(),
		// 	},
		// }

		soapRequest := NewRequestEnvelope(c.username, c.password)
		soapRequest.Body = req.RequestBodyInterface()

		enc := xml.NewEncoder(buf)
		enc.Indent("", "  ")
		err := enc.Encode(soapRequest)
		if err != nil {
			return nil, err
		}

		err = enc.Flush()
		if err != nil {
			return nil, err
		}
	}

	// create new http request
	r, err := http.NewRequest(req.Method(), req.URL().String(), buf)
	if err != nil {
		return nil, err
	}

	// values := url.Values{}
	// err = utils.AddURLValuesToRequest(values, req, true)
	// if err != nil {
	// 	return nil, err
	// }

	// optionally pass along context
	if ctx != nil {
		r = r.WithContext(ctx)
	}

	// set other headers
	r.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", c.MediaType(), c.Charset()))
	r.Header.Add("Accept", c.MediaType())
	r.Header.Add("User-Agent", c.UserAgent())
	// r.Header.Add("SOAPAction", req.SOAPAction())

	return r, nil
}

// Do sends an Client request and returns the Client response. The Client response is xml decoded and stored in the value
// pointed to by v, or returned as an error if an Client error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, body interface{}) (*http.Response, error) {
	if c.debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	httpResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, nil
	}

	if body == nil {
		return httpResp, err
	}

	mediaType, params, err := mime.ParseMediaType(httpResp.Header.Get("Content-Type"))
	if err != nil {
		return httpResp, err
	}

	mr := multipart.NewReader(httpResp.Body, params["boundary"])

	parts := []io.Reader{}
	if strings.HasPrefix(mediaType, "multipart/") {
		for {
			part, err := mr.NextPart()
			if errors.Is(err, io.EOF) {
				break
			}
			if err != nil {
				return httpResp, err
			}

			buf := new(bytes.Buffer)
			_, err = io.Copy(buf, part) // Fully consume the part
			if err != nil {
				return httpResp, err
			}

			parts = append(parts, buf)
			part.Close()
		}
	} else {
		parts = append(parts, httpResp.Body)
	}

	if len(parts) == 0 {
		return httpResp, errors.New("No body/parts found in the http response")
	} else if len(parts) > 1 {
		return httpResp, errors.New("More then 1 body/part found in the http response: we don't support this yet")
	}

	soapResponse := &ResponseEnvelope{
		// header: SOAPHeader{},
		Body: body,
	}
	soapError := SOAPError{Response: httpResp}
	errResp := &ResponseEnvelope{
		Header: SOAPHeader{},
		Body:   &soapError,
	}

	soapFault := SOAPFault{Response: httpResp}
	faultResp := &ResponseEnvelope{
		Header: SOAPHeader{},
		Body:   &soapFault,
	}

	statusResponseBody := StatusResponseBody{Response: httpResp}
	statusResp := &ResponseEnvelope{
		Header: SOAPHeader{},
		Body:   statusResponseBody,
	}

	err = c.Unmarshal(parts[0], soapResponse, errResp, faultResp, statusResp)
	if err != nil {
		return httpResp, err
	}

	if e, ok := soapResponse.Body.(error); ok {
		if e.Error() != "" {
			return httpResp, e
		}
	}

	if soapError.Error() != "" {
		return httpResp, soapError
	}

	if soapFault.Error() != "" {
		return httpResp, soapFault
	}

	// soapResponse := ResponseEnvelope{
	// 	// Header: Header{},
	// 	Body: Body{
	// 		ActionBody: body,
	// 	},
	// }

	// soapError := SoapError{}
	// err = c.Unmarshal(httpResp.Body, []interface{}{&soapResponse}, []interface{}{&soapError})
	// if err != nil {
	// 	return httpResp, err
	// }

	// if soapError.Error() != "" {
	// 	return httpResp, &ErrorResponse{Response: httpResp, Err: soapError}
	// }

	// i, ok := body.(interface{ ExceptionBlock() ExceptionBlock })
	// if ok {
	// 	if i.ExceptionBlock().ExceptionCode != 0 {
	// 		return httpResp, &ErrorResponse{Response: httpResp, Err: i.ExceptionBlock()}
	// 	}
	// }

	// if len(errorResponse.Messages) > 0 {
	// 	return httpResp, errorResponse
	// }

	return httpResp, nil
}

// func (c *Client) Unmarshal(r io.Reader, vv []interface{}, optionalVv []interface{}) error {
// 	if len(vv) == 0 && len(optionalVv) == 0 {
// 		return nil
// 	}

// 	b, err := ioutil.ReadAll(r)
// 	if err != nil {
// 		return err
// 	}

// 	for _, v := range vv {
// 		r := bytes.NewReader(b)
// 		dec := xml.NewDecoder(r)

// 		err := dec.Decode(v)
// 		if err != nil && err != io.EOF {
// 			return errors.WithStack((err))
// 		}
// 	}

// 	for _, v := range optionalVv {
// 		r := bytes.NewReader(b)
// 		dec := xml.NewDecoder(r)

// 		_ = dec.Decode(v)
// 	}

// 	return nil
// }

func (c *Client) Unmarshal(r io.Reader, vv ...interface{}) error {
	if len(vv) == 0 {
		return nil
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	errs := []error{}
	for _, v := range vv {
		r := bytes.NewReader(b)
		dec := xml.NewDecoder(r)

		err := dec.Decode(v)
		if err != nil && err != io.EOF {
			errs = append(errs, err)
		}
	}

	if len(errs) == len(vv) {
		// Everything errored
		msgs := make([]string, len(errs))
		for i, e := range errs {
			log.Println(e)
			msgs[i] = fmt.Sprint(e)
		}
		return errors.New(strings.Join(msgs, ", "))
	}

	return nil
}

// CheckResponse checks the Client response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. Client error responses are expected to have either no response
// body, or a xml response body that maps to ErrorResponse. Any other response
// body will be silently ignored.
func CheckResponse(r *http.Response) error {
	errorResponse := &ErrorResponse{Response: r}

	// Don't check content-lenght: a created response, for example, has no body
	// if r.Header.Get("Content-Length") == "0" {
	// 	errorResponse.Errors.Message = r.Status
	// 	return errorResponse
	// }

	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	// read data and copy it back
	data, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(data))
	if err != nil {
		return errorResponse
	}

	err = checkContentType(r)
	if err != nil {
		return errors.WithStack(err)
	}

	if r.ContentLength == 0 {
		return errors.New("response body is empty")
	}

	// convert xml to struct
	err = xml.Unmarshal(data, &errorResponse)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func checkContentType(response *http.Response) error {
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != mediaType {
		return fmt.Errorf("Expected Content-Type \"%s\", got \"%s\"", mediaType, contentType)
	}

	return nil
}

type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response
	Err      string
}

func (r *ErrorResponse) Error() string {
	return r.Err
}

type SOAPError struct {
	// HTTP response that caused this error
	Response *http.Response
}

func (e SOAPError) Error() string {
	return ""
}

type SOAPFault struct {
	// HTTP response that caused this error
	Response *http.Response

	Fault struct {
		Code struct {
			Value string `xml:"Value"`
		} `xml:"Code"`
		Reason struct {
			Text struct {
				Lang    string `xml:"xml:lang,attr"`
				Content string `xml:",chardata"`
			} `xml:"Text"`
		} `xml:"Reason"`
	} `xml:"Fault"`
}

func (f SOAPFault) Error() string {
	l := []string{f.Fault.Code.Value, f.Fault.Reason.Text.Content}
	if f.Fault.Code.Value != "" {
		l = append(l, f.Fault.Code.Value)
	}
	l = append(l, f.Fault.Reason.Text.Content)

	ll := []string{}
	for _, v := range l {
		if v != "" {
			ll = append(ll, v)
		}
	}

	ll = pie.Unique(pie.Strings(ll))

	// ll = pie.Strings(ll).Unique()

	return strings.Join(ll, ", ")
}

type StatusResponseBody struct {
	// HTTP response that caused this error
	Response *http.Response

	Node struct {
	} `xml:",any"`
}

// <?xml version="1.0" encoding="utf-8"?>
// <soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
//   <soap:Body>
//     <soap:Fault>
//       <faultcode>soap:Client</faultcode>
//       <faultstring>Server was unable to read request. ---&gt; There is an error in XML document (5, 34). ---&gt; Input string was not in a correct format.</faultstring>
//       <detail/>
//     </soap:Fault>
//   </soap:Body>
// </soap:Envelope>

// type SoapError struct {
// 	XMLName xml.Name `xml:"Envelope"`
// 	Body    struct {
// 		Fault struct {
// 			FaultCode   string `xml:"faultcode"`
// 			FaultString string `xml:"faultstring"`
// 		} `xml:"Fault"`
// 	} `xml:"Body"`
// }

// func (e SoapError) Error() string {
// 	if e.Body.Fault.FaultCode != "" || e.Body.Fault.FaultString != "" {
// 		return fmt.Sprintf("%s: %s", e.Body.Fault.FaultCode, e.Body.Fault.FaultString)
// 	}
// 	return ""
// }

// type SoapError struct {
// 	XMLName xml.Name `xml:"Envelope"`
// 	Body    struct {
// 		Fault struct {
// 			Code struct {
// 				Value string `xml:"Value"`
// 			} `xml:"Code"`
// 			Reason struct {
// 				Text struct {
// 					Lang    string `xml:"xml:lang,attr"`
// 					Content string `xml:",chardata"`
// 				} `xml:"Text"`
// 			} `xml:"Reason"`
// 		} `xml:"Fault"`
// 	} `xml:"Body"`
// }

// func (e SoapError) Error() string {
// 	if e.Body.Fault.Code.Value != "" || e.Body.Fault.Reason.Text.Content != "" {
// 		return fmt.Sprintf("%s: %s", e.Body.Fault.Code.Value, e.Body.Fault.Reason.Text.Content)
// 	}
// 	return ""
// }

// type ErrorResponse struct {
// 	// HTTP response that caused this error
// 	Response *http.Response

// 	// HTTP status code
// 	Err error
// }

// func (r *ErrorResponse) Error() string {
// 	return r.Err.Error()
// }

// type StatusResponseBody struct {
// 	Response *http.Response
// 	Node     struct{} `xml:",any"`
// }

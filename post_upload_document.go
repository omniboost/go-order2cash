package order2cash

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-order2cash/utils"
)

func (c *Client) NewPostUploadDocumentRequest() PostUploadDocumentRequest {
	r := PostUploadDocumentRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	r.requestHeader = r.NewRequestHeader()
	return r
}

type PostUploadDocumentRequest struct {
	client        *Client
	queryParams   *PostUploadDocumentQueryParams
	pathParams    *PostUploadDocumentPathParams
	method        string
	headers       http.Header
	requestBody   PostUploadDocumentRequestBody
	requestHeader PostUploadDocumentRequestHeader
}

func (r PostUploadDocumentRequest) NewQueryParams() *PostUploadDocumentQueryParams {
	return &PostUploadDocumentQueryParams{}
}

type PostUploadDocumentQueryParams struct {
}

func (p PostUploadDocumentQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PostUploadDocumentRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r PostUploadDocumentRequest) NewPathParams() *PostUploadDocumentPathParams {
	return &PostUploadDocumentPathParams{}
}

type PostUploadDocumentPathParams struct {
}

func (p *PostUploadDocumentPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostUploadDocumentRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *PostUploadDocumentRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostUploadDocumentRequest) Method() string {
	return r.method
}

func (r PostUploadDocumentRequest) NewRequestHeader() PostUploadDocumentRequestHeader {
	return PostUploadDocumentRequestHeader{}
}

func (r *PostUploadDocumentRequest) RequestHeader() *PostUploadDocumentRequestHeader {
	return &r.requestHeader
}

func (r *PostUploadDocumentRequest) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type PostUploadDocumentRequestHeader struct {
	// AuthenticationContext AuthenticationContext `xml:"AuthenticationContext"`
}

func (r PostUploadDocumentRequest) NewRequestBody() PostUploadDocumentRequestBody {
	return PostUploadDocumentRequestBody{}
}

type PostUploadDocumentRequestBody struct {
	// XMLName   xml.Name `xml:"rlx:pmsdoc_PostUploadDocument"`
	// SessionID string   `xml:"rlx:SessionID"`
	UploadDocumentRequest UploadDocumentRequest `xml:"ns:UploadDocumentRequest"`
}

func (r *PostUploadDocumentRequest) RequestBody() *PostUploadDocumentRequestBody {
	return &r.requestBody
}

func (r *PostUploadDocumentRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostUploadDocumentRequest) SetRequestBody(body PostUploadDocumentRequestBody) {
	r.requestBody = body
}

func (r *PostUploadDocumentRequest) NewResponseBody() *PostUploadDocumentResponseBody {
	return &PostUploadDocumentResponseBody{}
}

type PostUploadDocumentResponseBody struct {
	XMLName                xml.Name `xml:"Body"`
	UploadDocumentResponse struct {
		XMLName       xml.Name `xml:"UploadDocumentResponse"`
		NumberOfFiles int      `xml:"numberOfFiles,attr"`
		Result        string   `xml:"result,attr"`
		Message       string   `xml:"Message"`
		Xmlns         string   `xml:"xmlns,attr"`
	}
}

func (r *PostUploadDocumentResponseBody) Error() string {
	if r.UploadDocumentResponse.Result == "error" {
		return r.UploadDocumentResponse.Message
	}
	return ""
}

func (r *PostUploadDocumentRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *PostUploadDocumentRequest) Do() (PostUploadDocumentResponseBody, error) {
	var err error

	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}

type UploadDocumentRequest struct {
	XMLName        xml.Name       `xml:"ns:UploadDocumentRequest"`
	SenderID       string         `xml:"ns:senderId,attr"`
	GUID           string         `xml:"ns:guid,attr"`
	XMLFile        Base64Binary   `xml:"ns:XmlFile"`
	AttachmentFile []Base64Binary `xml:"ns:AttachmentFile,omitempty"`
}

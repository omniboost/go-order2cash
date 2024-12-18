package order2cash

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-order2cash/utils"
)

func (c *Client) NewGetUploadDocumentRequest() GetUploadDocumentRequest {
	r := GetUploadDocumentRequest{
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

type GetUploadDocumentRequest struct {
	client        *Client
	queryParams   *GetUploadDocumentQueryParams
	pathParams    *GetUploadDocumentPathParams
	method        string
	headers       http.Header
	requestBody   GetUploadDocumentRequestBody
	requestHeader GetUploadDocumentRequestHeader
}

func (r GetUploadDocumentRequest) NewQueryParams() *GetUploadDocumentQueryParams {
	return &GetUploadDocumentQueryParams{}
}

type GetUploadDocumentQueryParams struct {
}

func (p GetUploadDocumentQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetUploadDocumentRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetUploadDocumentRequest) NewPathParams() *GetUploadDocumentPathParams {
	return &GetUploadDocumentPathParams{}
}

type GetUploadDocumentPathParams struct {
}

func (p *GetUploadDocumentPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetUploadDocumentRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetUploadDocumentRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetUploadDocumentRequest) Method() string {
	return r.method
}

func (r GetUploadDocumentRequest) NewRequestHeader() GetUploadDocumentRequestHeader {
	return GetUploadDocumentRequestHeader{}
}

func (r *GetUploadDocumentRequest) RequestHeader() *GetUploadDocumentRequestHeader {
	return &r.requestHeader
}

func (r *GetUploadDocumentRequest) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type GetUploadDocumentRequestHeader struct {
	// AuthenticationContext AuthenticationContext `xml:"AuthenticationContext"`
}

func (r GetUploadDocumentRequest) NewRequestBody() GetUploadDocumentRequestBody {
	return GetUploadDocumentRequestBody{}
}

type GetUploadDocumentRequestBody struct {
	// XMLName   xml.Name `xml:"rlx:pmsdoc_GetUploadDocument"`
	// SessionID string   `xml:"rlx:SessionID"`
	UploadDocumentRequest UploadDocumentRequest `xml:"ns:UploadDocumentRequest"`
}

func (r *GetUploadDocumentRequest) RequestBody() *GetUploadDocumentRequestBody {
	return &r.requestBody
}

func (r *GetUploadDocumentRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetUploadDocumentRequest) SetRequestBody(body GetUploadDocumentRequestBody) {
	r.requestBody = body
}

func (r *GetUploadDocumentRequest) NewResponseBody() *GetUploadDocumentResponseBody {
	return &GetUploadDocumentResponseBody{}
}

type GetUploadDocumentResponseBody struct {
	XMLName                xml.Name `xml:"Body"`
	UploadDocumentResponse struct {
		XMLName       xml.Name `xml:"UploadDocumentResponse"`
		NumberOfFiles int      `xml:"numberOfFiles,attr"`
		Result        string   `xml:"result,attr"`
		Message       string   `xml:"Message"`
		Xmlns         string   `xml:"xmlns,attr"`
	}
}

func (r *GetUploadDocumentResponseBody) Error() string {
	if r.UploadDocumentResponse.Result == "error" {
		return r.UploadDocumentResponse.Message
	}
	return ""
}

func (r *GetUploadDocumentRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *GetUploadDocumentRequest) Do() (GetUploadDocumentResponseBody, error) {
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
	Guid           string         `xml:"ns:guid,attr"`
	XmlFile        Base64Binary   `xml:"ns:XmlFile"`
	AttachmentFile []Base64Binary `xml:"ns:AttachmentFile"`
}

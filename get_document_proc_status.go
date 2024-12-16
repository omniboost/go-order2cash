package order2cash

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-order2cash/utils"
)

func (c *Client) NewGetDocumentProcStatusRequest() GetDocumentProcStatusRequest {
	r := GetDocumentProcStatusRequest{
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

type GetDocumentProcStatusRequest struct {
	client        *Client
	queryParams   *GetDocumentProcStatusQueryParams
	pathParams    *GetDocumentProcStatusPathParams
	method        string
	headers       http.Header
	requestBody   GetDocumentProcStatusRequestBody
	requestHeader GetDocumentProcStatusRequestHeader
}

func (r GetDocumentProcStatusRequest) NewQueryParams() *GetDocumentProcStatusQueryParams {
	return &GetDocumentProcStatusQueryParams{}
}

type GetDocumentProcStatusQueryParams struct {
}

func (p GetDocumentProcStatusQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetDocumentProcStatusRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetDocumentProcStatusRequest) NewPathParams() *GetDocumentProcStatusPathParams {
	return &GetDocumentProcStatusPathParams{}
}

type GetDocumentProcStatusPathParams struct {
}

func (p *GetDocumentProcStatusPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetDocumentProcStatusRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetDocumentProcStatusRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetDocumentProcStatusRequest) Method() string {
	return r.method
}

func (r GetDocumentProcStatusRequest) NewRequestHeader() GetDocumentProcStatusRequestHeader {
	return GetDocumentProcStatusRequestHeader{}
}

func (r *GetDocumentProcStatusRequest) RequestHeader() *GetDocumentProcStatusRequestHeader {
	return &r.requestHeader
}

func (r *GetDocumentProcStatusRequest) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type GetDocumentProcStatusRequestHeader struct {
	// AuthenticationContext AuthenticationContext `xml:"AuthenticationContext"`
}

func (r GetDocumentProcStatusRequest) NewRequestBody() GetDocumentProcStatusRequestBody {
	return GetDocumentProcStatusRequestBody{}
}

type GetDocumentProcStatusRequestBody struct {
	// XMLName   xml.Name `xml:"rlx:pmsdoc_GetDocumentProcStatus"`
	// SessionID string   `xml:"rlx:SessionID"`
	DocumentProcStatusRequest DocumentProcStatusRequest `xml:"ns:CheckDocumentProcessingStatusRequest"`
}

func (r *GetDocumentProcStatusRequest) RequestBody() *GetDocumentProcStatusRequestBody {
	return &r.requestBody
}

func (r *GetDocumentProcStatusRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetDocumentProcStatusRequest) SetRequestBody(body GetDocumentProcStatusRequestBody) {
	r.requestBody = body
}

func (r *GetDocumentProcStatusRequest) NewResponseBody() *GetDocumentProcStatusResponseBody {
	return &GetDocumentProcStatusResponseBody{}
}

type GetDocumentProcStatusResponseBody struct {
	XMLName                               xml.Name `xml:"Body"`
	CheckDocumentProcessingStatusResponse struct {
		XMLName       xml.Name `xml:"CheckDocumentProcessingStatusResponse"`
		NumberOfFiles int      `xml:"numberOfFiles,attr"`
		Result        string   `xml:"result,attr"`
	}
}

func (r *GetDocumentProcStatusRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *GetDocumentProcStatusRequest) Do() (GetDocumentProcStatusResponseBody, error) {
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

type DocumentProcStatusRequest struct {
	XMLName  xml.Name `xml:"ns:CheckDocumentProcessingStatusRequest"`
	SenderID string   `xml:"ns:senderId,attr"`
	Guid     string   `xml:"ns:guid,attr"`
}

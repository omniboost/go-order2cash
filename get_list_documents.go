package order2cash

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-order2cash/utils"
)

func (c *Client) NewGetListDocumentsRequest() GetListDocumentsRequest {
	r := GetListDocumentsRequest{
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

type GetListDocumentsRequest struct {
	client        *Client
	queryParams   *GetListDocumentsQueryParams
	pathParams    *GetListDocumentsPathParams
	method        string
	headers       http.Header
	requestBody   GetListDocumentsRequestBody
	requestHeader GetListDocumentsRequestHeader
}

func (r GetListDocumentsRequest) NewQueryParams() *GetListDocumentsQueryParams {
	return &GetListDocumentsQueryParams{}
}

type GetListDocumentsQueryParams struct {
}

func (p GetListDocumentsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetListDocumentsRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetListDocumentsRequest) NewPathParams() *GetListDocumentsPathParams {
	return &GetListDocumentsPathParams{}
}

type GetListDocumentsPathParams struct {
}

func (p *GetListDocumentsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetListDocumentsRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetListDocumentsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetListDocumentsRequest) Method() string {
	return r.method
}

func (r GetListDocumentsRequest) NewRequestHeader() GetListDocumentsRequestHeader {
	return GetListDocumentsRequestHeader{}
}

func (r *GetListDocumentsRequest) RequestHeader() *GetListDocumentsRequestHeader {
	return &r.requestHeader
}

func (r *GetListDocumentsRequest) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type GetListDocumentsRequestHeader struct {
	// AuthenticationContext AuthenticationContext `xml:"AuthenticationContext"`
}

func (r GetListDocumentsRequest) NewRequestBody() GetListDocumentsRequestBody {
	return GetListDocumentsRequestBody{}
}

type GetListDocumentsRequestBody struct {
	// XMLName   xml.Name `xml:"rlx:pmsdoc_GetListDocuments"`
	// SessionID string   `xml:"rlx:SessionID"`
	ListDocumentsRequest ListDocumentsRequest `xml:"ns:ListDocumentsRequest"`
}

func (r *GetListDocumentsRequest) RequestBody() *GetListDocumentsRequestBody {
	return &r.requestBody
}

func (r *GetListDocumentsRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetListDocumentsRequest) SetRequestBody(body GetListDocumentsRequestBody) {
	r.requestBody = body
}

func (r *GetListDocumentsRequest) NewResponseBody() *GetListDocumentsResponseBody {
	return &GetListDocumentsResponseBody{}
}

type GetListDocumentsResponseBody struct {
	XMLName xml.Name `xml:"ns:ListDocumentsResponse"`

	SenderID               string `xml:"ns:senderId,attr"`
	DocumentId             string `xml:"ns:documentId,attr"`
	DocumentNumber         string `xml:"ns:documentNumber,attr"`
	DocumentDateStart      string `xml:"ns:documentDateStart,attr"`
	DocumentDateEnd        string `xml:"ns:documentDateEnd,attr"`
	ViewStatus             string `xml:"ns:viewStatus,attr"`
	DownloadStatus         string `xml:"ns:downloadStatus,attr"`
	InvoiceDebitCreditCode string `xml:"ns:invoiceDebitCreditCode,attr"`
}

func (r *GetListDocumentsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *GetListDocumentsRequest) Do() (GetListDocumentsResponseBody, error) {
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

type ListDocumentsRequest struct {
	SenderID          string `xml:"ns:senderId,attr"`
	DocumentNumber    string `xml:"ns:documentNumber,attr"`
	DocumentDateStart string `xml:"ns:documentDateStart,attr"`
	DocumentDateEnd   string `xml:"ns:documentDateEnd,attr"`
	ViewStatus        string `xml:"ns:viewStatus,attr"`
	DownloadStatus    string `xml:"ns:downloadStatus,attr"`
}

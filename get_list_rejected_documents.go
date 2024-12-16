package order2cash

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-order2cash/utils"
)

func (c *Client) NewGetListRejectedDocumentsRequest() GetListRejectedDocumentsRequest {
	r := GetListRejectedDocumentsRequest{
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

type GetListRejectedDocumentsRequest struct {
	client        *Client
	queryParams   *GetListRejectedDocumentsQueryParams
	pathParams    *GetListRejectedDocumentsPathParams
	method        string
	headers       http.Header
	requestBody   GetListRejectedDocumentsRequestBody
	requestHeader GetListRejectedDocumentsRequestHeader
}

func (r GetListRejectedDocumentsRequest) NewQueryParams() *GetListRejectedDocumentsQueryParams {
	return &GetListRejectedDocumentsQueryParams{}
}

type GetListRejectedDocumentsQueryParams struct {
}

func (p GetListRejectedDocumentsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetListRejectedDocumentsRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetListRejectedDocumentsRequest) NewPathParams() *GetListRejectedDocumentsPathParams {
	return &GetListRejectedDocumentsPathParams{}
}

type GetListRejectedDocumentsPathParams struct {
}

func (p *GetListRejectedDocumentsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetListRejectedDocumentsRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetListRejectedDocumentsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetListRejectedDocumentsRequest) Method() string {
	return r.method
}

func (r GetListRejectedDocumentsRequest) NewRequestHeader() GetListRejectedDocumentsRequestHeader {
	return GetListRejectedDocumentsRequestHeader{}
}

func (r *GetListRejectedDocumentsRequest) RequestHeader() *GetListRejectedDocumentsRequestHeader {
	return &r.requestHeader
}

func (r *GetListRejectedDocumentsRequest) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type GetListRejectedDocumentsRequestHeader struct {
	// AuthenticationContext AuthenticationContext `xml:"AuthenticationContext"`
}

func (r GetListRejectedDocumentsRequest) NewRequestBody() GetListRejectedDocumentsRequestBody {
	return GetListRejectedDocumentsRequestBody{}
}

type GetListRejectedDocumentsRequestBody struct {
	// XMLName   xml.Name `xml:"rlx:pmsdoc_GetListRejectedDocuments"`
	// SessionID string   `xml:"rlx:SessionID"`
	ListRejectedDocumentsRequest ListRejectedDocumentsRequest `xml:"ns:ListRejectedDocumentsRequest"`
}

func (r *GetListRejectedDocumentsRequest) RequestBody() *GetListRejectedDocumentsRequestBody {
	return &r.requestBody
}

func (r *GetListRejectedDocumentsRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetListRejectedDocumentsRequest) SetRequestBody(body GetListRejectedDocumentsRequestBody) {
	r.requestBody = body
}

func (r *GetListRejectedDocumentsRequest) NewResponseBody() *GetListRejectedDocumentsResponseBody {
	return &GetListRejectedDocumentsResponseBody{}
}

type GetListRejectedDocumentsResponseBody struct {
	XMLName                               xml.Name `xml:"Body"`
	ListRejectedDocumentsResponse struct {
		XMLName       xml.Name `xml:"ListRejectedDocumentsResponse"`
		NumberOfFiles int      `xml:"numberOfFiles,attr"`
		Result        string   `xml:"result,attr"`
	}
}

func (r *GetListRejectedDocumentsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *GetListRejectedDocumentsRequest) Do() (GetListRejectedDocumentsResponseBody, error) {
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

type ListRejectedDocumentsRequest struct {
	SenderID string `xml:"ns:senderId,attr"`
}

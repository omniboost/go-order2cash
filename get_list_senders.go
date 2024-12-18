package order2cash

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-order2cash/utils"
)

func (c *Client) NewGetListSendersRequest() GetListSendersRequest {
	r := GetListSendersRequest{
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

type GetListSendersRequest struct {
	client        *Client
	queryParams   *GetListSendersQueryParams
	pathParams    *GetListSendersPathParams
	method        string
	headers       http.Header
	requestBody   GetListSendersRequestBody
	requestHeader GetListSendersRequestHeader
}

func (r GetListSendersRequest) NewQueryParams() *GetListSendersQueryParams {
	return &GetListSendersQueryParams{}
}

type GetListSendersQueryParams struct {
}

func (p GetListSendersQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetListSendersRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetListSendersRequest) NewPathParams() *GetListSendersPathParams {
	return &GetListSendersPathParams{}
}

type GetListSendersPathParams struct {
}

func (p *GetListSendersPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetListSendersRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *GetListSendersRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetListSendersRequest) Method() string {
	return r.method
}

func (r GetListSendersRequest) NewRequestHeader() GetListSendersRequestHeader {
	return GetListSendersRequestHeader{}
}

func (r *GetListSendersRequest) RequestHeader() *GetListSendersRequestHeader {
	return &r.requestHeader
}

func (r *GetListSendersRequest) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type GetListSendersRequestHeader struct {
	// AuthenticationContext AuthenticationContext `xml:"AuthenticationContext"`
}

func (r GetListSendersRequest) NewRequestBody() GetListSendersRequestBody {
	return GetListSendersRequestBody{}
}

type GetListSendersRequestBody struct {
	XMLName xml.Name `xml:"ns:ListSendersRequest"`
}

func (r *GetListSendersRequest) RequestBody() *GetListSendersRequestBody {
	return &r.requestBody
}

func (r *GetListSendersRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetListSendersRequest) SetRequestBody(body GetListSendersRequestBody) {
	r.requestBody = body
}

func (r *GetListSendersRequest) NewResponseBody() *GetListSendersResponseBody {
	return &GetListSendersResponseBody{}
}

type GetListSendersResponseBody struct {
	XMLName                               xml.Name `xml:"Body"`
	CheckDocumentProcessingStatusResponse struct {
		XMLName       xml.Name `xml:"CheckDocumentProcessingStatusResponse"`
		NumberOfFiles int      `xml:"numberOfFiles,attr"`
		Result        string   `xml:"result,attr"`
		Message       string   `xml:"Message"`
	}
}

func (r *GetListSendersRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *GetListSendersRequest) Do() (GetListSendersResponseBody, error) {
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

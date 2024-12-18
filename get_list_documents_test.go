package order2cash_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetListDocumentsRequest(t *testing.T) {
	// base := order2cash.Base64Binary("test")
	// x, _ := xml.MarshalIndent(base, "", "  ")
	// log.Fatal(string(x))

	req := client.NewGetListDocumentsRequest()
	req.RequestBody().ListDocumentsRequest.SenderID = "bastionairport"
	req.RequestBody().ListDocumentsRequest.DocumentNumber = "test"
	req.RequestBody().ListDocumentsRequest.DocumentDateStart = ""
	req.RequestBody().ListDocumentsRequest.DocumentDateEnd = ""
	req.RequestBody().ListDocumentsRequest.ViewStatus = ""
	req.RequestBody().ListDocumentsRequest.DownloadStatus = ""
	// req.Parameters.PartnerCode = client.PartnerCode()
	// req.Parameters.HotelCode = client.HotelCode()
	// req.Parameters.PartnerToken = client.PartnerToken()
	// req.Body.BusinessDate = "2024-10-01"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

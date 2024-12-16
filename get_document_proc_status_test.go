package order2cash_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetDocumentProcStatusRequest(t *testing.T) {
	// base := order2cash.Base64Binary("test")
	// x, _ := xml.MarshalIndent(base, "", "  ")
	// log.Fatal(string(x))

	req := client.NewGetDocumentProcStatusRequest()
	req.RequestBody().DocumentProcStatusRequest.Guid = ""
	req.RequestBody().DocumentProcStatusRequest.SenderID = "bastionairport"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

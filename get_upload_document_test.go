package order2cash_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetUploadDocumentRequest(t *testing.T) {
	// base := order2cash.Base64Binary("test")
	// x, _ := xml.MarshalIndent(base, "", "  ")
	// log.Fatal(string(x))

	req := client.NewGetUploadDocumentRequest()
	req.RequestBody().UploadDocumentRequest.SenderID = "bastionairport"
	req.RequestBody().UploadDocumentRequest.Guid = "bastionairport"
	// req.RequestBody().UploadDocumentRequest.XmlFile = "bastionairport"
	// req.RequestBody().UploadDocumentRequest.AttachmentFile = ""
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

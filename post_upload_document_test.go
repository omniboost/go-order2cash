package order2cash_test

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/omniboost/go-order2cash"
)

func TestPostUploadDocumentRequest(t *testing.T) {
	// base := order2cash.Base64Binary("test")
	// x, _ := xml.MarshalIndent(base, "", "  ")
	// log.Fatal(string(x))

	filePath := "./test.xml"
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	req := client.NewPostUploadDocumentRequest()
	req.RequestBody().UploadDocumentRequest.SenderID = "bastionairport"
	req.RequestBody().UploadDocumentRequest.GUID = "test"
	req.RequestBody().UploadDocumentRequest.XMLFile = order2cash.Base64Binary(fileContent)
	// req.RequestBody().UploadDocumentRequest.AttachmentFile = ""
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

package order2cash_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetListRejectedDocumentsRequest(t *testing.T) {
	req := client.NewGetListRejectedDocumentsRequest()
	req.RequestBody().ListRejectedDocumentsRequest.SenderID = "bastionairport"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

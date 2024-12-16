package order2cash_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	order2cash "github.com/omniboost/go-order2cash"
)

var (
	client *order2cash.Client
)

func TestMain(m *testing.M) {
	var err error

	baseURLString := os.Getenv("BASE_URL")
	username := os.Getenv("O2C_USERNAME")
	password := os.Getenv("O2C_PASSWORD")
	// partnerToken := os.Getenv("PARTNER_TOKEN")
	// hotelCode := os.Getenv("HOTEL_CODE")
	// partnerCode := os.Getenv("PARTNER_CODE")
	debug := os.Getenv("DEBUG")
	var baseURL *url.URL

	client = order2cash.NewClient(nil)
	if debug != "" {
		client.SetDebug(true)
	}

	client.SetUsername(username)
	client.SetPassword(password)

	// client.SetPartnerToken(partnerToken)
	// client.SetHotelCode(hotelCode)
	// client.SetPartnerCode(partnerCode)

	if baseURLString != "" {
		baseURL, err = url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
	}

	if baseURL != nil {
		client.SetBaseURL(*baseURL)
	}

	m.Run()
}

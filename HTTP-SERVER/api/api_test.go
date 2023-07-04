package api

import (
	"HTTP-SERVER/receipt"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type receiptTest struct {
	json           []byte
	expectedOutput receipt.Receipt
	expectedPoints int
}

var receiptTests = []receiptTest{
	{
		[]byte(`{
			"retailer": "Target",
			"purchaseDate": "2022-01-01",
			"purchaseTime": "13:01",
			"items": [
			  {
				"shortDescription": "Mountain Dew 12PK",
				"price": "6.49"
			  },{
				"shortDescription": "Emils Cheese Pizza",
				"price": "12.25"
			  },{
				"shortDescription": "Knorr Creamy Chicken",
				"price": "1.26"
			  },{
				"shortDescription": "Doritos Nacho Cheese",
				"price": "3.35"
			  },{
				"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
				"price": "12.00"
			  }
			],
			"total": "35.35"
		} `),
		receipt.Receipt{
			Retailer:     "Target",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "13:01",
			Items: []receipt.ReceiptItem{
				{
					ShortDescription: "Mountain Dew 12PK",
					Price:            "6.49",
				}, {
					ShortDescription: "Emils Cheese Pizza",
					Price:            "12.25",
				}, {
					ShortDescription: "Knorr Creamy Chicken",
					Price:            "1.26",
				}, {
					ShortDescription: "Doritos Nacho Cheese",
					Price:            "3.35",
				}, {
					ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
					Price:            "12.00",
				},
			},
			Total: "35.35",
		},
		28,
	},
	{
		[]byte(`{
			"retailer": "M&M Corner Market",
			"purchaseDate": "2022-03-20",
			"purchaseTime": "14:33",
			"items": [
			  {
				"shortDescription": "Gatorade",
				"price": "2.25"
			  },{
				"shortDescription": "Gatorade",
				"price": "2.25"
			  },{
				"shortDescription": "Gatorade",
				"price": "2.25"
			  },{
				"shortDescription": "Gatorade",
				"price": "2.25"
			  }
			],
			"total": "9.00"
		}`),
		receipt.Receipt{
			Retailer:     "M&M Corner Market",
			PurchaseDate: "2022-03-20",
			PurchaseTime: "14:33",
			Items: []receipt.ReceiptItem{
				{
					ShortDescription: "Gatorade",
					Price:            "2.25",
				}, {
					ShortDescription: "Gatorade",
					Price:            "2.25",
				}, {
					ShortDescription: "Gatorade",
					Price:            "2.25",
				}, {
					ShortDescription: "Gatorade",
					Price:            "2.25",
				},
			},
			Total: "9.00",
		},
		109,
	},
}

func TestReceiptGetValidID(t *testing.T) {

	for _, test := range receiptTests {
		//Creates new request loaded with a receipt test entry
		request, err := http.NewRequest("POST", "http://localhost:8080/process", bytes.NewBuffer(test.json))
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()
		server := NewServer()
		handler := http.HandlerFunc(server.createReceipt())
		handler.ServeHTTP(recorder, request)
		status := recorder.Code
		//Checks the status of the HTTP Recorder
		if status != http.StatusOK {
			t.Errorf("Handler returned bad status code. Got %v, expected %v.", status, http.StatusOK)
		}
		//Checks the status of the HTTP Recorder
		var idJSON map[string]string
		err = json.Unmarshal(recorder.Body.Bytes(), &idJSON)
		id := idJSON["id"]
		if err != nil {
			t.Errorf("Couldn't Parse ID in URL. %v", err)
		}
		//Checks if the id is correctly corresponding to the receipt entry
		if server.receiptPoints[0].ID.String() != id {
			t.Errorf("HTTP Request returned bad ID. Got %s, expected %s", id, server.receiptPoints[0].ID.String())
		}
	}
}

func TestReceiptPointsCalculation(t *testing.T) {

	for _, test := range receiptTests {
		//Creates new request loaded with a receipt test entry
		request, err := http.NewRequest("POST", "http://localhost:8080/process", bytes.NewBuffer(test.json))
		if err != nil {
			t.Fatal(err)
		}
		request.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()
		server := NewServer()
		handler := http.HandlerFunc(server.createReceipt())
		handler.ServeHTTP(recorder, request)
		status := recorder.Code
		//Checks the status of the HTTP Recorder
		if status != http.StatusOK {
			t.Errorf("Handler returned bad status code. Got %v, expected %v.", status, http.StatusOK)
		}

		//Converts JSON returned from receipt creation into string map containing the ID of the receipt
		var idJSON map[string]string
		err = json.Unmarshal(recorder.Body.Bytes(), &idJSON)
		id := idJSON["id"]
		if err != nil {
			t.Errorf("Couldn't Parse ID in URL. %v", err)
		}
		if server.receiptPoints[0].ID.String() != id {
			t.Errorf("HTTP Request returned bad ID. Got %s, expected %s", id, server.receiptPoints[0].ID.String())
		}
		//Checks if points are correct
		points := server.receiptPoints[0].Points
		if points != test.expectedPoints {
			t.Fatalf("Incorrect Points Calculation. Got %d, expected %d.", points, test.expectedPoints)
		}

	}
}

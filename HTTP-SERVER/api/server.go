package api

import (
	"HTTP-SERVER/receipt"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
	//receipts []receipt.Receipt
	receiptPoints []receipt.ReceiptPoints
}

func NewServer() *Server {
	server := &Server{
		Router: mux.NewRouter(),
		//receipts: []receipt.Receipt{}, //used for debugging purposes
		receiptPoints: []receipt.ReceiptPoints{},
	}
	server.routes()
	return server
}

// All possible ways of interacting with the server
// API Routes
func (s *Server) routes() {
	//s.HandleFunc("/receipts", s.listReceipts()).Methods("GET") //used for debugging and getting list of receipts
	s.HandleFunc("/receipts/process", s.createReceipt()).Methods("POST")
	s.HandleFunc("/receipts/{id}", s.removeReceipt()).Methods("DELETE")
	s.HandleFunc("/receipts/{id}", s.getReceiptPoints()).Methods("GET")
}

// On POST HTTP Request, creates a new Receipt object with the given JSON data,
// then calculates the points for the receipt, then stores the points with the receipt ID in server memory (receiptPoints).
// The receipt ID is then outputed to the user in JSON format.
func (s *Server) createReceipt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rct receipt.Receipt
		err := json.NewDecoder(r.Body).Decode(&rct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var receiptPoints receipt.ReceiptPoints
		receiptPoints.ID = uuid.New()
		receiptPoints.Points = rct.CalculatePoints()
		s.receiptPoints = append(s.receiptPoints, receiptPoints)

		idOutput := map[string]string{"id": receiptPoints.ID.String()}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(idOutput)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

//used for debugging purposes
/*func (s *Server) listReceipts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.receiptPoints); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}*/

// On DELETE HTTP Request, tries to parse the ID given within the request and iterates over the list of ReceiptPoints objects until it finds it or finishes iterating over the list.
// It then deletes that object from the list or calls an error if something goes wrong or the item cannot be found
func (s *Server) removeReceipt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]   //gets the string containing the id that is inputed into the URL of the DELETE request
		id, err := uuid.Parse(idStr) //parses the id string
		if err != nil {              //issue parsing the id string
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		foundItem := false
		for i, receipt := range s.receiptPoints { //checks each ReceiptPoints object stored in the server
			if receipt.ID == id { //if that object's id is the one we are looking for
				s.receiptPoints = append(s.receiptPoints[:i], s.receiptPoints[i+1:]...) //then remove the object from the list (by splitting the list into two around the given object and rejoining them without the object)
				foundItem = true
				break
			}

		}
		if !foundItem {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
	}
}

// On GET HTTP Request, tries to parse the ID given within the request and iterates over the list of ReceiptPoints objects until it finds it or finishes iterating over the list.
// It then outputs the receipt points value in JSON format.
func (s *Server) getReceiptPoints() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		for _, receiptPoints := range s.receiptPoints {
			if receiptPoints.ID == id {
				pointsOutput := map[string]int{"points": receiptPoints.Points}
				w.Header().Set("Content-Type", "application/json")
				err = json.NewEncoder(w).Encode(pointsOutput)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				break
			}
		}
	}
}

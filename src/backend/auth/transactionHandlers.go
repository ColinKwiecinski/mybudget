package auth

// These handlers are designed to recieve input from the client side
// of the application and process and store the transaction data

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	// "time"
	"github.com/gorilla/mux"
)

type Transaction struct {
	ID     int64  `json:"id"`
	UID    int64  `json:"uid"`
	Name   string `json:"name"`
	Memo   string `json:"memo"`
	Date   string `json:"date"`
	Amount string `json:"amount"`
	Type   string `json:"type"`
}

// Endpoints:
// POST /transactions - add new transaction into users transaction table
// GET /transactions/{UID} - Get a specific user's transaction data
// DELETE /transaction/{UID}

func (ctx *HandlerContext) TransactionHandler(w http.ResponseWriter, r *http.Request) {

	// Test POST endpoint, test headers, check for valid body format,
	// parse body or single row. Attempt to input into db, return errors if something goes wrong
	if r.Method == http.MethodPost {
		temp := Transaction{}
		denc := json.NewDecoder(r.Body)
		if err := denc.Decode(&temp); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err := ctx.Users.InsertTransaction(&temp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("input successful"))
	} else {
		statusNotAllowed(w)
	}
}

func (ctx *HandlerContext) SpecificTransactionHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		endpoint := mux.Vars(r)["UID"]
		res, err := ctx.Users.GetTransactions(endpoint)
		if err != nil {
			http.Error(w, "error getting transactions", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		enc := json.NewEncoder(w)
		if err := enc.Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	} else if r.Method == http.MethodDelete {
		endpoint := mux.Vars(r)["UID"]
		endpointInt, err := strconv.ParseInt(endpoint, 0, 64)
		if err != nil {
			http.Error(w, "error bad id input", http.StatusBadRequest)
			return
		}
		err = ctx.Users.DeleteTransaction(endpointInt)
		if err != nil {
			http.Error(w, "error while deleting row", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("entry deleted"))
	} else {
		statusNotAllowed(w)
	}
}

// Response for when an unsupported method is attempted
func statusNotAllowed(w http.ResponseWriter) {
	http.Error(w, "invalid status method attempted", http.StatusMethodNotAllowed)
}

// Checks if the header starts with application/json
// Returns true if an there is no match, false if set properly
func checkJSONHeader(w http.ResponseWriter, r *http.Request) bool {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "request body must be of type application/json", http.StatusUnsupportedMediaType)
		return true
	}
	return false
}

func processCSV(w http.ResponseWriter, r *http.Request) bool {
	r.ParseMultipartForm(10 << 20) // Limit upload file size to 10MB
	// TODO: Undo underscore and use file. Not implemented yet
	_, handler, err := r.FormFile("myFile") // TODO: do I need to change myfile to actually grab the POST body?
	if err != nil {
		http.Error(w, "error while processing csv upload: "+err.Error(), http.StatusBadRequest)
		return false
	}
	fileName := handler.Filename
	tempFile, err := ioutil.TempFile("backend/handlers/temp-csv", "upload-*.csv")
	if err != nil {
		http.Error(w, "error processing csv upload: "+err.Error(), http.StatusBadRequest)
		return false
	}
	defer tempFile.Close()

	csvFile, err := os.Open(fileName) // TODO: verify usage of fileName instead of some other path string
	if err != nil {
		http.Error(w, "file not found when retrieving csv", http.StatusInternalServerError)
		return false
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	content, _ := reader.ReadAll()
	if len(content) < 1 {
		http.Error(w, "empty csv file", http.StatusInternalServerError)
		return false
	}
	// TODO: Left off here. Need to continue processing csv file. Find a way to marshall each column into the fields of the struct
	// and then create an array of struct objects and repeatedly loop/call the database function to upload the transactions into the db.

	return true
}

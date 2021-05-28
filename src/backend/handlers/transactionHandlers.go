// These handlers are designed to recieve input from the client side
// of the application and process and store the transaction data
package handlers

// Endpoints:
// POST /transactions - add new transaction into users transaction table
// GET /transactions{UID} - Get a specific user's transaction data
// PATCH /transaction/{UID}/{TID} - Update single row transaction
// DELETE /transaction/{UID}/{TID}

func (ctx *HandlerContext) TransactionHandler(w http.ResponseWriter, r *http.Request) {
	// Test POST endpoint, test headers, check for valid body format,
	// parse body or single row. Attempt to input into db, return errors if something goes wrong
}

func (ctx *HandlerContext) SpecificTransactionHandler(w http.ResponseWriter, r *http.Request) {
	// Test for each of the three GET PATCH DELETE methods. Verify headers. Verify body for patch
}


// Helper functions

// Encode user data to JSON and write to http output
func writeUser(w http.ResponseWriter, user *users.User, status int) bool {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	resp, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}
	enc.Encode(resp)
	return true
}

// Response for when an unsupported method is attempted
func statusNotAllowed(w http.ResponseWriter) {
	http.Error(w, "invalid status method attempted", http.StatusMethodNotAllowed)
}

// Checks if the header starts with application/json
// Returns true if an there is no match
func checkJSONHeader(w http.ResponseWriter, r *http.Request) bool {
	if (r.Header.Get("Content-Type") != "application/json") {
		http.Error(w, "request body must be of type application/json", http.StatusUnsupportedMediaType)
		return true
	}
	return false
}
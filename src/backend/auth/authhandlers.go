package auth

import (
	"assignments-limj27/servers/gateway/models/users"
	"encoding/json"
	"log"

	"github.com/gorilla/mux"
	"mybudget.com/src/backend/sessions"

	// "mybudget/src/backend/users"
	"net/http"
	"strconv"
	"time"
	// "mybudget.com/src/backend/sessions"
	// "mybudget.com/src/backend/users"
)

//Handler to handle /users
func (hc *HandlerContext) UsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST Method is allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Request body must be in JSON", http.StatusUnsupportedMediaType)
		return
	}
	newUser := &NewUser{}

	denc := json.NewDecoder(r.Body)
	if err := denc.Decode(newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := newUser.ToUser()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user2, err := hc.Users.Insert(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sessionState := SessionState{time.Now(), user2}
	_, err2 := sessions.BeginSession(hc.SigningKey, hc.Sessions, &sessionState, w)

	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	enc := json.NewEncoder(w)
	if err := enc.Encode(user2); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

//Handler to handle /users/{UID}
func (hc *HandlerContext) SpecificUserHandler(w http.ResponseWriter, r *http.Request) {
	stringID := mux.Vars(r)
	log.Print(stringID)
	sessionState := SessionState{}
	_, err := sessions.GetState(r, hc.SigningKey, hc.Sessions, &sessionState)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	user := sessionState.User
	if r.Method == "GET" {
		if stringID["UID"] != "me" {
			id, err := strconv.ParseInt(stringID["UID"], 10, 64)
			if err != nil {
				http.Error(w, "couldn't convert string to int", http.StatusBadRequest)
				return
			}
			log.Println(id)
			user, err = hc.Users.GetByID(id)
			if err != nil {
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}
			if stringID["UID"] != strconv.Itoa(int(user.ID)) {
				http.Error(w, "User not authanticated", http.StatusForbidden)
				return
			}
		}
		user, _ = hc.Users.GetByID(user.ID)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		enc := json.NewEncoder(w)
		if err := enc.Encode(user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else if r.Method == "PATCH" {
		currUser := sessionState.User
		if stringID["UID"] != "me" && stringID["UID"] != strconv.Itoa(int(currUser.ID)) {
			http.Error(w, "User not authanticated", http.StatusForbidden)
			return
		}
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Request body must be JSON", http.StatusUnsupportedMediaType)
			return
		}
		newUpdates := &Updates{}
		denc := json.NewDecoder(r.Body)
		if err := denc.Decode(newUpdates); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user2, err2 := hc.Users.Update(currUser.ID, newUpdates)
		if err2 != nil {
			http.Error(w, err2.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		enc := json.NewEncoder(w)
		if err := enc.Encode(user2); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else if r.Method == "DELETE" {
		currUser := sessionState.User
		hc.Users.Delete(currUser.ID)

		_, err := sessions.EndSession(r, hc.SigningKey, hc.Sessions)
		if err != nil {
			http.Error(w, "Couldn't end session", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("User Deleted"))
	} else {
		http.Error(w, "Only PATCH, DELETE, and GET Method is allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (hc *HandlerContext) SessionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST Method is allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Request body must be in JSON", http.StatusUnsupportedMediaType)
		return
	}
	creds := &users.Credentials{}

	denc := json.NewDecoder(r.Body)
	if err := denc.Decode(creds); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := hc.Users.GetByEmail(creds.Email)
	if err != nil {
		http.Error(w, "Invalid Creditionals", http.StatusUnauthorized)
		return
	}

	sessionState := SessionState{time.Now(), user}
	_, err2 := sessions.BeginSession(hc.SigningKey, hc.Sessions, sessionState, w)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	enc := json.NewEncoder(w)
	if err := enc.Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (hc *HandlerContext) SpecificSessionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Only Delete Method is allowed", http.StatusMethodNotAllowed)
		return
	}
	stringID := mux.Vars(r)["UID"]
	if stringID != "mine" {
		http.Error(w, "User not authticated", http.StatusForbidden)
		return
	}
	sessions.EndSession(r, hc.SigningKey, hc.Sessions)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("signed out"))
}

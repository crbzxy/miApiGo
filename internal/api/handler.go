package api

import (
	"encoding/json"
	"miapigo/internal/user"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		users, err := user.GetAllUsers()
		if err != nil {
			http.Error(w, "Failed to get users", http.StatusInternalServerError)
			return
		}
		respondJSON(w, http.StatusOK, users)
	case "POST":
		var u user.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		createdUser, err := user.AddUser(u)
		if err != nil {
			http.Error(w, "Failed to add user", http.StatusInternalServerError)
			return
		}
		respondJSON(w, http.StatusCreated, createdUser)
	}
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		u, err := user.GetUserByID(id)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		respondJSON(w, http.StatusOK, u)
	case "PUT":
		var u user.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		updatedUser, err := user.UpdateUser(id, u)
		if err != nil {
			http.Error(w, "Failed to update user", http.StatusInternalServerError)
			return
		}
		respondJSON(w, http.StatusOK, updatedUser)
	case "DELETE":
		if err := user.DeleteUser(id); err != nil {
			http.Error(w, "Failed to delete user", http.StatusInternalServerError)
			return
		}
		respondJSON(w, http.StatusOK, nil)
	}
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

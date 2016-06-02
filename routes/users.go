package routes

import (
	"net/http"
)

// POST /api/users
func HandlePostUsers(w http.ResponseWriter, r *http.Request) {}

func HandleGetUserList(w http.ResponseWriter, r *http.Request) {}

// GET /api/users/<id>
func HandleGetUserByID(w http.ResponseWriter, r *http.Request) {}

// DEL /api/users/<id>
func HandleDelUser(w http.ResponseWriter, r *http.Request) {}

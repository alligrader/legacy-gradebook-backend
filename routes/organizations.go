package routes

import (
	"net/http"
)

// POST /api/organizations
func HandlePostOrganizations(w http.ResponseWriter, r *http.Request) {}

// GET /api/organizations
func HandleGetOrganizationList(w http.ResponseWriter, r *http.Request) {}

// GET /api/organizations/<id>
func HandleGetOrganizationByID(w http.ResponseWriter, r *http.Request) {}

// POST /api/organizations/<id>/teachers
func HandleAddNewTeachers(w http.ResponseWriter, r *http.Request) {}

// POST /api/organizations/<id>/billing
func HandleAddBilling(w http.ResponseWriter, r *http.Request) {}

// GET /api/organizations/<id>/billing
func HandleGetBilling(w http.ResponseWriter, r *http.Request) {}

// DEL /api/organizations/<id>/billing
func HandleDelBilling(w http.ResponseWriter, r *http.Request) {}

// DEL /api/organizations/<id>
func HandleDelOrganization(w http.ResponseWriter, r *http.Request) {}

// DEL /api/organizations/<id>/classes/<id>
func HandleDelClass(w http.ResponseWriter, r *http.Request) {}

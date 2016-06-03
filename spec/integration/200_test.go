package integration

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alligrader/gradebook-backend/routes"
)

func TestServer(t *testing.T) {

	requestLayout := []struct {
		Method string
		UrlExt string
	}{
		{"POST", "/classes/1/users"},
		{"POST", "/classes"},
		{"GET", "/classes/1"},
		{"POST", "/organizations"},
		{"GET", "/organizations/1"},
		{"POST", "/organizations/1/teachers"},
		{"POST", "/organizations/1/billing"},
		{"GET", "/organizations/1/billing"},
		{"DELETE", "/organizations/1/billing"},
		{"DELETE", "/organizations/1"},
		{"DELETE", "/organizations/1/classes/1"},
	}
	urlTmpl := "http://localhost:8000/api%s"

	for _, req := range requestLayout {
		urlExt, meth := req.UrlExt, req.Method

		url := fmt.Sprintf(urlTmpl, urlExt)
		r, err := http.NewRequest(meth, url, nil)
		if err != nil {
			t.Error(err)
		}
		w := httptest.NewRecorder()
		routes.R.ServeHTTP(w, r)
		if w.Code != 200 {
			t.Errorf("For URL %v:\nExpected Response Code %v\nFound Response Code %v\nDumping body:\n%v", url, 200, w.Code, w.Body.String())
		}
	}
}

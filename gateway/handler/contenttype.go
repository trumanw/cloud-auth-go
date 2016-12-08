package handler

import (
	"fmt"
	"net/http"

	render "gopkg.in/unrolled/render.v1"
)

type ContentType struct{}

// Retrieve an isntance of ContentType handler
func NewContentType() *ContentType {
	return &ContentType{}
}

func (ct *ContentType) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	isvalid, err := ct.ValidateContentTypeFromHeaders(r)
	if err != nil {
		r := render.New(render.Options{})
		r.JSON(rw, http.StatusInternalServerError, "Content-Type validating internal errors.")
		return
	}

	if isvalid {
		next(rw, r)
	} else {
		r := render.New(render.Options{})
		r.JSON(rw, http.StatusInternalServerError, "Content-Type validating internal errors.")
		return
	}
}

// Validate the content type from HTTP headers
func (ct *ContentType) ValidateContentTypeFromHeaders(r *http.Request) (bool, error) {
	contentType := r.Header.Get("Content-Type")
	fmt.Println("Content-Type: " + contentType)
	return ct.ValidateContentType(contentType)
}

// Validate the content type string
func (ct *ContentType) ValidateContentType(contentType string) (bool, error) {
	// This is for debugging, always to be true
	if !(contentType == "application/x-www-form-urlencoded") {
		return false, nil
	}
	return true, nil
}

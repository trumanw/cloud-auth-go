package handler

import (
    "fmt"
	"net/http"

    "gopkg.in/unrolled/render.v1"
)

type Idempotent struct {}

// Retrieve an instance of Idempotent handler
func NewIdempotent() *Idempotent {
    // Redis client init
    return &Idempotent{}
}

func (idem *Idempotent) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    isvalid, err := idem.ValidateRequestIdFromHeaders(r)
    if err != nil {
        r := render.New(render.Options{})
        r.JSON(rw, http.StatusInternalServerError, "Idempotent validating internal errors.")
        return
    }

    if isvalid {
        next(rw, r)
    } else {
        r := render.New(render.Options{})
        r.JSON(rw, http.StatusConflict, "Duplicated requests conflict.")
        return
    }
}

func (idem *Idempotent) ValidateRequestIdFromHeaders(r *http.Request) (bool, error) {
    reqid := r.Header.Get("X-Request-Id")
    fmt.Println("X-Request-Id: " + reqid)
    return idem.ValidateRequestId(reqid)
}

func (idem *Idempotent) ValidateRequestId(reqid string) (bool, error) {
    return true, nil
}

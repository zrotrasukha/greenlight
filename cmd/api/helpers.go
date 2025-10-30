package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	// ParseInt has been used here instead of Itoa is because Itoa has
	// hardcoded values, and it set bitsize to be according to the
	// cpu program is being currently running, it could be 8, 16, 32, 64
	// here we are setting it 64 and it won't change, it can store upto 9 quintillion
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

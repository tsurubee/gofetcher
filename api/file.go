package api

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func FetchAuthorizedKeys(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Hello, %s\n", p.ByName("username"))
}
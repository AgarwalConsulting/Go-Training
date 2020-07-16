package transport

import (
	"net/http"

	"algogrit.com/emp-server/auth"
	"algogrit.com/emp-server/employee/endpoints"
	basic "github.com/go-kit/kit/auth/basic"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(e endpoints.Endpoint) http.Handler {
	r := mux.NewRouter()

	basicAuthMdw := basic.AuthMiddleware(auth.Username, auth.Password, auth.Realm)

	r.Handle("/employees/{id}", httptransport.NewServer(
		basicAuthMdw(e.Show),
		decodeShowRequest,
		encodeShowResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
	)).Methods("GET")

	r.Handle("/employees", httptransport.NewServer(
		basicAuthMdw(e.Create),
		decodeCreateRequest,
		encodeCreateResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
	)).Methods("POST")

	return r
}

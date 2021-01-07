package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"algogrit.com/emp-server/employee/endpoints"
	"github.com/gorilla/mux"
)

func decodeCreateRequest(ctx context.Context, req *http.Request) (request interface{}, err error) {
	var createReq endpoints.CreateRequest

	json.NewDecoder(req.Body).Decode(&createReq.Employee)

	return createReq, nil
}

func encodeCreateResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	createRes := response.(endpoints.CreateResponse)

	return json.NewEncoder(w).Encode(createRes.Employee)
}

func decodeShowRequest(ctx context.Context, req *http.Request) (request interface{}, err error) {
	pathVars := mux.Vars(req)

	id, err := strconv.Atoi(pathVars["id"])

	if err != nil {
		return nil, err
	}

	request = endpoints.ShowRequest{id}

	return request, nil
}

func encodeShowResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	showResponse := response.(endpoints.ShowResponse)
	return json.NewEncoder(w).Encode(showResponse.Employee)
}

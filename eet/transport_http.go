package eet

import (
	"context"
	"net/http"
	"encoding/json"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	httptransport "github.com/go-kit/kit/transport/http"
)

type JoinMeetingRequest struct {

}

type LeaveMeetingRequest struct {

}

// MakeHTTPHandler returns a handler for the iteming service.
func MakeItemHTTPHandler(ctx context.Context, endpoints ServiceEndpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	r.Methods("POST").Path("/join_meeting/").Handler(httptransport.NewServer(
		endpoints.AddMeetingMemberEndpoint,
		decodeJoinMeetingRequest,
		encodeJoinMeetingResponse,
	))
	r.Methods("POST").Path("/leave_meeting/").Handler(httptransport.NewServer(
		endpoints.RemoveMeetingMemberEndpoint,
		decodeLeaveMeetingRequest,
		encodeLeaveMeetingResponse,
	))

	return r
}

func decodeJoinMeetingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request JoinMeetingRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeLeaveMeetingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request LeaveMeetingRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeLeaveMeetingResponse(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	return json.NewEncoder(w).Encode(resp)
}

func encodeJoinMeetingResponse(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	return json.NewEncoder(w).Encode(resp)
}

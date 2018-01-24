package eet

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints collects all of the endpoints that compose an item service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type ServiceEndpoints struct {
	AddMeetingMemberEndpoint    endpoint.Endpoint
	RemoveMeetingMemberEndpoint endpoint.Endpoint
}

func (se *ServiceEndpoints) CreateMeetingMembers(ctx context.Context, item []*MeetingMember) ([]*MeetingMember, error) {
	items, err := se.AddMeetingMemberEndpoint(ctx, item)
	if i, ok := items.([]*MeetingMember); ok {
		return i, err
	}
	return nil, err
}

func (se *ServiceEndpoints) RemoveMeetingMembers(ctx context.Context, item []*MeetingMember) ([]*MeetingMember, error) {
	items, err := se.RemoveMeetingMemberEndpoint(ctx, item)
	if i, ok := items.([]*MeetingMember); ok {
		return i, err
	}
	return nil, err
}

// MakeServerEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the provided service. Useful in a item server.
func MakeItemServerEndpoints(s Service) ServiceEndpoints {
	return ServiceEndpoints{
		AddMeetingMemberEndpoint:    makeAddMeetingMemberEndpoint(s),
		RemoveMeetingMemberEndpoint: makeRemoveMeetingMemberEndpoint(s),
	}
}

func makeRemoveMeetingMemberEndpoint(ts Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return ts.RemoveMeetingMembers(ctx, nil)
	}
}

func makeAddMeetingMemberEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var items = request.([]*MeetingMember)
		return s.AddMeetingMembers(ctx, items)
	}
}

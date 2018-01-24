package eet

import (
	"time"
	"context"
)

// Service is the interface that provides the basic Order methods.
type Service interface {
	// AddMeetingMembers creates a new meeting request
	AddMeetingMembers(ctx context.Context, mr []*MeetingMember) ([]*MeetingMember, error)
	// RemoveeMeetingMembers creates a new meeting request
	RemoveMeetingMembers(ctx context.Context, mr []*MeetingMember) ([]*MeetingMember, error)
}

type MeetingMemberFilter struct {
	GroupId  string
	From     time.Time
	To       time.Time
}

// Repository provides access to a meeting request store.
type Repository interface {
	StoreMeeetingRequest(mr *MeetingMember) error
	FindMeetingMembers(filter *MeetingMemberFilter) ([]*MeetingMember, error)
	DeleteMeetingMembers(ids []int) (error)

	// Ping contacts the repository to see if there are any errors communicating
	Ping() error
}

type service struct {
	repo Repository
}

// NewService returns a new instance of the default Order Service.
func NewService(meetingMembers Repository) Service {
	return &service{
		repo: meetingMembers,
	}
}

// MeetingMember contains only top level information about the item. All info about charges is contained on
// the child `item` object
type MeetingMember struct {
	Id         string       `db:"id"`
	SlackUserId       string       `db:"slack_group_id"`
	SlackGroupId string `db:"slack_group_id"`
	CreatedAt time.Time `db:"created_at"`
	MeetingAt time.Time `db:"lunch_at"`
}

func (s *service) AddMeetingMembers(ctx context.Context, members []*MeetingMember) ([]*MeetingMember, error) {
	return members, nil
}

func (s *service) RemoveMeetingMembers(ctx context.Context, members []*MeetingMember) ([]*MeetingMember, error) {
	return members, nil
}

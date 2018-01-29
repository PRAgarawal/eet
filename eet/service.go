package eet

import (
	"context"
	"time"
)

// Service is the interface that provides the basic Order methods.
type Service interface {
	// AddMeetingMembers creates a new meeting request
	AddMeetingMembers(ctx context.Context, mr []*MeetingMember) ([]*MeetingMember, error)
	// RemoveeMeetingMembers creates a new meeting request
	RemoveMeetingMembers(ctx context.Context, mr []*MeetingMember) ([]*MeetingMember, error)
}

type MeetingMemberFilter struct {
	SlackIds    []string
	SlackTeamId string
	LunchAt     time.Time
}

// Repository provides access to a meeting request store.
type Repository interface {
	StoreMeeetingRequest(mr *MeetingMember) error
	FindMeetingMembers(filter *MeetingMemberFilter) ([]*MeetingMember, error)
	FindMeetingMembersByTeam(filter *MeetingMemberFilter) ([]*MeetingMember, error)
	DeleteMeetingMembers(ids []int) error

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
	Id          string    `db:"id"`
	SlackUserId string    `db:"slack_group_id"`
	SlackTeamId string    `db:"slack_team_id"`
	CreatedAt   time.Time `db:"created_at"`
	MeetingAt   time.Time `db:"meeting_at"`
	IsDeleted   bool      `db:"is_deleted"`
}

func (s *service) AddMeetingMembers(ctx context.Context, members []*MeetingMember) ([]*MeetingMember, error) {
	return members, nil
}

func (s *service) RemoveMeetingMembers(ctx context.Context, members []*MeetingMember) ([]*MeetingMember, error) {
	return members, nil
}

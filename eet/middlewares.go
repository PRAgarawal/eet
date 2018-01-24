package eet

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(Service) Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   Service
	logger log.Logger
}

func (mw loggingMiddleware) AddMeetingMembers(ctx context.Context, mrs []*MeetingMember) (requests []*MeetingMember, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "AddMeetingMembers", "took", time.Since(begin), "err", err, "severity", "INFO")
	}(time.Now())
	return mw.next.AddMeetingMembers(ctx, mrs)
}

func (mw loggingMiddleware) RemoveMeetingMembers(ctx context.Context, mrs []*MeetingMember) (requests []*MeetingMember, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "RemoveMeetingMembers", "took", time.Since(begin), "err", err, "severity", "INFO")
	}(time.Now())
	return mw.next.RemoveMeetingMembers(ctx, mrs)
}

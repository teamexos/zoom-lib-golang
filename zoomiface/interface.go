package zoomiface

import "github.com/teamexos/zoom-lib-golang"

// ClientAPI provides an interface for Zoom Client
type ClientAPI interface {
	CreateMeeting(opts zoom.CreateMeetingOptions) (zoom.Meeting, error)
	ListMeetings(opts zoom.ListMeetingsOptions) (zoom.ListMeetingsResponse, error)
	RegisterForMeeting(opts zoom.MeetingRegistrant) (zoom.RegisterForMeetingResponse, error)
	RegisterForWebinar(opts zoom.WebinarRegistrant) (zoom.RegisterForWebinarResponse, error)
	UpdateRegistrantsStatus(opts zoom.MeetingRegistrantsStatus) error
	UpdateWebinarRegistrantsStatus(opts zoom.WebinarRegistrantsStatus) error
}

// make sure zoom.Client type satisfies the ClientAPI interface
var _ ClientAPI = (*zoom.Client)(nil)

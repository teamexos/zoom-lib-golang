package zoom

import "fmt"

const (
	// RegisterForMeetingPath - path for registering a user for a meeting
	RegisterForMeetingPath = "/meetings/%d/registrants"
)

// MeetingRegistrant contains options for meeting registration
type MeetingRegistrant struct {
	MeetingID int    `json:"-" url:"-"`
	Email     string `json:"email" url:"-"`
	FirstName string `json:"first_name" url:"-"`
	LastName  string `json:"last_name" url:"-"`
}

// RegisterForMeetingResponse is the response object returned when registering for a meeting
type RegisterForMeetingResponse struct {
	MeetingID    int    `json:"id"`
	RegistrantID string `json:"registrant_id"`
	Topic        string `json:"topic"`
	StartTime    *Time  `json:"start_time"`
	JoinURL      *URL   `json:"join_url"`
}

// RegisterForMeeting registers a user for a meeting
func (c *Client) RegisterForMeeting(opts MeetingRegistrant) (RegisterForMeetingResponse, error) {
	var ret = RegisterForMeetingResponse{}
	return ret, c.requestV2(requestV2Opts{
		Method:         Post,
		Path:           fmt.Sprintf(RegisterForMeetingPath, opts.MeetingID),
		URLParameters:  opts,
		DataParameters: opts,
		Ret:            &ret,
	})
}

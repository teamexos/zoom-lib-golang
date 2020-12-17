package zoom

import (
	"fmt"
)

const (
	ReportMeetingParticipantPath = "/report/meetings/%v/participants"
)

// ReportMeetingParticipant represents a participant for the particular meeting
type ReportMeetingParticipant struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	UserEmail string `json:"user_email"`
	JoinTime  Time   `json:"join_time"`
	LeaveTime Time   `json:"leave_time"`
	Duration  int    `json:"duration"`
}

// ReportMeetingParticipantOptions defines request options
type ReportMeetingParticipantOptions struct {
	MeetingID     int
	MeetingUUID   string
	PageSize      int    `json:"page_size"`
	NextPageToken string `json:"next_page_token"`
}

// ReportMeetingParticipantResponse represents the response from the API
type ReportMeetingParticipantResponse struct {
	PageCount     int                        `json:"page_count"`
	PageSize      int                        `json:"page_size"`
	TotalRecords  int                        `json:"total_records"`
	NextPageToken string                     `json:"next_page_token"`
	Participants  []ReportMeetingParticipant `json:"participants"`
}

// ReportMeetingParticipantResponse returns the list of meeting participants
//
// NOTE:
// Zoom will return an error (3001 Meeting does not exist) if none of the participants other than the organizer
// has joined the meeting
// https://devforum.zoom.us/t/meeting-does-not-exist-error-when-meeting-exists/19128/12
//
// Depending on when the user starts the meeting, Zoom can change the UUID of the meeting
// https://devforum.zoom.us/t/why-end-meeting-doesnt-end-it-actually/4808/17
// https://devforum.zoom.us/t/uuid-gets-changed/6429
func (c *Client) ReportMeetingParticipants(opts ReportMeetingParticipantOptions) (ReportMeetingParticipantResponse, error) {
	var (
		ret  = ReportMeetingParticipantResponse{}
		path string
	)

	if len(opts.MeetingUUID) > 0 {
		path = fmt.Sprintf(ReportMeetingParticipantPath, opts.MeetingUUID)
	} else {
		path = fmt.Sprintf(ReportMeetingParticipantPath, opts.MeetingID)
	}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          path,
		URLParameters: &opts,
		Ret:           &ret,
	})
}

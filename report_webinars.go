// Report endpoints
package zoom

import (
	"fmt"
	"strconv"
)

const (
	ParticipantReportPath = "/report/webinars/%s/participants"
)

// WebinarParticipant represents a participant for a particular webinar
type WebinarParticipant struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	UserEmail string `json:"user_email"`
	JoinTime  Time   `json:"join_time"`
	LeaveTime Time   `json:"leave_time"`
	Duration  int    `json:"duration"`
}

// ParticipantsOptions contains options for Participants.
type ParticipantsOptions struct {
	WebinarID     int    `url:"-"`
	WebinarUUID   string `url:"-"`
	PageSize      int    `json:"page_size"`
	NextPageToken string `json:"next_page_token"`
}

// ParticipantsResponse represents the response from the API
type ParticipantsResponse struct {
	PageCount     int                  `json:"page_count"`
	PageSize      int                  `json:"page_size"`
	TotalRecords  int                  `json:"total_records"`
	NextPageToken string               `json:"next_page_token"`
	Participants  []WebinarParticipant `json:"participants"`
}

// Participants delegates to the Client method of the same name
func Participants(opts ParticipantsOptions) (ParticipantsResponse, error) {
	return defaultClient.Participants(opts)
}

// Participants calls /report/webinars/{webinarId}/participants,
// listing all participants for the given webinar
func (c *Client) Participants(opts ParticipantsOptions) (ParticipantsResponse, error) {
	var (
		ret  = ParticipantsResponse{}
		path string
	)
	if len(opts.WebinarUUID) > 0 {
		path = fmt.Sprintf(ParticipantReportPath, opts.WebinarUUID)
	} else {
		path = fmt.Sprintf(ParticipantReportPath, strconv.Itoa(opts.WebinarID))
	}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          path,
		URLParameters: &opts,
		Ret:           &ret,
	})
}

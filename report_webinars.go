// Report endpoints
package zoom

import "fmt"

const (
	ParticipantReportPath = "/report/webinars/%d/participants"
)

// WebinarParticipant represents a participant for a particular webinar
type WebinarParticipant struct {
	ID        string
	UserID    string
	Name      string
	UserEmail string
	JoinTime  Time
	LeaveTime Time
	Duration  int
}

// ParticipantsOptions contains options for Participants.
type ParticipantsOptions struct {
	WebinarID     int    `url:"-"`
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
	var ret = ParticipantsResponse{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          fmt.Sprintf(ParticipantReportPath, opts.WebinarID),
		URLParameters: &opts,
		Ret:           &ret,
	})
}

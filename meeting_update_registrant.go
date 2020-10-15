package zoom

import "fmt"

const (
	// UpdateRegistrantsStatusPath defines a path for updating status of meeting registrants
	UpdateRegistrantsStatusPath = "/meetings/%d/registrants/status"

	// MeetingRegistrantStatusActionApprove defines the action to approve the registrant of the meeting
	MeetingRegistrantStatusActionApprove MeetingRegistrantStatusActionType = "approve"

	// MeetingRegistrantStatusActionCancel defines the action to cancel the registrant of the meeting
	MeetingRegistrantStatusActionCancel MeetingRegistrantStatusActionType = "cancel"

	// MeetingRegistrantStatusActionDeny defines the action to deny the registrant of the meeting
	MeetingRegistrantStatusActionDeny MeetingRegistrantStatusActionType = "deny"
)

type (
	// MeetingRegistrantStatusActionType defines the type of actions on the registrant status
	MeetingRegistrantStatusActionType string

	// MeetingRegistrantDetails contains identification of the meeting registrant
	MeetingRegistrantDetails struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}

	// MeetingRegistrantsStatus contains options for updating status
	MeetingRegistrantsStatus struct {
		MeetingID     int                               `json:"-" url:"-"`
		OccurrenceIDs string                            `json:"-" url:"occurrence_id,omitempty"`
		Action        MeetingRegistrantStatusActionType `json:"action" url:"-"`
		Registrants   []MeetingRegistrantDetails        `json:"registrants" url:"-"`
	}
)

// UpdateRegistrantsStatus updates a status of meeting registrants
func (c *Client) UpdateRegistrantsStatus(opts MeetingRegistrantsStatus) error {
	return c.requestV2(requestV2Opts{
		Method:         Put,
		Path:           fmt.Sprintf(UpdateRegistrantsStatusPath, opts.MeetingID),
		URLParameters:  opts,
		DataParameters: opts,
	})
}

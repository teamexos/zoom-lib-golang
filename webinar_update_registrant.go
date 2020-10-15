package zoom

import "fmt"

const (
	// UpdateWebinarRegistrantsStatusPath defines a path for updating status of the webinar registrants
	UpdateWebinarRegistrantsStatusPath = "/webinars/%d/registrants/status"

	// WebinarRegistrantStatusActionApprove defines the action to approve the registrant of the webinar
	WebinarRegistrantStatusActionApprove WebinarRegistrantStatusActionType = "approve"

	// WebinarRegistrantStatusActionCancel defines the action to cancel the registrant of the webinar
	WebinarRegistrantStatusActionCancel WebinarRegistrantStatusActionType = "cancel"

	// WebinarRegistrantStatusActionDeny defines the action to deny the registrant of the webinar
	WebinarRegistrantStatusActionDeny WebinarRegistrantStatusActionType = "deny"
)

type (
	// WebinarRegistrantStatusActionType defines the type of actions on the registrant status
	WebinarRegistrantStatusActionType string

	// WebinarRegistrantDetails contains identification of the webinar registrant
	WebinarRegistrantDetails struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}

	// WebinarRegistrantsStatus contains options for updating status
	WebinarRegistrantsStatus struct {
		WebinarID     int                               `json:"-" url:"-"`
		OccurrenceIDs string                            `json:"-" url:"occurrence_id,omitempty"`
		Action        WebinarRegistrantStatusActionType `json:"action" url:"-"`
		Registrants   []WebinarRegistrantDetails        `json:"registrants" url:"-"`
	}
)

// UpdateWebinarRegistrantsStatus updates a status of webinar registrants
func (c *Client) UpdateWebinarRegistrantsStatus(opts WebinarRegistrantsStatus) error {
	return c.requestV2(requestV2Opts{
		Method:         Put,
		Path:           fmt.Sprintf(UpdateWebinarRegistrantsStatusPath, opts.WebinarID),
		URLParameters:  opts,
		DataParameters: opts,
	})
}

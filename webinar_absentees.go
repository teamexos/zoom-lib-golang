package zoom

import "fmt"

const (
	// WebinarAbsenteePath - v2 path for webinar absentees
	WebinarAbsenteePath = "/past_webinars/%s/absentees"
)

// ListWebinarAbsenteeOptions - options for listing webinar absentees
type ListWebinarAbsenteeOptions struct {
	WebinarUUID  string `url:"-"`
	PageSize     *int   `url:"page_size,omitempty"`
	PageNumber   *int   `url:"page_number,omitempty"`
	OccurrenceID string `url:"occurrence_id,omitempty"`
}

// ListWebinarAbsenteeResponse - response for listing webinar registrants
type ListWebinarAbsenteeResponse struct {
	PageCount    int                 `json:"page_count"`
	PageNumber   int                 `json:"page_number"`
	PageSize     int                 `json:"page_size"`
	TotalRecords int                 `json:"total_records"`
	Registrants  []WebinarRegistrant `json:"registrants"`
}

// ListWebinarAbsentees lists webinar absentees using the default client.
func ListWebinarAbsentees(opts ListWebinarAbsenteeOptions) (ListWebinarAbsenteeResponse, error) {
	return defaultClient.ListWebinarAbsentees(opts)
}

// ListWebinarAbsentees lists webinar absentees using client c
func (c *Client) ListWebinarAbsentees(opts ListWebinarAbsenteeOptions) (ListWebinarAbsenteeResponse, error) {
	var ret = ListWebinarAbsenteeResponse{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          fmt.Sprintf(WebinarAbsenteePath, opts.WebinarUUID),
		URLParameters: opts,
		Ret:           &ret,
	})
}

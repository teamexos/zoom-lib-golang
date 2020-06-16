// Metrics endpoints for Webinars
package zoom

const (
	// WebinarMetricsPath - v2 webinar metrics
	WebinarMetricsPath = "/metrics/webinars"
)

// WebinarMetric represents metrics for a particular webinar
type WebinarMetric struct {
	UUID             string `json:"uuid"`
	ID               int64  `json:"id"`
	Topic            string `json:"topic"`
	Host             string `json:"host"`
	Email            string `json:"email"`
	Department       string `json:"dept"`
	UserType         string `json:"user_type"`
	StartTime        Time   `json:"start_time"`
	EndTime          Time   `json:"end_time"`
	Duration         string `json:"duration"`
	Participants     int64  `json:"participants"`
	HasPSTN          bool   `json:"has_pstn"`
	HasVOIP          bool   `json:"has_voip"`
	Has3rdPartyAudio bool   `json:"has_3rd_party_audio"`
	HasVideo         bool   `json:"has_video"`
	HasScreenShare   bool   `json:"has_screen_share"`
	HasRecording     bool   `json:"has_recording"`
	HasSIP           bool   `json:"has_sip"`
}

// WebinarsMetricsOptions contains options for WebinarMetrics.
// url:"page_size,omitempty"
type WebinarMetricsOptions struct {
	Type          string `url:"type"`
	From          string `url:"from"`
	To            string `url:"to"`
	PageSize      int    `url:"page_size,omitempty"`
	NextPageToken string `url:"next_page_token,omitempty"`
}

// WebinarMetricsResponse represents the response from the API
type WebinarMetricsResponse struct {
	From          string          `json:"from"`
	To            string          `json:"to"`
	PageCount     int             `json:"page_count"`
	PageSize      int             `json:"page_size"`
	TotalRecords  int             `json:"total_records"`
	NextPageToken string          `json:"next_page_token"`
	Webinars      []WebinarMetric `json:"webinars"`
}

// WebinarMetrics delegates to the Client method of the same name
func WebinarMetrics(opts WebinarMetricsOptions) (WebinarMetricsResponse, error) {
	return defaultClient.WebinarMetrics(opts)
}

// WebinarMetrics calls /metrics/webinars/, listing all webinars in the given time range
func (c *Client) WebinarMetrics(opts WebinarMetricsOptions) (WebinarMetricsResponse, error) {
	var ret = WebinarMetricsResponse{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          WebinarMetricsPath,
		URLParameters: &opts,
		Ret:           &ret,
	})
}

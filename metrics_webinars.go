// Metrics endpoints for Webinars
package zoom

const (
	// WebinarMetricsPath - v2 webinar metrics
	WebinarMetricsPath = "/metrics/webinars"
)

// WebinarMetric represents metrics for a particular webinar
type WebinarMetric struct {
	UUID             string
	ID               int64
	Topic            string
	Host             string
	Email            string
	UserType         string
	StartTime        string
	EndTime          Time
	Duration         string
	Participants     int64
	HasPSTN          bool
	HasVOIP          bool
	Has3rdPartyAudio bool
	HasVideo         bool
	HasScreenShare   bool
	HasRecording     bool
	HasSIP           bool
}

// WebinarsMetricsOptions contains options for WebinarMetrics.
type WebinarMetricsOptions struct {
	Type          string `json:"type"`
	From          string `json:"from"`
	To            string `json:"to"`
	PageSize      int    `json:"page_size"`
	NextPageToken string `json:"next_page_token"`
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

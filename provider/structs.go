package provider

const (
	EventTypeNotify        = "notify"
	EventTypeImpression    = "impression"
	EventTypeClick         = "click"
	EventTypeStart         = "start"
	EventTypeFirstQuartile = "firstQuartile"
	EventTypeMidpoint      = "midpoint"
	EventTypeThirdQuartile = "thirdQuartile"
	EventTypeComplete      = "complete"
)

type (
	item struct {
		ID             string  `json:"id"`
		Title          string  `json:"title"`
		VideoDuration  string  `json:"video_duration"`
		VideoURL       string  `json:"video_url"`
		DestinationUrl string  `json:"destination_url"`
		Bid            float64 `json:"bid"`
	}

	itemEvents struct {
		ID            string `json:"id"`
		Notify        int64  `json:"notify"`
		Impression    int64  `json:"impression"`
		Start         int64  `json:"start"`
		FirstQuartile int64  `json:"first_quartile"`
		Midpoint      int64  `json:"midpoint"`
		ThirdQuartile int64  `json:"third_quartile"`
		Complete      int64  `json:"complete"`
		Click         int64  `json:"click"`
	}
)

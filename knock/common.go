package knock

import "time"

type PageInfo struct {
	PageSize int    `json:"page_size,omitempty"`
	Before   string `json:"before,omitempty"`
	After    string `json:"after,omitempty"`
}

type ChannelData struct {
	ChannelID string                 `json:"channel_id"`
	Data      map[string]interface{} `json:"data"`
}

type PreferenceSet struct {
	ID           string                 `json:"id"`
	Workflows    map[string]interface{} `json:"workflows"`
	Categories   map[string]interface{} `json:"categories"`
	ChannelTypes map[string]interface{} `json:"channel_types"`
}

func ParseRFC3339Timestamp(input string) time.Time {
	out, _ := time.Parse(time.RFC3339, input)
	return out
}

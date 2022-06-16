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
	Workflows    map[string]interface{} `json:"workflows,omitempty"`
	Categories   map[string]interface{} `json:"categories,omitempty"`
	ChannelTypes map[string]interface{} `json:"channel_types,omitempty"`
}

func ParseRFC3339Timestamp(input string) time.Time {
	out, _ := time.Parse(time.RFC3339, input)
	return out
}

func PreferencesMapAppend(m map[string]interface{}, new map[string]interface{}) map[string]interface{} {
	if m == nil {
		m = new
	} else {
		for k, v := range new {
			m[k] = v
		}
	}
	return m
}

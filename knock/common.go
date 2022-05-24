package knock

import "time"

type PageInfo struct {
	PageSize int    `json:"page_size,omitempty"`
	Before   string `json:"before,omitempty"`
	After    string `json:"after,omitempty"`
}

func ParseRFC3339Timestamp(input string) time.Time {
	out, _ := time.Parse(time.RFC3339, input)
	return out
}

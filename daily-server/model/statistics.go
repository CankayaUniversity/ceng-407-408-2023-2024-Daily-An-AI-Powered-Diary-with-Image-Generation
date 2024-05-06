package model

// StatisticsDTO represents statistical data about a user's activity
type StatisticsDTO struct {
	Dates          []string `json:"date"`
	Likes          int      `json:"likes"`          // Number of likes received
	Views          int      `json:"views"`          // Number of views
	DailiesWritten int      `json:"dailiesWritten"` // Number of dailies written
	Mood           string   `json:"mood"`           // Current mood based on user's entries
	Streak         int      `json:"streak"`         // Current streak of daily entries
	Topic          string   `json:"topic"`          // Currently focused topic
}
